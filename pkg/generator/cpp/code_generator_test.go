package cpp

import (
	"os"
	"testing"

	"github.com/Jusonex/RELang/pkg/model"
	"github.com/stretchr/testify/assert"
)

const outputFile = "../../../test/testoutput.h"

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
	gen := NewCodeGenerator(outputFile)
	defer gen.Close()

	assert.NotNil(t, gen.Emitter)
	_, err := os.Stat(outputFile)
	assert.Nil(t, err)
}

func TestGeneratesChunk(t *testing.T) {
	gen := NewCodeGenerator(outputFile)
	chunk := prepareTestChunk()

	gen.Generate(chunk)
	gen.Close()

	stat, err := os.Stat(outputFile)
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 100)
}
