package model

type Variable struct {
	Name         string
	Type         string
	MemoryOffset *uint64
}

func NewVariable() *Variable {
	return &Variable{}
}
