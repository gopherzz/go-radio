package elements

import (
	"go-radio/pkg/radio/scheme"
	"go-radio/pkg/radio/signal"
)

type pair struct {
	inner   []scheme.Scheme
	outputs []signal.Signal
}

func (u pair) Eval(signals []signal.Signal) []signal.Signal {
	mainInputs := u.inner[0].InputsLength()
	mainOutputs := u.inner[0].Eval(signals[:mainInputs])
	restInputs := signals[mainInputs:]
	restOutputs := u.inner[1].Eval(append(mainOutputs, restInputs...))

	return restOutputs
}

func (u pair) InputsLength() int {
	var length int
	for _, s := range u.inner {
		length += s.InputsLength()
	}
	return length
}

func (u pair) OutputLength() int {
	return u.inner[len(u.inner)-1].OutputLength()
}

func Pair(inner ...scheme.Scheme) *pair {
	return &pair{inner: inner}
}
