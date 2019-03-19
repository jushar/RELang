package cpp

import (
	"github.com/Jusonex/RELang/pkg/generator"
	"github.com/Jusonex/RELang/pkg/model"
)

// TransformerTypeTransformations transforms RELang types to C++ types
type TransformerTypeTransformations struct {
}

// Transform applies available type information to functions and variables in the chunk
func (s *TransformerTypeTransformations) Transform(chunk *model.Chunk) {
	generator.ApplyForAllFunctions(chunk, func(function *model.Function, method bool, virtual bool) {
		ApplyFunctionTypeTransformations(function)
	})
	generator.ApplyForAllVariables(chunk, func(variable *model.Variable) {
		ApplyVariableTypeTransformations(variable)
	})
}
