package main

import (
	"go-radio/pkg/radio"
	"go-radio/pkg/radio/inputs"
	lib2 "go-radio/pkg/radio/lib"
)

func main() {
	simpleScheme := lib2.NewBoard(
		lib2.NewUnion(
			lib2.NewEnd(),
			lib2.NewEnd(),
		),
		lib2.NewEnd(),
	)
	r := radio.New(simpleScheme)
	r.Start(inputs.SimpleInputs(0, 1, 1, 1))

	r.PrintResult()
}