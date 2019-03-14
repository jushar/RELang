package generator

import (
	"github.com/Jusonex/RELang/pkg/model"
)

// Holds information about the C++ code generator
// that accepts the parsed chunk as input and
// generates a C++ header file for it
type CppCodeGenerator struct {
	// Low-level code emitter
	Emitter *CppCodeEmitter
}

const (
	// TODO: Make configurable for x64 calling convention
	DEFAULT_FUNCTION_CALLING_CONVENTION = "__stdcall"
	DEFAULT_METHOD_CALLING_CONVENTION   = "__thiscall"
)

// Constructs a new cpp code generator
func NewCppCodeGenerator(path string) *CppCodeGenerator {
	return &CppCodeGenerator{
		Emitter: NewCppCodeEmitter(path),
	}
}

// Disposes associated resources
func (s *CppCodeGenerator) Close() {
	s.Emitter.Close()
}

// Invokes the generator for a specific chunk
func (s *CppCodeGenerator) Generate(chunk *model.Chunk) {
	// Prepare chunk for code generation
	// e.g. add missing fields
	s.CreateMemoryAddresses(chunk)
	s.CreateVariablePads(chunk)
	s.CreateVirtualFunctionPads(chunk)
	s.CreateCallingConventions(chunk)
	s.ApplyTypeTransformations(chunk)

	// After preperations are done, emit the code as is
	s.EmitCode(chunk)
}

// Fills missing attributes for a chunk in subsequent manner
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

// Creates padding attributes for field gaps for the specified chunk
func (s *CppCodeGenerator) CreateVariablePads(chunk *model.Chunk) {
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

// Creates pads for virtual functions
func (s *CppCodeGenerator) CreateVirtualFunctionPads(chunk *model.Chunk) {
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

// Adds default calling conventions to the specified chunk
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

// ApplyTypeTransformations applies available type information to functions and variables in the chunk
func (s *CppCodeGenerator) ApplyTypeTransformations(chunk *model.Chunk) {
	ApplyForAllFunctions(chunk, func(function *model.Function) {
		ApplyFunctionTypeTransformations(function)
	})
	ApplyForAllVariables(chunk, func(variable *model.Variable) {
		ApplyVariableTypeTransformations(variable)
	})
}

func (s *CppCodeGenerator) GetRequiredForwardDeclarations(chunk *model.Chunk) []string {
	var forwardDeclarations []string // TODO: Use a Set for this

	for _, class := range chunk.Classes {
		for _, variable := range class.Variables {
			forwardDeclarations = append(forwardDeclarations, variable.GetRequiredForwardDeclarations()...)
		}

		for _, function := range class.Functions {
			forwardDeclarations = append(forwardDeclarations, function.GetRequiredForwardDeclarations()...)
		}

		for _, function := range class.VirtualFunctions {
			forwardDeclarations = append(forwardDeclarations, function.GetRequiredForwardDeclarations()...)
		}
	}

	for _, variable := range chunk.GlobalVariables {
		forwardDeclarations = append(forwardDeclarations, variable.GetRequiredForwardDeclarations()...)
	}

	for _, function := range chunk.GlobalFunctions {
		forwardDeclarations = append(forwardDeclarations, function.GetRequiredForwardDeclarations()...)
	}

	return forwardDeclarations
}

// Emits the code for a specific chunk.
// Assumes that the input is well-filled and prepared.
func (s *CppCodeGenerator) EmitCode(chunk *model.Chunk) {
	s.Emitter.EmitHeader()

	// Emit forward declarations
	s.Emitter.EmitLine("//////////////////////////////", false)
	s.Emitter.EmitLineComment("Forward declarations")
	for _, ptrType := range s.GetRequiredForwardDeclarations(chunk) {
		s.Emitter.EmitForwardDeclaration(ptrType)
	}
	s.Emitter.EmitLine("", false)

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

		// Emit raw blocks
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Raw blocks")
		for _, rawBlock := range class.RawBlocks {
			s.Emitter.EmitLine(rawBlock.Content, false)
		}

		s.Emitter.EmitClassDeclarationEnd()
		s.Emitter.EmitClassSizeAssertion(class.Name, class.GetSize())
	}

	// Emit global functions
	s.Emitter.EmitLine("", false)
	s.Emitter.EmitLine("//////////////////////////////", false)
	s.Emitter.EmitLineComment("Global functions")
	for _, function := range chunk.GlobalFunctions {
		s.Emitter.EmitFunctionDeclaration(function, false)
	}

	// Emit global raw blocks
	s.Emitter.EmitLine("//////////////////////////////", false)
	s.Emitter.EmitLineComment("Raw blocks")
	for _, rawBlock := range chunk.RawBlocks {
		s.Emitter.EmitLine(rawBlock.Content, false)
	}
}
