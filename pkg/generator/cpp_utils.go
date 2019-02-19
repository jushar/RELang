package generator

import "strings"

func mapCppVariables(vs []CppVariable, f func(CppVariable) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Converts an array of cpp variables to a fully typed comma-separated function parameter string
func CppFunctionParametersToString(params []CppVariable) string {
	return strings.Join(mapCppVariables(params, func(param CppVariable) string {
		return param.Type + " " + param.Name
	}), ", ")
}

// Converts an array of cpp varaibles to a comma-separated list of parameter names
func CppFunctionParameterNamesToString(params []CppVariable) string {
	return strings.Join(mapCppVariables(params, func(param CppVariable) string {
		return param.Name
	}), ", ")
}

// Converts an array of cpp varaibles to a comma-separated list of parameter types
func CppFunctionParameterTypesToString(params []CppVariable) string {
	return strings.Join(mapCppVariables(params, func(param CppVariable) string {
		return param.Type
	}), ", ")
}
