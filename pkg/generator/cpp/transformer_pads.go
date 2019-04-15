package cpp

import (
	"github.com/Jusonex/RELang/pkg/model"
	log "github.com/sirupsen/logrus"
)

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

		// Remove vtable ptr from variable list
		if class.HasVirtualMembers() && len(newVariables) > 0 {
			variable := newVariables[0]
			if *variable.MemoryOffset != 0 {
				log.WithFields(log.Fields{"class": class.Name, "variable": variable.Name, "offset": *variable.MemoryOffset}).Fatal(
					"internal compiler error: class variables do not start at 0")
			}

			if variable.Size < uint64(model.POINTER_SIZE) {
				log.WithFields(log.Fields{"class": class.Name, "variable": variable.Name, "size": variable.Size}).Fatal(
					"internal compiler error: first member size is smaller than ptr size")
			}

			variable.Size = variable.Size - uint64(model.POINTER_SIZE)

			// If size is zero, remove entirely
			if variable.Size == 0 {
				newVariables = newVariables[1 : len(newVariables)-1]
			}
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
