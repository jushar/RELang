package generator

import (
	"strings"

	"github.com/Jusonex/RELang/pkg/model"
)

func mapCppParameters(vs []model.Parameter, f func(model.Parameter) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// CppFunctionParametersToString converts an array of cpp variables to a fully typed comma-separated function parameter string
func CppFunctionParametersToString(params []model.Parameter) string {
	return strings.Join(mapCppParameters(params, func(param model.Parameter) string {
		return param.Type + " " + param.Name
	}), ", ")
}

// CppFunctionParameterNamesToString converts an array of cpp varaibles to a comma-separated list of parameter names
func CppFunctionParameterNamesToString(params []model.Parameter, perfectForward bool) string {
	return strings.Join(mapCppParameters(params, func(param model.Parameter) string {
		if perfectForward {
			return "std::forward<" + param.Type + ">(" + param.Name + ")"
		}

		return param.Name
	}), ", ")
}

// CppFunctionParameterTypesToString converts an array of cpp varaibles to a comma-separated list of parameter types
func CppFunctionParameterTypesToString(params []model.Parameter) string {
	return strings.Join(mapCppParameters(params, func(param model.Parameter) string {
		return param.Type
	}), ", ")
}

// ApplyForAllFunctions applies a functor to all types of functions (global/class, virtual/non-virtual).
// Functor parameters: function, isMethod, isVirtual
func ApplyForAllFunctions(chunk *model.Chunk, f func(*model.Function, bool, bool)) {
	for _, class := range chunk.Classes {
		for _, function := range class.Functions {
			f(function, true, false)
		}

		for _, function := range class.VirtualFunctions {
			f(function, true, true)
		}
	}

	for _, function := range chunk.GlobalFunctions {
		f(function, false, false)
	}
}

// ApplyForAllVariables applies a functor to all types of variables (global/class)
func ApplyForAllVariables(chunk *model.Chunk, f func(*model.Variable)) {
	for _, class := range chunk.Classes {
		for _, variable := range class.Variables {
			f(variable)
		}
	}

	for _, variable := range chunk.GlobalVariables {
		f(variable)
	}
}
