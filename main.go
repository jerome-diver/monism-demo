package main

import (
	. "github.com/jerome-diver/monism-demo/interfaces"
	. "github.com/jerome-diver/monism-demo/monism"
)

// Modifier Monsim for Something to modify
func setValue(value string) Modifier {
	return func(m *Monism) Monism {
		that := Monism(*m).(Something)
		that.Value = value
		return that
	}
}

// Modifier Monsim for Something to modify
func setData(data int) Modifier {
	return func(m *Monism) Monism {
		that := Monism(*m).(Something)
		that.Data = data
		return that
	}
}

// MetaDefiner Monsim for Printer owned string to modify
func sayTitle(say string) Modifier {
	return func(m *Monism) Monism {
		that := MetaDefiner{Say: say, Recover: m}
		return that
	}
}

func main() {
	printer := DefaultSomething()
	sayNothing := ModifyMonism[Printer](printer)
	sayHello := ModifyMonism[Printer](printer, sayTitle("Helo"))
	sayHelli := ModifyMonism[Printer](printer, setData(10), sayTitle("Heli"), setValue("defined"))
	sayNothing.Exec()
	sayHello.Exec()
	sayHelli.Exec()
}
