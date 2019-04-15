package cpp

import (
	"testing"

	"github.com/Jusonex/RELang/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestCreatesVariablePads(t *testing.T) {
	chunk := prepareTestChunk()

	transformerMemoryAddresses := &TransformerMemoryAddresses{}
	transformerMemoryAddresses.Transform(chunk)

	transformer := &TransformerVariablePads{}
	transformer.Transform(chunk)

	variables := chunk.Classes[0].Variables
	assert.Equal(t, 5, len(variables))
	assert.False(t, variables[2].Public)
	assert.Equal(t, uint64(5), *variables[2].MemoryOffset)
	assert.Equal(t, uint64(16), *variables[3].MemoryOffset)
}

func TestCorrectlyConsiderVirtualMembers(t *testing.T) {
	chunk := model.NewChunk()
	class := model.NewClass("testClassWithVirtual")
	chunk.AddClass(class)

	var1 := model.NewVariable()
	var1.Name = "prop1"
	var1.Type = "int32"
	offset := uint64(4)
	var1.MemoryOffset = &offset
	class.AddVariable(var1)

	func1 := model.NewFunction("func1")
	func1.Modifier = "virtual"

	transformerMemoryAddresses := &TransformerMemoryAddresses{}
	transformerMemoryAddresses.Transform(chunk)

	transformer := &TransformerVariablePads{}
	transformer.Transform(chunk)

	assert.Equal(t, uint64(8), class.GetSize())
}
