// This code in auto generated. DO NOT EDIT.

package binary

import (
	"io"
	"time"

	"github.com/leonnicolas/genstrument/examples/pkg"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type InstrumentedInterface struct {
	Interface
	cv *prometheus.CounterVec
	hv *prometheus.HistogramVec
}

func NewInstrumentedInterface(impl Interface, r prometheus.Registerer) *InstrumentedInterface {
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
		}, []string{"method", "result"})

	return &InstrumentedInterface{
		Interface: impl,
		cv:        cv,
		hv:        hv,
	}
}

func (c *InstrumentedInterface) AcceptIOReader(_c0 io.Reader) (int, error) {
	t := time.Now()
	_a0, err := c.Interface.AcceptIOReader(_c0)
	if err != nil {
		c.cv.WithLabelValues("AcceptIOReader", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("AcceptIOReader", "success").Inc()
	c.hv.WithLabelValues("AcceptIOReader").Observe(time.Since(t).Seconds())
	return _a0, err
}

func (c *InstrumentedInterface) AccpeptStructFromSubPkg(_c0 *pkg.EmptyStruct) error {
	t := time.Now()
	err := c.Interface.AccpeptStructFromSubPkg(_c0)
	if err != nil {
		c.cv.WithLabelValues("AccpeptStructFromSubPkg", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("AccpeptStructFromSubPkg", "success").Inc()
	c.hv.WithLabelValues("AccpeptStructFromSubPkg").Observe(time.Since(t).Seconds())
	return err
}

func (c *InstrumentedInterface) ReturnIOReaderAndError() (io.Reader, error) {
	t := time.Now()
	_a0, err := c.Interface.ReturnIOReaderAndError()
	if err != nil {
		c.cv.WithLabelValues("ReturnIOReaderAndError", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("ReturnIOReaderAndError", "success").Inc()
	c.hv.WithLabelValues("ReturnIOReaderAndError").Observe(time.Since(t).Seconds())
	return _a0, err
}

func (c *InstrumentedInterface) ReturnInt() (int, error) {
	t := time.Now()
	_a0, err := c.Interface.ReturnInt()
	if err != nil {
		c.cv.WithLabelValues("ReturnInt", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("ReturnInt", "success").Inc()
	c.hv.WithLabelValues("ReturnInt").Observe(time.Since(t).Seconds())
	return _a0, err
}

func (c *InstrumentedInterface) ReturnStructFromSubPkg() (*pkg.EmptyStruct, error) {
	t := time.Now()
	_a0, err := c.Interface.ReturnStructFromSubPkg()
	if err != nil {
		c.cv.WithLabelValues("ReturnStructFromSubPkg", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("ReturnStructFromSubPkg", "success").Inc()
	c.hv.WithLabelValues("ReturnStructFromSubPkg").Observe(time.Since(t).Seconds())
	return _a0, err
}

func (c *InstrumentedInterface) Simple() error {
	t := time.Now()
	err := c.Interface.Simple()
	if err != nil {
		c.cv.WithLabelValues("Simple", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("Simple", "success").Inc()
	c.hv.WithLabelValues("Simple").Observe(time.Since(t).Seconds())
	return err
}
