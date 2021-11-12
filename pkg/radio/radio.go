package radio

import (
	"fmt"
	"go-radio/pkg/radio/scheme"
	"go-radio/pkg/radio/signal"
)

type radio struct {
	scheme  scheme.Scheme
	outputs []signal.Signal
}

func New(scheme scheme.Scheme) *radio {
	return &radio{
		scheme:  scheme,
	}
}

func (r *radio) Start(inputs []signal.Signal) {
	 r.outputs = r.scheme.Eval(inputs)
}

func (r *radio) PrintResult() {
	fmt.Printf("%v", r.outputs)
}