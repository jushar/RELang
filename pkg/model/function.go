package model

import "fmt"

// Represents a function parameter
type Parameter struct {
	Name string
	Type string
}

// Represents a function modifier
type FunctionModifier = string

// Represents a function calling convention
type CallingConvention = string

// Represents a function within a chunk or class
type Function struct {
	Name              string
	ReturnType        string
	Modifier          FunctionModifier
	Params            []Parameter
	CallingConvention CallingConvention
	MemoryAddress     *uint64
	Public            bool
}

// Constructs a new function by the given name
func NewFunction(name string) *Function {
	return &Function{
		Name:   name,
		Public: true,
	}
}

// Constructs a new function pad
func NewFunctionPad(memoryOffset uint64) *Function {
	addr := memoryOffset

	return &Function{
		Name:              fmt.Sprintf("vpad_%x", memoryOffset),
		ReturnType:        "void",
		Modifier:          "virtual",
		CallingConvention: "",
		MemoryAddress:     &addr,
		Public:            false,
	}
}

// Adds a parameter to the function
func (s *Function) AddParameter(name string, parameterType string) {
	s.Params = append(s.Params, Parameter{Name: name, Type: parameterType})
}
