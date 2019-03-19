package cpp

import "github.com/Jusonex/RELang/pkg/model"

// TransformerMemoryAddresses fills in implicit memory addresses/offsets
type TransformerMemoryAddresses struct {
}

// Transform fills missing attributes for a chunk in subsequent manner
func (s *TransformerMemoryAddresses) Transform(chunk *model.Chunk) {
	for _, class := range chunk.Classes {
		currentAddr := uint64(0)

		// Fill memory offsets for class variables
		for _, variable := range class.Variables {
			if variable.MemoryOffset != nil {
				currentAddr = *variable.MemoryOffset
			} else {
				addr := currentAddr
				variable.MemoryOffset = &addr
			}

			size, err := model.GetInbuiltTypeSize(variable.Type)
			if err != nil { // TODO: Support custom types
				panic(err)
			}
			currentAddr = currentAddr + uint64(size)

			variable.Size = uint64(size)
		}

		// Fill memory offsets for virtual methods
		currentAddr = uint64(0)
		for _, function := range class.VirtualFunctions {
			if function.MemoryAddress != nil {
				currentAddr = *function.MemoryAddress
			} else {
				addr := currentAddr
				function.MemoryAddress = &addr
			}

			currentAddr = currentAddr + uint64(model.POINTER_SIZE)
		}
	}
}
