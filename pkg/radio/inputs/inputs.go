package inputs

import (
	"go-radio/pkg/radio/signal"
)

func SimpleInputs(inputs ...uint8) []signal.Signal {
	var signals []signal.Signal
	for _, i := range inputs {
		signals = append(signals, signal.New(i > 0))
	}
	return signals
}

func Inputs(inputs ...signal.Signal) []signal.Signal {
	var signals []signal.Signal
	for _, sig := range inputs {
		signals = append(signals, sig)
	}
	return signals
}