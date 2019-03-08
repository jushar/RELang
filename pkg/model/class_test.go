package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassGetSize(t *testing.T) {
	class := NewClass("testClass")

	assert.Equal(t, uint64(0), class.GetSize())

	var1 := NewVariable()
	var1.Name = "prop1"
	var1.Type = "int32"
	addr := uint64(0)
	var1.MemoryOffset = &addr
	var1.Size = 4
	class.AddVariable(var1)

	var2 := NewVariable()
	var2.Name = "prop2"
	var2.Type = "int8"
	addr2 := uint64(4)
	var2.MemoryOffset = &addr2
	var2.Size = 1
	class.AddVariable(var2)	

	assert.Equal(t, uint64(5), class.GetSize())
}
