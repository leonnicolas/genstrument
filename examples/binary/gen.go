package binary

import (
	"io"

	"github.com/leonnicolas/genstrument/examples/pkg"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type InstrumentedInterface struct {
	Interface
	cv *prometheus.CounterVec
}

func NewInstrumentedInterface(impl Interface, r prometheus.Registerer) *InstrumentedInterface {
	cv := promauto.With(r).NewCounterVec(prometheus.CounterOpts{
		Name: "metric_name_total",
		Help: "help to the metric",
	}, []string{"method", "result"})

	return &InstrumentedInterface{
		Interface: impl,
		cv:        cv,
	}
}

func (c *InstrumentedInterface) AcceptIOReader(_c0 io.Reader) error {
	err := c.Interface.AcceptIOReader(_c0)
	if err != nil {
		c.cv.WithLabelValues("AcceptIOReader", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("AcceptIOReader", "success").Inc()
	return err
}

func (c *InstrumentedInterface) AccpeptStructFromSubPkg(_c0 *pkg.EmptyStruct) error {
	err := c.Interface.AccpeptStructFromSubPkg(_c0)
	if err != nil {
		c.cv.WithLabelValues("AccpeptStructFromSubPkg", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("AccpeptStructFromSubPkg", "success").Inc()
	return err
}

func (c *InstrumentedInterface) ReturnIOReaderAndError() (io.Reader, error) {
	_a0, err := c.Interface.ReturnIOReaderAndError()
	if err != nil {
		c.cv.WithLabelValues("ReturnIOReaderAndError", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("ReturnIOReaderAndError", "success").Inc()
	return _a0, err
}

func (c *InstrumentedInterface) ReturnInt() error {
	err := c.Interface.ReturnInt()
	if err != nil {
		c.cv.WithLabelValues("ReturnInt", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("ReturnInt", "success").Inc()
	return err
}

func (c *InstrumentedInterface) ReturnStructFromSubPkg() (*pkg.EmptyStruct, error) {
	_a0, err := c.Interface.ReturnStructFromSubPkg()
	if err != nil {
		c.cv.WithLabelValues("ReturnStructFromSubPkg", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("ReturnStructFromSubPkg", "success").Inc()
	return _a0, err
}

func (c *InstrumentedInterface) Simple() error {
	err := c.Interface.Simple()
	if err != nil {
		c.cv.WithLabelValues("Simple", "error").Inc()
		return err
	}
	c.cv.WithLabelValues("Simple", "success").Inc()
	return err
}
