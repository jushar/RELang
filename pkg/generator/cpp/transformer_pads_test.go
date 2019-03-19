package cpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesVariablePads(t *testing.T) {
	gen := NewCodeGenerator(outputFile)
	defer gen.Close()
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
