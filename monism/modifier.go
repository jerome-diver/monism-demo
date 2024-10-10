package monism

import (
	. "github.com/jerome-diver/monism-demo/interfaces"
)

// Alias as Modifier
// to be a Monism
// They will be a function to call as a Strategy
type Modifier func(*Monism) Monism

// Handle Monism contract
func (m Modifier) Exec(printer ...Monism) Monism {
	return m(&printer[0])
}

func (m Modifier) ID() string {
	return "Modifier"
}
