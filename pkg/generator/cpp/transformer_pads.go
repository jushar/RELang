package cpp

import "github.com/Jusonex/RELang/pkg/model"

// TransformerVariablePads inserts variable pads
type TransformerVariablePads struct {
}

// Transform creates padding attributes for field gaps for the specified chunk
func (s *TransformerVariablePads) Transform(chunk *model.Chunk) {
	for _, class := range chunk.Classes {
		var newVariables []*model.Variable

		currentAddr := uint64(0)
		for _, variable := range class.Variables {
			size, err := model.GetInbuiltTypeSize(variable.Type)
			if err != nil { // TODO: Support custom types
				panic(err)
			}

			// If the stored memory offset doesn't match the theoretical,
			// we have to insert a pad
			if diff := *variable.MemoryOffset - currentAddr; diff > 0 {
				newVariables = append(newVariables, model.NewVariablePad(currentAddr, diff))
				currentAddr = currentAddr + diff
			}

			// Add actual variable
			newVariables = append(newVariables, variable)
			currentAddr = currentAddr + uint64(size)
		}

		class.Variables = newVariables
	}
}

// TransformerVirtualFunctionPads inserts virtual function pads
type TransformerVirtualFunctionPads struct {
}

// Transform creates pads for virtual functions
func (s *TransformerVirtualFunctionPads) Transform(chunk *model.Chunk) {
	for _, class := range chunk.Classes {
		var newFunctions []*model.Function

		currentAddr := uint64(0)
		for _, function := range class.VirtualFunctions {
			// If the stored memory offset doesn't match the theoretical,
			// we have to insert a pad
			if diff := *function.MemoryAddress - currentAddr; diff > 0 {
				for i := uint64(0); i < diff/uint64(model.POINTER_SIZE); i = i + 1 {
					newFunctions = append(newFunctions, model.NewFunctionPad(currentAddr))
					currentAddr = currentAddr + uint64(model.POINTER_SIZE)
				}
				// TODO: Throw error if not multiple of POINTER_SIZE or interleaving
			}

			// Add actual variable
			newFunctions = append(newFunctions, function)
			currentAddr = currentAddr + uint64(model.POINTER_SIZE)
		}

		class.VirtualFunctions = newFunctions
	}
}
