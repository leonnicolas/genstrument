// This code in auto generated. DO NOT EDIT.

package httphandler

import (
	"net/http"

	"github.com/leonnicolas/genstrument/examples/pkg"
	"github.com/metalmatze/signal/server/signalhttp"
	"github.com/prometheus/client_golang/prometheus"
)

type InstrumentedServer struct {
	Server
	signalhttp.HandlerInstrumenter
}

func NewInstrumentedServer(impl Server, r prometheus.Registerer) *InstrumentedServer {
	i := signalhttp.NewHandlerInstrumenter(r, []string{"handler"})
	return &InstrumentedServer{impl, i}
}

func (i *InstrumentedServer) ExternalTypeParam(w http.ResponseWriter, r *http.Request, _c2 pkg.EmptyStruct) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		i.Server.ExternalTypeParam(w, r, _c2)
	}
	i.NewHandler(prometheus.Labels{"handler": "ExternalTypeParam"}, http.HandlerFunc(handler))(w, r)
}

func (i *InstrumentedServer) NoParam(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		i.Server.NoParam(w, r)
	}
	i.NewHandler(prometheus.Labels{"handler": "NoParam"}, http.HandlerFunc(handler))(w, r)
}

func (i *InstrumentedServer) PointerParam(w http.ResponseWriter, r *http.Request, _c2 *string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		i.Server.PointerParam(w, r, _c2)
	}
	i.NewHandler(prometheus.Labels{"handler": "PointerParam"}, http.HandlerFunc(handler))(w, r)
}

func (i *InstrumentedServer) SliceParam(w http.ResponseWriter, r *http.Request, _c2 []string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		i.Server.SliceParam(w, r, _c2)
	}
	i.NewHandler(prometheus.Labels{"handler": "SliceParam"}, http.HandlerFunc(handler))(w, r)
}

func (i *InstrumentedServer) StringParam(w http.ResponseWriter, r *http.Request, _c2 string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		i.Server.StringParam(w, r, _c2)
	}
	i.NewHandler(prometheus.Labels{"handler": "StringParam"}, http.HandlerFunc(handler))(w, r)
}

func (i *InstrumentedServer) TypeParam(w http.ResponseWriter, r *http.Request, _c2 Param) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		i.Server.TypeParam(w, r, _c2)
	}
	i.NewHandler(prometheus.Labels{"handler": "TypeParam"}, http.HandlerFunc(handler))(w, r)
}
