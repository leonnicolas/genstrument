package binary

import (
	"io"

	"github.com/leonnicolas/genstrument/examples/pkg"
)

//go:generate go run ../../ --file-path binary.go -p Interface --metric-help "help to the metric" --metric-name metric_name_total -o gen.go

type Interface interface {
	Simple() error
	ReturnInt() (int error)
	AcceptIOReader(io.Reader) (int error)
	ReturnIOReaderAndError() (io.Reader, error)
	ReturnStructFromSubPkg() (*pkg.EmptyStruct, error)
	AccpeptStructFromSubPkg(*pkg.EmptyStruct) error
}
