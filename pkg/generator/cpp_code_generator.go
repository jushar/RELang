package generator

import "github.com/Jusonex/RELang/pkg/model"

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
	s.EmitCode(chunk)
}

func (s *CppCodeGenerator) CreateMemoryAddresses(chunk *model.Chunk) {
	// TODO
}

func (s *CppCodeGenerator) CreateVariablePads(chunk *model.Chunk) {
	// TODO
}

func (s *CppCodeGenerator) CreateVirtualFunctionPads(chunk *model.Chunk) {
	// TODO
}

func (s *CppCodeGenerator) EmitCode(chunk *model.Chunk) {
	s.Emitter.EmitHeader()

	// Emit classes
	for _, class := range chunk.Classes {
		s.Emitter.EmitClassDeclarationStart(class.Name, class.BaseClasses)
		s.Emitter.EmitPublicBlock()

		// Emit virtual functions
		for _, function := range class.VirtualFunctions {
			s.Emitter.EmitVirtualFunctionDeclaration(function.Name, function.ReturnType, function.Params, *function.MemoryAddress, function.CallingConvention)
		}

		// Emit functions
		for _, function := range class.Functions {
			s.Emitter.EmitFunctionDeclaration(function.Name, function.ReturnType, function.Params, *function.MemoryAddress, function.CallingConvention, true)
		}

		// Emit variables
		for _, variable := range class.Variables {
			s.Emitter.EmitVariableDeclaration(variable.Name, variable.Type, *variable.MemoryOffset)
		}

		s.Emitter.EmitClassDeclarationEnd()
	}
}
