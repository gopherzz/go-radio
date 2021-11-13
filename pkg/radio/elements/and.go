package elements

import (
	"go-radio/pkg/radio/signal"
)

// And Logical Operation scheme
type and struct {
}

func End() *and {
	return &and{}
}

func (e and) Eval(signals []signal.Signal) []signal.Signal {
	return []signal.Signal{
		signal.New(signals[0].Turn && signals[1].Turn),
	}
}

func (e and) InputsLength() int {
	return 2
}

func (e and) OutputLength() int {
	return 1
}
