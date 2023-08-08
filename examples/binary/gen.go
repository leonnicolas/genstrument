// This code is auto generated. DO NOT EDIT.

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
	defer func() {
		c.hv.WithLabelValues("AcceptIOReader").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.Interface.AcceptIOReader(_c0)
	if err != nil {
		c.cv.WithLabelValues("AcceptIOReader", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("AcceptIOReader", "success").Inc()

	return _a0, err
}

func (c *InstrumentedInterface) AccpeptStructFromSubPkg(_c0 *pkg.EmptyStruct) error {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccpeptStructFromSubPkg").Observe(time.Since(t).Seconds())
	}()
	err := c.Interface.AccpeptStructFromSubPkg(_c0)
	if err != nil {
		c.cv.WithLabelValues("AccpeptStructFromSubPkg", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("AccpeptStructFromSubPkg", "success").Inc()

	return err
}

func (c *InstrumentedInterface) AccpeptVariadic(_c0 ...int) error {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccpeptVariadic").Observe(time.Since(t).Seconds())
	}()
	err := c.Interface.AccpeptVariadic(_c0...)
	if err != nil {
		c.cv.WithLabelValues("AccpeptVariadic", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("AccpeptVariadic", "success").Inc()

	return err
}

func (c *InstrumentedInterface) AccpeptVariadicStdLib(_c0 ...io.Reader) error {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccpeptVariadicStdLib").Observe(time.Since(t).Seconds())
	}()
	err := c.Interface.AccpeptVariadicStdLib(_c0...)
	if err != nil {
		c.cv.WithLabelValues("AccpeptVariadicStdLib", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("AccpeptVariadicStdLib", "success").Inc()

	return err
}

func (c *InstrumentedInterface) AccpeptVariadicStructFromSubPkg(_c0 *pkg.EmptyStruct, _c1 ...*pkg.EmptyStruct) error {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccpeptVariadicStructFromSubPkg").Observe(time.Since(t).Seconds())
	}()
	err := c.Interface.AccpeptVariadicStructFromSubPkg(_c0, _c1...)
	if err != nil {
		c.cv.WithLabelValues("AccpeptVariadicStructFromSubPkg", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("AccpeptVariadicStructFromSubPkg", "success").Inc()

	return err
}

func (c *InstrumentedInterface) ReturnIOReaderAndError() (io.Reader, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("ReturnIOReaderAndError").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.Interface.ReturnIOReaderAndError()
	if err != nil {
		c.cv.WithLabelValues("ReturnIOReaderAndError", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("ReturnIOReaderAndError", "success").Inc()

	return _a0, err
}

func (c *InstrumentedInterface) ReturnInt() (int, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("ReturnInt").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.Interface.ReturnInt()
	if err != nil {
		c.cv.WithLabelValues("ReturnInt", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("ReturnInt", "success").Inc()

	return _a0, err
}

func (c *InstrumentedInterface) ReturnStructFromSubPkg() (*pkg.EmptyStruct, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("ReturnStructFromSubPkg").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.Interface.ReturnStructFromSubPkg()
	if err != nil {
		c.cv.WithLabelValues("ReturnStructFromSubPkg", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("ReturnStructFromSubPkg", "success").Inc()

	return _a0, err
}

func (c *InstrumentedInterface) Simple() error {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("Simple").Observe(time.Since(t).Seconds())
	}()
	err := c.Interface.Simple()
	if err != nil {
		c.cv.WithLabelValues("Simple", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("Simple", "success").Inc()

	return err
}
