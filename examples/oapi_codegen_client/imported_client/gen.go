// This code in auto generated. DO NOT EDIT.

package importedclient

import (
	"context"
	"fmt"
	"io"
	"time"

	oapicodegenclient "github.com/leonnicolas/genstrument/examples/oapi_codegen_client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type InstrumentedClientWithResponsesInterface struct {
	oapicodegenclient.ClientWithResponsesInterface
	cv *prometheus.CounterVec
	hv *prometheus.HistogramVec
}

func NewInstrumentedClientWithResponsesInterface(impl oapicodegenclient.ClientWithResponsesInterface, r prometheus.Registerer) *InstrumentedClientWithResponsesInterface {
	hv := promauto.With(r).NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "metric_name_duration_seconds",
			Help:    "help to the metric",
			Buckets: prometheus.DefBuckets,
		}, []string{"method"})

	cv := promauto.With(r).NewCounterVec(
		prometheus.CounterOpts{
			Name: "metric_name_total",
			Help: "help to the metric",
		}, []string{"method", "code"})

	return &InstrumentedClientWithResponsesInterface{
		ClientWithResponsesInterface: impl,
		cv:                           cv,
		hv:                           hv,
	}
}

func (c *InstrumentedClientWithResponsesInterface) AddPetWithBodyWithResponse(_c0 context.Context, _c1 string, _c2 io.Reader, _c3 ...oapicodegenclient.RequestEditorFn) (*oapicodegenclient.AddPetResponse, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AddPetWithBodyWithResponse").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientWithResponsesInterface.AddPetWithBodyWithResponse(_c0, _c1, _c2, _c3...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("AddPetWithBodyWithResponse", fmt.Sprint(res.StatusCode())).Inc()

	return res, err
}

func (c *InstrumentedClientWithResponsesInterface) AddPetWithResponse(_c0 context.Context, _c1 oapicodegenclient.NewPet, _c2 ...oapicodegenclient.RequestEditorFn) (*oapicodegenclient.AddPetResponse, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AddPetWithResponse").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientWithResponsesInterface.AddPetWithResponse(_c0, _c1, _c2...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("AddPetWithResponse", fmt.Sprint(res.StatusCode())).Inc()

	return res, err
}

func (c *InstrumentedClientWithResponsesInterface) DeletePetWithResponse(_c0 context.Context, _c1 int64, _c2 ...oapicodegenclient.RequestEditorFn) (*oapicodegenclient.DeletePetResponse, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("DeletePetWithResponse").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientWithResponsesInterface.DeletePetWithResponse(_c0, _c1, _c2...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("DeletePetWithResponse", fmt.Sprint(res.StatusCode())).Inc()

	return res, err
}

func (c *InstrumentedClientWithResponsesInterface) FindPetByIDWithResponse(_c0 context.Context, _c1 int64, _c2 ...oapicodegenclient.RequestEditorFn) (*oapicodegenclient.FindPetByIDResponse, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("FindPetByIDWithResponse").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientWithResponsesInterface.FindPetByIDWithResponse(_c0, _c1, _c2...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("FindPetByIDWithResponse", fmt.Sprint(res.StatusCode())).Inc()

	return res, err
}

func (c *InstrumentedClientWithResponsesInterface) FindPetsWithResponse(_c0 context.Context, _c1 *oapicodegenclient.FindPetsParams, _c2 ...oapicodegenclient.RequestEditorFn) (*oapicodegenclient.FindPetsResponse, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("FindPetsWithResponse").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientWithResponsesInterface.FindPetsWithResponse(_c0, _c1, _c2...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("FindPetsWithResponse", fmt.Sprint(res.StatusCode())).Inc()

	return res, err
}
