package elements

import (
	"go-radio/pkg/radio/scheme"
	"go-radio/pkg/radio/signal"
)

/// Pair scheme for pass first input signals to first scheme and
/// first scheme outputs + rest input signals to second scheme
/// 					  | [sig1, sig2] => [scheme1] -> [out]
///                       |                                |
/// [sig1, sig2, sig3] -> |   +----------------------------+
///                       |   |
///                       | [out, sig3]  => [scheme2]
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
