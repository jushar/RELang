package parser

type GeneratorContext int32

type GeneratorContextStack struct {
	stack []GeneratorContext
}

const (
	CONTEXT_CLASS_DECL    GeneratorContext = 0
	CONTEXT_VARIABLE_DECL GeneratorContext = 1
	CONTEXT_FUNCTION_DECL GeneratorContext = 2
)

// Creates a new generator context stack
func NewGeneratorContextStack() *GeneratorContextStack {
	return new(GeneratorContextStack)
}

// Pushes a context to the top of the stack
func (s *GeneratorContextStack) Push(context GeneratorContext) {
	s.stack = append(s.stack, context)
}

// Removes the top element of the stack that is a specific context
// panics, if not
func (s *GeneratorContextStack) Pop(context GeneratorContext) {
	if s.Top() != context {
		panic("Stack corrupted. Invalid top element, because of broken unwinding")
	}

	s.stack = s.stack[0 : len(s.stack)-1]
}

// Returns the top element of the stack
func (s *GeneratorContextStack) Top() GeneratorContext {
	l := len(s.stack)
	if l == 0 {
		panic("Stack empty")
	}
	return s.stack[l-1]
}

// Checks whether or not the stack contains a specific context
func (s *GeneratorContextStack) Contains(context GeneratorContext) bool {
	for _, v := range s.stack {
		if v == context {
			return true
		}
	}

	return false
}
