package monism

import (
	. "github.com/jerome-diver/monism-demo/interfaces"
)

type Acceptable interface {
	Printer
}

// --- First Class, pure and Variadic function ---
// Get one or many Monism as argument
// and return a Printer Monism
//
// First Monism arg has to be a printer
// Next have to be Modifier and/or MetaDefiner
//
// This HOF use two design pattern: Decorator and Strategy patterns
// And is a composit function as it return function too.
func ModifyMonism[T Acceptable](first Monism, modifiers ...Monism) Monism {
	var next, after Monism
	switch first.(type) { // Check
	case T:
		next = first
	default:
		panic("first Monism has to be a Printer")
	}
	for _, modify := range modifiers {
		after = modify.Exec(next)
		switch after.ID() {
		case "Something":
			next = after
		case "MetaDefiner":
			meta := after.(MetaDefiner)
			printer_before := after.Exec().(Something)
			printer_before.Title = meta.Say
			next = printer_before
		}
	}
	return next
}
