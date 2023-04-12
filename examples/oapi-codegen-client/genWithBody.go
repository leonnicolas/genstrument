// This code in auto generated. DO NOT EDIT.

package oapicodegenclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type InstrumentedClientInterface struct {
	ClientInterface
	cv *prometheus.CounterVec
	hv *prometheus.HistogramVec
}

func NewInstrumentedClientInterface(impl ClientInterface, r prometheus.Registerer) *InstrumentedClientInterface {
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

	return &InstrumentedClientInterface{
		ClientInterface: impl,
		cv:              cv,
		hv:              hv,
	}
}

func (c *InstrumentedClientInterface) AddPet(_c0 context.Context, _c1 NewPet, _c2 ...RequestEditorFn) (*http.Response, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AddPet").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientInterface.AddPet(_c0, _c1, _c2...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("AddPet", fmt.Sprint(res.StatusCode)).Inc()

	return res, err
}

func (c *InstrumentedClientInterface) AddPetWithBody(_c0 context.Context, _c1 string, _c2 io.Reader, _c3 ...RequestEditorFn) (*http.Response, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AddPetWithBody").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientInterface.AddPetWithBody(_c0, _c1, _c2, _c3...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("AddPetWithBody", fmt.Sprint(res.StatusCode)).Inc()

	return res, err
}

func (c *InstrumentedClientInterface) DeletePet(_c0 context.Context, _c1 int64, _c2 ...RequestEditorFn) (*http.Response, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("DeletePet").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientInterface.DeletePet(_c0, _c1, _c2...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("DeletePet", fmt.Sprint(res.StatusCode)).Inc()

	return res, err
}

func (c *InstrumentedClientInterface) FindPetByID(_c0 context.Context, _c1 int64, _c2 ...RequestEditorFn) (*http.Response, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("FindPetByID").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientInterface.FindPetByID(_c0, _c1, _c2...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("FindPetByID", fmt.Sprint(res.StatusCode)).Inc()

	return res, err
}

func (c *InstrumentedClientInterface) FindPets(_c0 context.Context, _c1 *FindPetsParams, _c2 ...RequestEditorFn) (*http.Response, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("FindPets").Observe(time.Since(t).Seconds())
	}()
	res, err := c.ClientInterface.FindPets(_c0, _c1, _c2...)
	if err != nil {
		return res, err
	}
	c.cv.WithLabelValues("FindPets", fmt.Sprint(res.StatusCode)).Inc()

	return res, err
}
