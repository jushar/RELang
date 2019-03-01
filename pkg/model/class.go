package model

type Class struct {
	Name             string
	BaseClasses      []string
	Variables        []*Variable
	Functions        []*Function
	VirtualFunctions []*Function
	StaticFunctions  []*Function
}

func NewClass(name string) *Class {
	return &Class{Name: name}
}

func (s *Class) AddBaseClass(name string) {
	s.BaseClasses = append(s.BaseClasses, name)
}

func (s *Class) AddFunction(function *Function) {
	switch function.Modifier {
	case "virtual":
		s.VirtualFunctions = append(s.VirtualFunctions, function)
		break
	case "static":
		s.StaticFunctions = append(s.StaticFunctions, function)
		break
	default:
		s.Functions = append(s.Functions, function)
	}
}
