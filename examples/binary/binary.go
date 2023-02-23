package binary

import (
	"io"

	"github.com/leonnicolas/genstrument/examples/pkg"
)

type Interface interface {
	Simple() error
	ReturnInt() (int error)
	AcceptIOReader(io.Reader) (int error)
	ReturnIOReaderAndError() (io.Reader, error)
	ReturnStructFromSubPkg() (*pkg.EmptyStruct, error)
	AccpeptStructFromSubPkg(*pkg.EmptyStruct) error
}
