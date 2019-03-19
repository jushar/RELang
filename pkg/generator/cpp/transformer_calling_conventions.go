package cpp

import (
	"github.com/Jusonex/RELang/pkg/generator"
	"github.com/Jusonex/RELang/pkg/model"
)

// TransformerCallingConventions inserts missing calling conventions
type TransformerCallingConventions struct {
}

// Transform adds default calling conventions to the specified chunk
func (s *TransformerCallingConventions) Transform(chunk *model.Chunk) {
	generator.ApplyForAllFunctions(chunk, func(function *model.Function, method bool, virtual bool) {
		if function.CallingConvention == "" {
			if method {
				function.CallingConvention = DefaultMethodCallingConvention
			} else {
				function.CallingConvention = DefaultFunctionCallingConvention
			}
		}
	})
}
