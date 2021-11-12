package utils

import (
	"go-radio/pkg/radio/signal"
)

func JoinSignals(sig [][]signal.Signal) []signal.Signal {
	var result []signal.Signal
	for _, s := range sig {
		result = append(result, s...)
	}
	return result
}