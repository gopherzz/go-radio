package elements

import (
	"go-radio/pkg/radio/scheme"
	"go-radio/pkg/radio/signal"
)

/// Board scheme for simple linear schemes run.
/// [signal1, signal2]=>[or]->[not]
type board struct {
	inner      []scheme.Scheme
	pos, count int
	outputs    []signal.Signal
}

func (b board) Eval(signals []signal.Signal) []signal.Signal {
	b.outputs = signals
	for {
		if b.pos >= b.count {
			break
		}
		b.outputs = b.inner[b.pos].Eval(b.outputs)
		b.pos++
	}
	return b.outputs
}

func (b board) InputsLength() int {
	return b.inner[0].InputsLength()
}

func (b board) OutputLength() int {
	return b.inner[b.count-1].OutputLength()
}

func Board(inner ...scheme.Scheme) board {
	return board{
		inner: inner,
		count: len(inner),
	}
}
