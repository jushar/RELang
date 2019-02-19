package generator

import (
	"fmt"

	"github.com/Jusonex/RELang/pkg/parser"
)

type RELangGenerator struct {
	*parser.BaseRELangListener

	Emitter      *CppCodeEmitter
	ContextStack *GeneratorContextStack
	State        struct {
		FunctionName  string
		MemoryAddress string
		VariableName  string
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
	s.Emitter.EmitClassDeclarationStart(ctx.Name().GetText())
	s.ContextStack.Push(CONTEXT_CLASS_DECL)
}

func (s *RELangGenerator) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.Emitter.EmitClassDeclarationEnd()
	s.ContextStack.Pop(CONTEXT_CLASS_DECL)
}

func (s *RELangGenerator) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	fmt.Println("declare function: ", ctx.Name())
}

func (s *RELangGenerator) EnterFunctionReturnType(ctx *parser.FunctionReturnTypeContext) {

}

func (s *RELangGenerator) EnterFunctionParamList(ctx *parser.FunctionParamListContext) {

}

func (s *RELangGenerator) EnterFunctionParameter(ctx *parser.FunctionParameterContext) {

}

func (s *RELangGenerator) EnterVariableDeclaration(ctx *parser.VariableDeclarationContext) {

}

func (s *RELangGenerator) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	if s.ContextStack.Contains(CONTEXT_CLASS_DECL) {
		s.Emitter.EmitClassVariableDeclaration(CppVariable{Name: ctx.Name().GetText(), Type: s.State.VariableName}, s.State.MemoryAddress)
	} else {
		s.Emitter.EmitVariableDeclaration(CppVariable{Name: ctx.Name().GetText(), Type: s.State.VariableName}, s.State.MemoryAddress)
	}
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
