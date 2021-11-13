package elements

import (
	"go-radio/internal/utils"
	"go-radio/pkg/radio/scheme"
	"go-radio/pkg/radio/signal"
)

type union struct {
	inner   []scheme.Scheme
	outputs []signal.Signal
}

func (u union) Eval(signals []signal.Signal) []signal.Signal {
	var inpIdx int
	var outputs [][]signal.Signal
	outputs = make([][]signal.Signal, len(u.inner))

	for idx, s := range u.inner {
		inputsLength := s.InputsLength()
		outputs[idx] = append([]signal.Signal{}, s.Eval(signals[inpIdx:inputsLength+inpIdx])...)
		inpIdx = inputsLength
	}

	return utils.JoinSignals(outputs)
}

func (u union) InputsLength() int {
	var length int
	for _, s := range u.inner {
		length += s.InputsLength()
	}
	return length
}

func (u union) OutputLength() int {
	var length int
	for _, s := range u.inner {
		length += s.OutputLength()
	}
	return length
}

func Union(inner ...scheme.Scheme) *union {
	return &union{inner: inner}
}
