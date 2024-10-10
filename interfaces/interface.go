package interfaces

type Monism interface {
	Exec(...Monism) Monism // will do something special in relation with the Monism identity
	ID() string
}

type Printer interface {
	Print()
}
