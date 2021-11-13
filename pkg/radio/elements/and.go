package elements

import (
	"go-radio/pkg/radio/signal"
)

type end struct {
}

func End() *end {
	return &end{}
}

func (e end) Eval(signals []signal.Signal) []signal.Signal {
	return []signal.Signal{
		signal.New(signals[0].Turn && signals[1].Turn),
	}
}

func (e end) InputsLength() int {
	return 2
}

func (e end) OutputLength() int {
	return 1
}
