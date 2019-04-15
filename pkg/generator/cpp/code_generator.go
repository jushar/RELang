package cpp

import (
	"github.com/Jusonex/RELang/pkg/generator"
	"github.com/Jusonex/RELang/pkg/model"
	log "github.com/sirupsen/logrus"
)

// CodeGenerator holds information about the C++ code generator
// that accepts the parsed chunk as input and
// generates a C++ header file for it
type CodeGenerator struct {
	// Low-level code emitter
	Emitter      *CodeEmitter
	Transformers []generator.Transformer
}

// TODO: Make configurable for x64 calling convention
const (
	// DefaultFunctionCallingConvention defines the default function calling convention
	DefaultFunctionCallingConvention = "__stdcall"
	// DefaultMethodCallingConvention defines the default method calling convention
	DefaultMethodCallingConvention = "__thiscall"
)

// NewCodeGenerator constructs a new cpp code generator
func NewCodeGenerator(path string) *CodeGenerator {
	return &CodeGenerator{
		Emitter: NewCodeEmitter(path),
		Transformers: []generator.Transformer{
			&TransformerMemoryAddresses{},
			&TransformerVariablePads{},
			&TransformerVirtualFunctionPads{},
			&TransformerCallingConventions{},
			&TransformerTypeTransformations{},
		},
	}
}

// Close disposes associated resources
func (s *CodeGenerator) Close() {
	s.Emitter.Close()
}

// Generate invokes the generator for a specific chunk
func (s *CodeGenerator) Generate(chunk *model.Chunk) {
	// Prepare chunk for code generation
	// by executing all transformers
	for _, transformer := range s.Transformers {
		transformer.Transform(chunk)
	}

	// After preperations are done, emit the code as is
	s.emitCode(chunk)
}

// GetRequiredForwardDeclarations returns the required forward declarations
// for a given chunk
func (s *CodeGenerator) GetRequiredForwardDeclarations(chunk *model.Chunk) []string {
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
func (s *CodeGenerator) emitCode(chunk *model.Chunk) {
	s.Emitter.EmitHeader()
	s.Emitter.EmitIncludeGuard()

	// Emit default includes
	s.Emitter.EmitIncludeStatement("cstdint", false)
	s.Emitter.EmitIncludeStatement("utility", false)
	s.Emitter.EmitLine("", false)

	// Emit forward declarations
	if forwardDecls := s.GetRequiredForwardDeclarations(chunk); len(forwardDecls) > 0 {
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Forward declarations")
		for _, ptrType := range s.GetRequiredForwardDeclarations(chunk) {
			s.Emitter.EmitForwardDeclaration(ptrType)
		}
		s.Emitter.EmitLine("", false)
	}

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
			s.Emitter.EmitClassVariableDeclaration(variable)
		}

		// Emit raw blocks
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Raw blocks")
		for _, rawBlock := range class.RawBlocks {
			s.Emitter.EmitLine(rawBlock.Content, false)
		}

		s.Emitter.EmitClassDeclarationEnd()

		classSize := class.GetSize()
		if classSize == 0 {
			// C++ standard says that size of 0 is not allowed
			classSize = 1
		}

		if class.HasVirtualMembers() {
			// Check if a variable collides with the vtable ptr
			for _, variable := range class.Variables {
				offset := *variable.MemoryOffset
				if offset < uint64(model.POINTER_SIZE) {
					log.WithFields(log.Fields{"class": class.Name, "variable": variable.Name}).Error("member variables collides with vtable pointer")
				}
			}
		}

		s.Emitter.EmitClassSizeAssertion(class.Name, classSize)
	}

	// Emit global functions
	if len(chunk.GlobalFunctions) > 0 {
		s.Emitter.EmitLine("", false)
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Global functions")
		for _, function := range chunk.GlobalFunctions {
			s.Emitter.EmitFunctionDeclaration(function, false)
		}
	}

	// Emit global variables
	if len(chunk.GlobalVariables) > 0 {
		s.Emitter.EmitLine("", false)
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Global variables")
		for _, variable := range chunk.GlobalVariables {
			s.Emitter.EmitGlobalVariableDeclaration(variable)
		}
	}

	// Emit global raw blocks
	if len(chunk.RawBlocks) > 0 {
		s.Emitter.EmitLine("", false)
		s.Emitter.EmitLine("//////////////////////////////", false)
		s.Emitter.EmitLineComment("Raw blocks")
		for _, rawBlock := range chunk.RawBlocks {
			s.Emitter.EmitLine(rawBlock.Content, false)
		}
	}
}
