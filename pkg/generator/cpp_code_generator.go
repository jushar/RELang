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
	// TODO
}

func (s *CppCodeGenerator) CreateVirtualFunctionPads(chunk *model.Chunk) {
	// TODO
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
		s.Emitter.EmitPublicBlock()

		// Emit virtual functions
		s.Emitter.EmitLineComment("Virtual functions")
		for _, function := range class.VirtualFunctions {
			s.Emitter.EmitVirtualFunctionDeclaration(function.Name, function.ReturnType, function.Params, *function.MemoryAddress, function.CallingConvention)
		}

		// Emit functions
		s.Emitter.EmitLine("", false)
		s.Emitter.EmitLineComment("Functions")
		for _, function := range class.Functions {
			s.Emitter.EmitFunctionDeclaration(function.Name, function.ReturnType, function.Params, *function.MemoryAddress, function.CallingConvention, true)
		}

		// Emit variables
		s.Emitter.EmitLineComment("Variables")
		for _, variable := range class.Variables {
			s.Emitter.EmitVariableDeclaration(variable.Name, variable.Type, *variable.MemoryOffset)
		}

		s.Emitter.EmitClassDeclarationEnd()
	}
}
