package httphandler

import (
	"net/http"

	"github.com/leonnicolas/genstrument/examples/pkg"
)

//go:generate go run ../../ --file-path http-handler.go --pattern Server --mode handler -o gen.go

type Param struct{}

type Server interface {
	NoParam(w http.ResponseWriter, r *http.Request)
	StringParam(w http.ResponseWriter, r *http.Request, id string)
	TypeParam(w http.ResponseWriter, r *http.Request, params Param)
	ExternalTypeParam(w http.ResponseWriter, r *http.Request, params pkg.EmptyStruct)
	PointerParam(w http.ResponseWriter, r *http.Request, id *string)
	SliceParam(w http.ResponseWriter, r *http.Request, id []string)
	InvalidParams(id []string)
	InvalidReturns(w http.ResponseWriter, r *http.Request) error
}
