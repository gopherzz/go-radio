package main

import (
	"go-radio/pkg/radio"
	elements "go-radio/pkg/radio/elements"
	"go-radio/pkg/radio/inputs"
)

func main() {
	simpleScheme := elements.Board(
		elements.Pair(
			elements.Board(
				elements.Union(
					elements.Or(),
					elements.End(),
				),
				elements.Or(),
			),
			elements.Xor(),
		),
	)
	r := radio.New(simpleScheme)
	r.Start(inputs.SimpleInputs(0, 1, 0, 1, 0))

	r.PrintResult()
}
