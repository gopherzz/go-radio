package elements

import (
	"go-radio/pkg/radio/signal"
)

// And Logical Operation scheme
type not struct {
}

func Not() *not {
	return &not{}
}

func (n not) Eval(signals []signal.Signal) []signal.Signal {
	return []signal.Signal{
		signal.New(!signals[0].Turn),
	}
}

func (n not) InputsLength() int {
	return 1
}

func (n not) OutputLength() int {
	return 1
}
