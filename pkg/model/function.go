package model

type Parameter struct {
	Name string
	Type string
}

type FunctionModifier = string
type CallingConvention = string

type Function struct {
	Name              string
	ReturnType        string
	Modifier          FunctionModifier
	Params            []Parameter
	CallingConvention CallingConvention
	MemoryAddress     *uint64
}

func NewFunction() *Function {
	return &Function{}
}

func (s *Function) AddParameter(name string, parameterType string) {
	s.Params = append(s.Params, Parameter{Name: name, Type: parameterType})
}
