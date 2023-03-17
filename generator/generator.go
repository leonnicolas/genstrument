package generator

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"go/types"
	"io"
	"strings"
	"text/template"

	"github.com/vektra/mockery/v2/pkg"
	"golang.org/x/tools/imports"
)

type GeneratorMode string

const (
	Binary  GeneratorMode = "binary"
	Handler GeneratorMode = "handler"
)

var mode2tmpl = map[GeneratorMode]string{
	Binary:  "binary.tmpl",
	Handler: "handler.tmpl",
}

type Config struct {
	FilePath       string
	Pattern        string
	MetricName     string
	MetricHelp     string
	MetricHistName string
	MetricHistHelp string
	Out            io.Writer
	Mode           GeneratorMode
}

//go:embed templates/*
var tmpls embed.FS

type InstrumentedInterface struct {
	*pkg.Interface
	InstrumentedTypeName string
	c                    *Config
}

func (ii *InstrumentedInterface) Methods() (methods []*Method) {
	for _, m := range ii.Interface.Methods() {
		methods = append(methods, &Method{m, ii})
	}
	return
}

func (ii *InstrumentedInterface) MetricName() string {
	return ii.c.MetricName
}

func (ii *InstrumentedInterface) MetricHelp() string {
	return ii.c.MetricHelp
}

func (ii *InstrumentedInterface) MetricHistName() string {
	if n := ii.c.MetricHistName; n != "" {
		return n
	}
	return fmt.Sprintf("%s_duration_seconds", strings.TrimSuffix(ii.MetricName(), "_total"))
}

func (ii *InstrumentedInterface) MetricHistHelp() string {
	if n := ii.c.MetricHistHelp; n != "" {
		return n
	}
	return ii.MetricHelp()
}

func (ii *InstrumentedInterface) typeName(t types.Type) (string, *types.TypeName, error) {
	switch t := t.(type) {
	case *types.TypeParam:
		return "", t.Obj(), nil
	case *types.Named:
		return "", t.Obj(), nil
	case *types.Pointer:
		s, _t, err := ii.typeName(t.Elem())
		return "*" + s, _t, err
	case *types.Slice:
		s, _t, err := ii.typeName(t.Elem())
		return "[]" + s, _t, err
	case *types.Array:
		s, _t, err := ii.typeName(t.Elem())
		return fmt.Sprintf("[%d]%s", t.Len(), s), _t, err
	default:
		return "", nil, fmt.Errorf("unsupported type: %q", t.String())
	}
}

func (ii *InstrumentedInterface) Imports() []string {
	h := make(map[string]struct{})
	h["github.com/prometheus/client_golang/prometheus"] = struct{}{}
	h["github.com/prometheus/client_golang/prometheus/promauto"] = struct{}{}

	for _, m := range ii.Methods() {
		l := m.Signature.Params().Len()

		for i := 0; i < l; i++ {
			v := m.Signature.Params().At(i)
			_, tn, err := ii.typeName(v.Type())
			if err == nil {
				h[tn.Pkg().Path()] = struct{}{}
			}
		}
		l = m.Signature.Results().Len()

		for i := 0; i < l; i++ {
			v := m.Signature.Results().At(i)
			_, tn, err := ii.typeName(v.Type())
			if err == nil && tn.Pkg() != nil {
				h[tn.Pkg().Path()] = struct{}{}
			}
		}
	}

	delete(h, ii.Pkg.Path())

	s := make([]string, 0, len(h))
	for k := range h {
		s = append(s, k)
	}

	return s
}

type Method struct {
	*pkg.Method
	parentInterface *InstrumentedInterface
}

func (m *Method) ParamsWithTypes() (str string) {
	l := m.Signature.Params().Len()
	strs := make([]string, l)

	for i := 0; i < l; i++ {
		s, tn, err := m.parentInterface.typeName(m.Signature.Params().At(i).Type())
		if err != nil {
			strs[i] = fmt.Sprintf("_c%d %s", i, m.Signature.Params().At(i).Type().String())
			continue
		}
		if tn.Pkg().Name() == m.parentInterface.Pkg.Name() {
			strs[i] = fmt.Sprintf("_c%d %s", i, fmt.Sprintf("%s%s", s, tn.Name()))
			continue
		}
		strs[i] = fmt.Sprintf("_c%d %s", i, fmt.Sprintf("%s%s.%s", s, tn.Pkg().Name(), tn.Name()))
	}

	if m.parentInterface.c.Mode == Handler {
		strs = strs[2:]
	}
	return strings.Join(strs, ",")
}

func (m *Method) ParamsWithoutTypes() (str string) {
	l := m.Signature.Params().Len()
	strs := make([]string, l)

	for i := 0; i < l; i++ {
		strs[i] = fmt.Sprintf("_c%d", i)
	}
	if m.parentInterface.c.Mode == Handler {
		strs = strs[2:]
	}
	return strings.Join(strs, ",")
}

func (m *Method) ResultsWithoutTypes() (str string) {
	l := m.Signature.Results().Len()
	strs := make([]string, l)

	for i := 0; i < l; i++ {
		if i == l-1 {
			strs[i] = "err"

			continue
		}
		strs[i] = fmt.Sprintf("_a%d", i)
	}
	return strings.Join(strs, ",")
}

func (m *Method) ResultTypes() (str string) {
	if m.Signature == nil {
		return ""
	}
	r := m.Signature.Results()
	if r == nil {
		return ""
	}
	l := r.Len()
	strs := make([]string, l)

	for i := 0; i < l; i++ {
		s, tn, err := m.parentInterface.typeName(r.At(i).Type())
		if err != nil {
			strs[i] = r.At(i).Type().String()
			continue
		}
		if tn.Pkg() == nil {
			strs[i] = fmt.Sprintf("%s%s", s, tn.Name())
			continue
		}
		if tn.Pkg().Name() == m.parentInterface.Pkg.Name() {
			strs[i] = fmt.Sprintf("%s%s", s, tn.Name())
			continue
		}
		strs[i] = fmt.Sprintf("%s%s.%s", s, tn.Pkg().Name(), tn.Id())
	}
	return strings.Join(strs, ",")
}

func (m *Method) ReturnsError() bool {
	r := m.Signature.Results()
	if r == nil {
		return false
	}
	return r.At(m.Signature.Results().Len()-1).Type().String() == "error"
}

func (m *Method) IsHandler() bool {
	r := m.Signature.Results()
	if r != nil && r.Len() > 0 {
		return false
	}
	p := m.Signature.Params()
	if p == nil {
		return false
	}
	if l := p.Len(); l < 2 {
		return false
	}
	return p.At(0).Type().String() == "net/http.ResponseWriter" && p.At(1).Type().String() == "*net/http.Request"
}

func (c *Config) load(ctx context.Context) (*InstrumentedInterface, error) {
	p := pkg.NewParser(nil)
	if err := p.Parse(ctx, c.FilePath); err != nil {
		return nil, fmt.Errorf("failed to parse file %q: %w", c.FilePath, err)
	}
	if err := p.Load(); err != nil {
		return nil, fmt.Errorf("failed to load package from %q: %w", c.FilePath, err)
	}
	iface, err := p.Find(c.Pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to find interface %q: %w", c.Pattern, err)
	}

	return &InstrumentedInterface{
		Interface:            iface,
		InstrumentedTypeName: fmt.Sprintf("Instrumented%s", iface.Name),
		c:                    c,
	}, nil
}

// The NeedTypes LoadMode bit sets this field for packages matching the
// patterns; type information for dependencies may be missing or incomplete,
// unless NeedDeps and NeedImports are also set.

func (c *Config) Generate(ctx context.Context) error {
	if _, ok := mode2tmpl[c.Mode]; !ok {
		return fmt.Errorf("unsopperted generator mode: %s", c.Mode)
	}
	iiface, err := c.load(ctx)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFS(tmpls, "templates/*.tmpl")
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}
	buf := bytes.NewBuffer(nil)
	if err := tmpl.ExecuteTemplate(buf, mode2tmpl[c.Mode], iiface); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	opt := &imports.Options{Comments: true}

	res, err := imports.Process(c.FilePath, buf.Bytes(), opt)
	if err != nil {
		return fmt.Errorf("failed to process imports: %w", err)
	}

	buf.Reset()
	_, err = io.Copy(c.Out, bytes.NewBuffer(res))

	return err
}
