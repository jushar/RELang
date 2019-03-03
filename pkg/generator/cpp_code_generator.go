package generator

import (
	"github.com/Jusonex/RELang/pkg/model"
)

type CppCodeGenerator struct {
	Emitter *CppCodeEmitter
}

const (
	// TODO: Make configurable for x64 calling convention
	DEFAULT_FUNCTION_CALLING_CONVENTION = "__stdcall"
	DEFAULT_METHOD_CALLING_CONVENTION   = "__thiscall"
)

func NewCppCodeGenerator(path string) *CppCodeGenerator {
	return &CppCodeGenerator{
		Emitter: NewCppCodeEmitter(path),
	}
}

func (s *CppCodeGenerator) Close() {
	s.Emitter.Close()
}

func (s *CppCodeGenerator) Generate(chunk *model.Chunk) {
	s.CreateMemoryAddresses(chunk)
	s.CreateVariablePads(chunk)
	s.CreateVirtualFunctionPads(chunk)
	s.CreateCallingConventions(chunk)
	s.EmitCode(chunk)
}

func (s *CppCodeGenerator) CreateMemoryAddresses(chunk *model.Chunk) {
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

			size, err := GetInbuiltTypeSize(variable.Type)
			if err != nil { // TODO: Support custom types
				panic(err)
			}
			currentAddr = currentAddr + uint64(size)
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

			currentAddr = currentAddr + uint64(POINTER_SIZE)
		}
	}
}

func (s *CppCodeGenerator) CreateVariablePads(chunk *model.Chunk) {
	for _, class := range chunk.Classes {
		var newVariables []*model.Variable

		currentAddr := uint64(0)
		for _, variable := range class.Variables {
			size, err := GetInbuiltTypeSize(variable.Type)
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

func (s *CppCodeGenerator) CreateVirtualFunctionPads(chunk *model.Chunk) {
	for _, class := range chunk.Classes {
		var newFunctions []*model.Function

		currentAddr := uint64(0)
		for _, function := range class.VirtualFunctions {
			// If the stored memory offset doesn't match the theoretical,
			// we have to insert a pad
			if diff := *function.MemoryAddress - currentAddr; diff > 0 {
				for i := uint64(0); i < diff/uint64(POINTER_SIZE); i = i + 1 {
					newFunctions = append(newFunctions, model.NewFunctionPad(currentAddr))
					currentAddr = currentAddr + uint64(POINTER_SIZE)
				}
				// TODO: Throw error if not multiple of POINTER_SIZE or interleaving
			}

			// Add actual variable
			newFunctions = append(newFunctions, function)
			currentAddr = currentAddr + uint64(POINTER_SIZE)
		}

		class.VirtualFunctions = newFunctions
	}
}

func (s *CppCodeGenerator) CreateCallingConventions(chunk *model.Chunk) {
	for _, class := range chunk.Classes {
		for _, function := range class.Functions {
			if function.CallingConvention == "" {
				function.CallingConvention = DEFAULT_METHOD_CALLING_CONVENTION
			}
		}

		for _, function := range class.VirtualFunctions {
			if function.CallingConvention == "" {
				function.CallingConvention = DEFAULT_METHOD_CALLING_CONVENTION
			}
		}
	}

	for _, function := range chunk.GlobalFunctions {
		if function.CallingConvention == "" {
			function.CallingConvention = DEFAULT_FUNCTION_CALLING_CONVENTION
		}
	}
}

func (s *CppCodeGenerator) EmitCode(chunk *model.Chunk) {
	s.Emitter.EmitHeader()

	// Emit classes
	for _, class := range chunk.Classes {
		s.Emitter.EmitClassDeclarationStart(class.Name, class.BaseClasses)

		// Emit virtual functions
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Virtual functions")
		for _, function := range class.VirtualFunctions {
			s.Emitter.EmitVirtualFunctionDeclaration(function)
		}

		// Emit functions
		s.Emitter.EmitLine("", false)
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Functions")
		for _, function := range class.Functions {
			s.Emitter.EmitFunctionDeclaration(function, true)
		}

		// Emit variables
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Variables")
		for _, variable := range class.Variables {
			s.Emitter.EmitVariableDeclaration(variable)
		}

		s.Emitter.EmitClassDeclarationEnd()
	}

	// Emit global functions
	s.Emitter.EmitLine("", false)
	for _, function := range chunk.GlobalFunctions {
		s.Emitter.EmitLineComment("Global functions")

		s.Emitter.EmitFunctionDeclaration(function, false)
	}
}
