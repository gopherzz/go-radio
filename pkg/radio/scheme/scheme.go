package scheme

import (
	"go-radio/pkg/radio/signal"
)

type Scheme interface {
	Eval([]signal.Signal) []signal.Signal
	InputsLength() int
	OutputLength() int
}