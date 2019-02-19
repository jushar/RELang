package generator

import (
	"strconv"

	"github.com/Jusonex/RELang/pkg/parser"
)

type RELangGenerator struct {
	*parser.BaseRELangListener

	Emitter      *CppCodeEmitter
	ContextStack *GeneratorContextStack
	State        struct {
		FunctionName              string
		FunctionParams            []CppVariable
		FunctionReturnType        string
		FunctionCallingConvention string
		MemoryAddress             string
		VariableType              string

		CurrentMemoryAddressOffset uint
		NextAttributeSize          uint8
	}
}

func NewRELangGenerator(outputFilePath string) *RELangGenerator {
	g := new(RELangGenerator)

	g.Emitter = NewCppCodeEmitter(outputFilePath)
	g.ContextStack = NewGeneratorContextStack()

	return g
}

func (s *RELangGenerator) Close() {
	s.Emitter.Close()
}

func (s *RELangGenerator) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	baseClasses := []string{}
	for i, baseClass := range ctx.AllName() {
		if i > 0 {
			baseClasses = append(baseClasses, baseClass.GetText())
		}
	}

	s.Emitter.EmitClassDeclarationStart(ctx.Name(0).GetText(), baseClasses)
	s.ContextStack.Push(CONTEXT_CLASS_DECL)
	s.State.CurrentMemoryAddressOffset = 0
}

func (s *RELangGenerator) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.Emitter.EmitClassDeclarationEnd()
	s.ContextStack.Pop(CONTEXT_CLASS_DECL)
}

func (s *RELangGenerator) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.ContextStack.Push(CONTEXT_FUNCTION_DECL)
}

func (s *RELangGenerator) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.ContextStack.Pop(CONTEXT_FUNCTION_DECL)

	s.Emitter.EmitFunctionDeclaration(ctx.Name().GetText(), s.State.FunctionReturnType, s.State.FunctionParams,
		s.State.MemoryAddress, s.State.FunctionCallingConvention, s.ContextStack.Contains(CONTEXT_CLASS_DECL))
}

func (s *RELangGenerator) EnterFunctionReturnType(ctx *parser.FunctionReturnTypeContext) {
	s.State.FunctionReturnType = ctx.GetText()
}

func (s *RELangGenerator) EnterCallingConvention(ctx *parser.CallingConventionContext) {
	s.State.FunctionCallingConvention = ctx.GetText()
}

func (s *RELangGenerator) ExitFunctionParameter(ctx *parser.FunctionParameterContext) {
	s.State.FunctionParams = append(s.State.FunctionParams, CppVariable{Name: ctx.Name().GetText(), Type: s.State.VariableType})
}

func (s *RELangGenerator) EnterVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	s.ContextStack.Push(CONTEXT_VARIABLE_DECL)
}

func (s *RELangGenerator) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	s.ContextStack.Pop(CONTEXT_VARIABLE_DECL)

	s.Emitter.EmitVariableDeclaration(CppVariable{Name: ctx.Name().GetText(), Type: s.State.VariableType},
		s.State.CurrentMemoryAddressOffset, s.ContextStack.Contains(CONTEXT_CLASS_DECL))

	s.State.CurrentMemoryAddressOffset = s.State.CurrentMemoryAddressOffset + uint(s.State.NextAttributeSize)
}

func (s *RELangGenerator) EnterMemoryAddress(ctx *parser.MemoryAddressContext) {
	s.State.MemoryAddress = ctx.HexInteger().GetText()
	// TODO: Validate memory address

	if s.ContextStack.Contains(CONTEXT_CLASS_DECL) && s.ContextStack.Top() == CONTEXT_VARIABLE_DECL {
		memoryOffset, err := strconv.ParseUint(s.State.MemoryAddress[2:len(s.State.MemoryAddress)], 10, 32)
		// TODO: Handle error properly
		if err != nil {
			panic(err)
		}
		s.State.CurrentMemoryAddressOffset = uint(memoryOffset) // TODO: Make sure memoryOffset is not smaller than CurrentMemoryAddressOffset
	}
}

func (s *RELangGenerator) EnterPointer(ctx *parser.PointerContext) {
	s.State.VariableType = ctx.Name().GetText() + "*"

	if s.ContextStack.Contains(CONTEXT_CLASS_DECL) && s.ContextStack.Top() == CONTEXT_VARIABLE_DECL {
		s.State.NextAttributeSize = POINTER_SIZE
	}
}

func (s *RELangGenerator) EnterPrimitiveType(ctx *parser.PrimitiveTypeContext) {
	s.State.VariableType = ctx.GetText()

	if s.ContextStack.Contains(CONTEXT_CLASS_DECL) && s.ContextStack.Top() == CONTEXT_VARIABLE_DECL {
		s.State.NextAttributeSize = TYPES_SIZE[s.State.VariableType]
	}
}
