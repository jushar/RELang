package model

// Represents a file that consists of classes, global functions and -variables
type Chunk struct {
	Classes         []*Class
	GlobalFunctions []*Function
	GlobalVariables []*Variable
	RawBlocks       []*RawBlock
}

// Constructs a new chunk
func NewChunk() *Chunk {
	return &Chunk{}
}

// Adds a class to the chunk
func (s *Chunk) AddClass(class *Class) {
	s.Classes = append(s.Classes, class)
}

// Adds a global function to the chunk
func (s *Chunk) AddFunction(function *Function) {
	s.GlobalFunctions = append(s.GlobalFunctions, function)
}

// Adds a global variable to the chunk
func (s *Chunk) AddVariable(variable *Variable) {
	s.GlobalVariables = append(s.GlobalVariables, variable)
}

// Adds a raw block to the chunk
func (s *Chunk) AddRawBlock(rawBlock *RawBlock) {
	s.RawBlocks = append(s.RawBlocks, rawBlock)
}
