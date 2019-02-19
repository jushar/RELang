package generator

import (
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
		VariableName              string
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
	s.State.FunctionParams = append(s.State.FunctionParams, CppVariable{Name: ctx.Name().GetText(), Type: s.State.VariableName})
}

func (s *RELangGenerator) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	s.Emitter.EmitVariableDeclaration(CppVariable{Name: ctx.Name().GetText(), Type: s.State.VariableName},
		s.State.MemoryAddress, s.ContextStack.Contains(CONTEXT_CLASS_DECL))
}

func (s *RELangGenerator) EnterMemoryAddress(ctx *parser.MemoryAddressContext) {
	s.State.MemoryAddress = ctx.HexInteger().GetText()
}

func (s *RELangGenerator) EnterPointer(ctx *parser.PointerContext) {
	s.State.VariableName = ctx.Name().GetText() + "*"
}

func (s *RELangGenerator) EnterPrimitiveType(ctx *parser.PrimitiveTypeContext) {
	s.State.VariableName = ctx.GetText()
}
