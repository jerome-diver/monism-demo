package monism

import (
	"encoding/json"
	"fmt"

	. "github.com/jerome-diver/monism-demo/interfaces"
)

// Something is a Monism struct with
// stored data to use first to build
// and handle a Print interface
type Something struct {
	Value, Title string
	Data         int
}

// Handle Printer conract
func (s Something) Print() {
	out, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Title, string(out))
}

// Handle Monism contract
func (s Something) Exec(m ...Monism) Monism {
	s.Print()
	return s
}

func (s Something) ID() string {
	return "Something"
}

func DefaultSomething() Something {
	return Something{
		Value: "",
		Title: "Something",
		Data:  0,
	}
}
