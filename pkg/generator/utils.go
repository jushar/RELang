package generator

import "github.com/Jusonex/RELang/pkg/model"

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
