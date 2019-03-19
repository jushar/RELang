package cpp

import (
	"strings"

	"github.com/Jusonex/RELang/pkg/model"
)

func mapParameters(vs []model.Parameter, f func(model.Parameter) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// FunctionParametersToString converts an array of cpp variables to a fully typed comma-separated function parameter string
func FunctionParametersToString(params []model.Parameter) string {
	return strings.Join(mapParameters(params, func(param model.Parameter) string {
		return param.Type + " " + param.Name
	}), ", ")
}

// FunctionParameterNamesToString converts an array of cpp varaibles to a comma-separated list of parameter names
func FunctionParameterNamesToString(params []model.Parameter, perfectForward bool) string {
	return strings.Join(mapParameters(params, func(param model.Parameter) string {
		if perfectForward {
			return "std::forward<" + param.Type + ">(" + param.Name + ")"
		}

		return param.Name
	}), ", ")
}

// FunctionParameterTypesToString converts an array of cpp varaibles to a comma-separated list of parameter types
func FunctionParameterTypesToString(params []model.Parameter) string {
	return strings.Join(mapParameters(params, func(param model.Parameter) string {
		return param.Type
	}), ", ")
}
