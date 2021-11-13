package scheme

import (
	"go-radio/pkg/radio/signal"
)

//Scheme Interface for declare any radio-element
type Scheme interface {
	Eval([]signal.Signal) []signal.Signal
	InputsLength() int
	OutputLength() int
}
