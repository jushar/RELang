package generator

import (
	"github.com/Jusonex/RELang/pkg/model"
)

var inbuiltTypes = map[string]string{
	"int8":    "std::int8_t",
	"int16":   "std::int16_t",
	"int32":   "std::int32_t",
	"int64":   "std::int64_t",
	"uint8":   "std::uint8_t",
	"uint16":  "std::uint16_t",
	"uint32":  "std::uint32_t",
	"uint64":  "std::uint64_t",
	"float32": "float",
	"float64": "double",
}

// TransformInbuiltType transforms a RELang type to a C++ type
// or returns the type itself if no transformation is available
func TransformInbuiltType(typeName string) string {
	if newType, ok := inbuiltTypes[typeName]; ok {
		return newType
	}

	return typeName
}

// ApplyVariableTypeTransformations applies available type transformations to a type
func ApplyVariableTypeTransformations(variable *model.Variable) {
	variable.Type = TransformInbuiltType(variable.Type)
}

// ApplyFunctionTypeTransformations applies available type transformations to a type
func ApplyFunctionTypeTransformations(function *model.Function) {
	function.ReturnType = TransformInbuiltType(function.ReturnType)

	for idx, param := range function.Params {
		function.Params[idx].Type = TransformInbuiltType(param.Type)
	}
}
