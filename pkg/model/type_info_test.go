package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInbuiltTypeSize(t *testing.T) {
	// valid primitive type
	size, err := GetInbuiltTypeSize("int16")
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), size)

	// valid pointer type
	size, err = GetInbuiltTypeSize("bla*")
	assert.Nil(t, err)
	assert.Equal(t, POINTER_SIZE, size)

	// invalid type
	size, err = GetInbuiltTypeSize("abc")
	assert.NotNil(t, err)
}
