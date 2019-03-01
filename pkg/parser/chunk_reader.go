package parser

import (
	"strconv"

	"github.com/Jusonex/RELang/pkg/model"
)

type ChunkReader struct {
	*BaseRELangListener

	ContextStack *GeneratorContextStack

	State struct {
		Chunk    *model.Chunk
		Class    *model.Class
		Function *model.Function
		Variable *model.Variable

		VariableType string
	}
}

func NewChunkReader(outputFilePath string) *ChunkReader {
	s := new(ChunkReader)

	s.State.Chunk = model.NewChunk()
	s.ContextStack = NewGeneratorContextStack()

	return s
}

func (s *ChunkReader) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	s.State.Class = model.NewClass(ctx.Name(0).GetText())

	for i, baseClass := range ctx.AllName() {
		if i > 0 {
			s.State.Class.AddBaseClass(baseClass.GetText())
		}
	}

	s.ContextStack.Push(CONTEXT_CLASS_DECL)
}

func (s *ChunkReader) ExitClassDeclaration(ctx *ClassDeclarationContext) {
	s.State.Chunk.AddClass(s.State.Class)

	s.ContextStack.Pop(CONTEXT_CLASS_DECL)
	s.State.Class = nil
}

func (s *ChunkReader) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {
	s.State.Function = model.NewFunction()
	s.ContextStack.Push(CONTEXT_FUNCTION_DECL)
}

func (s *ChunkReader) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {
	s.ContextStack.Pop(CONTEXT_FUNCTION_DECL)

	inClass := s.ContextStack.Contains(CONTEXT_CLASS_DECL)
	if inClass {
		s.State.Class.AddFunction(s.State.Function)
	} else {
		s.State.Chunk.AddFunction(s.State.Function)
	}

	s.State.Function = nil
}

func (s *ChunkReader) EnterFunctionModifier(ctx *FunctionModifierContext) {
	if s.ContextStack.Top() != CONTEXT_FUNCTION_DECL {
		// TODO: Log error
	}

	s.State.Function.Modifier = ctx.GetText()
}

func (s *ChunkReader) EnterFunctionReturnType(ctx *FunctionReturnTypeContext) {
	if s.ContextStack.Top() != CONTEXT_FUNCTION_DECL {
		// TODO: Log error
	}

	s.State.Function.ReturnType = ctx.GetText()
}

func (s *ChunkReader) EnterCallingConvention(ctx *CallingConventionContext) {
	if s.ContextStack.Top() != CONTEXT_FUNCTION_DECL {
		// TODO: Log error
	}

	s.State.Function.CallingConvention = ctx.GetText()
}

func (s *ChunkReader) ExitFunctionParameter(ctx *FunctionParameterContext) {
	s.State.Function.AddParameter(ctx.Name().GetText(), s.State.VariableType)
}

func (s *ChunkReader) EnterVariableDeclaration(ctx *VariableDeclarationContext) {
	s.State.Variable = model.NewVariable()
	s.ContextStack.Push(CONTEXT_VARIABLE_DECL)
}

func (s *ChunkReader) ExitVariableDeclaration(ctx *VariableDeclarationContext) {
	s.ContextStack.Pop(CONTEXT_VARIABLE_DECL)
	s.State.Variable = nil
}

func (s *ChunkReader) EnterMemoryAddress(ctx *MemoryAddressContext) {
	// Memory address might not be supplied, so just ignore
	// TODO: Empty field can be removed from the grammer (=> ?-notation)
	if ctx.HexInteger() == nil {
		return
	}

	addressStr := ctx.HexInteger().GetText()
	memoryAddress, err := strconv.ParseUint(addressStr[2:], 16, 64)
	// TODO: Handle error
	if err != nil {
		panic(err)
	}

	if s.ContextStack.Top() == CONTEXT_FUNCTION_DECL {
		s.State.Function.MemoryAddress = &memoryAddress
	} else if s.ContextStack.Top() == CONTEXT_VARIABLE_DECL {
		s.State.Variable.MemoryOffset = &memoryAddress
	} else {
		// TODO: Log error
	}
}

func (s *ChunkReader) EnterPointer(ctx *PointerContext) {
	s.State.VariableType = ctx.Name().GetText() + "*"
}

func (s *ChunkReader) EnterPrimitiveType(ctx *PrimitiveTypeContext) {
	s.State.VariableType = ctx.GetText()
}
