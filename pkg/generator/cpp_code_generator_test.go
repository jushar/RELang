package generator

import (
	"os"
	"testing"

	"github.com/Jusonex/RELang/pkg/model"
	"github.com/stretchr/testify/assert"
)

const OUTPUT_FILE = "../../test/testoutput.h"

func prepareTestChunk() *model.Chunk {
	chunk := model.NewChunk()
	class := model.NewClass("testClass")
	chunk.AddClass(class)

	var1 := model.NewVariable()
	var1.Name = "prop1"
	var1.Type = "int32"
	class.AddVariable(var1)

	var2 := model.NewVariable()
	var2.Name = "prop2"
	var2.Type = "int8"
	class.AddVariable(var2)

	var3 := model.NewVariable()
	var3.Name = "prop3"
	var3.Type = "float64"
	addr := uint64(16)
	var3.MemoryOffset = &addr
	class.AddVariable(var3)

	var4 := model.NewVariable()
	var4.Name = "prop4"
	var4.Type = "uint16"
	class.AddVariable(var4)

	return chunk
}

func TestConstructValidGenerator(t *testing.T) {
	gen := NewCppCodeGenerator(OUTPUT_FILE)
	defer gen.Close()

	assert.NotNil(t, gen.Emitter)
	_, err := os.Stat(OUTPUT_FILE)
	assert.Nil(t, err)
}

func TestGeneratesChunk(t *testing.T) {
	gen := NewCppCodeGenerator(OUTPUT_FILE)
	chunk := prepareTestChunk()

	gen.Generate(chunk)
	gen.Close()

	stat, err := os.Stat(OUTPUT_FILE)
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 100)
}

func TestCreatesMemoryAddresses(t *testing.T) {
	gen := NewCppCodeGenerator(OUTPUT_FILE)
	defer gen.Close()
	chunk := prepareTestChunk()

	gen.CreateMemoryAddresses(chunk)

	variables := chunk.Classes[0].Variables
	assert.Equal(t, uint64(0), *variables[0].MemoryOffset)
	assert.Equal(t, uint64(4), *variables[1].MemoryOffset)
	assert.Equal(t, uint64(16), *variables[2].MemoryOffset)
	assert.Equal(t, uint64(24), *variables[3].MemoryOffset)

	// TODO: Test virtual functions too
}

func TestCreatesVariablePads(t *testing.T) {
	gen := NewCppCodeGenerator(OUTPUT_FILE)
	defer gen.Close()
	chunk := prepareTestChunk()

	gen.CreateMemoryAddresses(chunk)
	gen.CreateVariablePads(chunk)

	variables := chunk.Classes[0].Variables
	assert.Equal(t, 5, len(variables))
	assert.False(t, variables[2].Public)
	assert.Equal(t, uint64(5), *variables[2].MemoryOffset)
	assert.Equal(t, uint64(16), *variables[3].MemoryOffset)
}
