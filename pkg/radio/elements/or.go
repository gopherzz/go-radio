package elements

import (
	"go-radio/pkg/radio/signal"
)

/// Or Logical Operation scheme
type or struct {
}

func (o or) Eval(signals []signal.Signal) []signal.Signal {
	return []signal.Signal{
		signal.New(signals[0].Turn || signals[1].Turn),
	}
}

func (o or) InputsLength() int {
	return 2
}

func (o or) OutputLength() int {
	return 1
}

func Or() *or {
	return &or{}
}
