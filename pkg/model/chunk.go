package model

type Chunk struct {
	Classes         []*Class
	GlobalFunctions []*Function
	GlobalVariables []*Variable
}

func NewChunk() *Chunk {
	return &Chunk{}
}

func (s *Chunk) AddClass(class *Class) {
	s.Classes = append(s.Classes, class)
}

func (s *Chunk) AddFunction(function *Function) {
	s.GlobalFunctions = append(s.GlobalFunctions, function)
}

func (s *Chunk) AddVariable(variable *Variable) {
	s.GlobalVariables = append(s.GlobalVariables, variable)
}
