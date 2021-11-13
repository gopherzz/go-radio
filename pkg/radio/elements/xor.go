package elements

import (
	"go-radio/pkg/radio/signal"
)

// Xor Logical Operation scheme
type xor struct {
}

func (x xor) Eval(signals []signal.Signal) []signal.Signal {
	res := !(signals[0].Turn == signals[1].Turn)
	return []signal.Signal{
		signal.New(res),
	}
}

func (x xor) InputsLength() int {
	return 2
}

func (x xor) OutputLength() int {
	return 1
}

func Xor() *xor {
	return &xor{}
}
