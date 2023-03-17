package generator

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestLoad(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	buf := bytes.NewBuffer(nil)
	c := Config{
		FilePath:   "../examples/binary/binary.go",
		Out:        buf,
		MetricName: "metric_name_total",
		MetricHelp: "help to the metric",
		Pattern:    "Interface",
	}
	i, err := c.load(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Run("Imports", func(t *testing.T) {
		i.Imports()
	})
	t.Run("Methods", func(t *testing.T) {
		for _, m := range i.Methods() {
			t.Run("ParamsWithTypes", func(t *testing.T) {
				m.ParamsWithTypes()
			})
			t.Run("ParamsWithoutTypes", func(t *testing.T) {
				m.ParamsWithoutTypes()
			})
			t.Run("ResultsWithoutTypes", func(t *testing.T) {
				m.ResultsWithoutTypes()
			})
			t.Run("ResultsTypes", func(t *testing.T) {
				m.ResultTypes()
			})
		}
	})
}

func TestGenerateBinary(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	buf := bytes.NewBuffer(nil)
	c := Config{
		FilePath:   "../examples/binary/binary.go",
		Out:        buf,
		MetricName: "metric_name_total",
		MetricHelp: "help to the metric",
		Pattern:    "Interface",
		Mode:       Binary,
	}

	err := c.Generate(ctx)
	if err != nil {
		t.Error(err)
	}

	f, err := os.ReadFile("../examples/binary/gen.go")
	if err != nil {
		t.Error(err)
	}

	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(string(f), string(buf.Bytes()), false)
	if !bytes.Equal(f, buf.Bytes()) {
		if len(diffs) > 0 {
			t.Error(dmp.DiffPrettyText(diffs))
		}
	}
}

func TestGenerateHandler(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	buf := bytes.NewBuffer(nil)
	c := Config{
		FilePath: "../examples/http-handler/http-handler.go",
		Out:      buf,
		Pattern:  "Server",
		Mode:     Handler,
	}

	err := c.Generate(ctx)
	if err != nil {
		t.Error(err)
	}

	f, err := os.ReadFile("../examples/http-handler/gen.go")
	if err != nil {
		t.Error(err)
	}

	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(string(f), string(buf.Bytes()), false)
	if !bytes.Equal(f, buf.Bytes()) {
		if len(diffs) > 0 {
			t.Error(dmp.DiffPrettyText(diffs))
		}
	}
}

func TestImport(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	c := Config{
		FilePath: "../examples/binary/binary.go",
		Pattern:  "Interface",
	}

	i, err := c.load(ctx)
	if err != nil {
		t.Error(err)
	}

	// Just make sure that this does not panic
	i.Imports()
}
