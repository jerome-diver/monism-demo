package monism

import (
	. "github.com/jerome-diver/monism-demo/interfaces"
)

// MetaDefiner is a Monism
// It has to store the defined Monism at function call time
// To do something with Meta Monism Object
type MetaDefiner struct {
	Recover *Monism
	Say     string
}

// Handle Monism contract
func (d MetaDefiner) Exec(m ...Monism) Monism {
	return *d.Recover
}

func (p MetaDefiner) ID() string {
	return "MetaDefiner"
}
