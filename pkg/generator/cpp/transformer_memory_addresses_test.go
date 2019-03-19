package cpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesMemoryAddresses(t *testing.T) {
	gen := NewCodeGenerator(outputFile)
	defer gen.Close()
	chunk := prepareTestChunk()

	transformer := &TransformerMemoryAddresses{}
	transformer.Transform(chunk)

	variables := chunk.Classes[0].Variables
	assert.Equal(t, uint64(0), *variables[0].MemoryOffset)
	assert.Equal(t, uint64(4), *variables[1].MemoryOffset)
	assert.Equal(t, uint64(16), *variables[2].MemoryOffset)
	assert.Equal(t, uint64(24), *variables[3].MemoryOffset)

	// TODO: Test virtual functions too
}
