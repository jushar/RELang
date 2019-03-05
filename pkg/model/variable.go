package model

import "fmt"

// Represents a variable within a class or chunk
type Variable struct {
	Name         string
	Type         string
	MemoryOffset *uint64
	Public       bool
}

// Constructs a new variable
func NewVariable() *Variable {
	return &Variable{Public: true}
}

// Constructs a new variable padding (gap filler)
func NewVariablePad(memoryOffset uint64, size uint64) *Variable {
	addr := memoryOffset

	return &Variable{
		// HACK: Size was embedded in name, fix this one day!
		Name:         fmt.Sprintf("_pad_%x[%d]", memoryOffset, size),
		Type:         "char",
		MemoryOffset: &addr,
		Public:       false,
	}
}
