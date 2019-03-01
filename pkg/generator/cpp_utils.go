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

// Converts an array of cpp variables to a fully typed comma-separated function parameter string
func CppFunctionParametersToString(params []model.Parameter) string {
	return strings.Join(mapCppParameters(params, func(param model.Parameter) string {
		return param.Type + " " + param.Name
	}), ", ")
}

// Converts an array of cpp varaibles to a comma-separated list of parameter names
func CppFunctionParameterNamesToString(params []model.Parameter) string {
	return strings.Join(mapCppParameters(params, func(param model.Parameter) string {
		return param.Name
	}), ", ")
}

// Converts an array of cpp varaibles to a comma-separated list of parameter types
func CppFunctionParameterTypesToString(params []model.Parameter) string {
	return strings.Join(mapCppParameters(params, func(param model.Parameter) string {
		return param.Type
	}), ", ")
}
