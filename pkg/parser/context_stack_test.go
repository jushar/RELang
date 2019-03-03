package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushPopContextStack(t *testing.T) {
	stack := NewGeneratorContextStack()

	stack.Push(CONTEXT_CLASS_DECL)
	assert.Equal(t, 1, stack.Size())

	stack.Push(CONTEXT_CLASS_DECL)
	stack.Push(CONTEXT_VARIABLE_DECL)
	assert.Equal(t, 3, stack.Size())

	stack.Pop(CONTEXT_VARIABLE_DECL)
	assert.Equal(t, 2, stack.Size())
}

func TestStackTopElement(t *testing.T) {
	stack := NewGeneratorContextStack()

	stack.Push(CONTEXT_CLASS_DECL)
	assert.Equal(t, CONTEXT_CLASS_DECL, stack.Top())

	stack.Push(CONTEXT_VARIABLE_DECL)
	assert.Equal(t, CONTEXT_VARIABLE_DECL, stack.Top())

	stack.Pop(CONTEXT_VARIABLE_DECL)
	assert.Equal(t, CONTEXT_CLASS_DECL, stack.Top())
}

func TestStackContains(t *testing.T) {
	stack := NewGeneratorContextStack()
	assert.False(t, stack.Contains(CONTEXT_CLASS_DECL))

	stack.Push(CONTEXT_CLASS_DECL)
	assert.True(t, stack.Contains(CONTEXT_CLASS_DECL))

	stack.Push(CONTEXT_VARIABLE_DECL)
	assert.True(t, stack.Contains(CONTEXT_CLASS_DECL))
	assert.True(t, stack.Contains(CONTEXT_VARIABLE_DECL))
}
