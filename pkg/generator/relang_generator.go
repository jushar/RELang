package generator

import (
	"fmt"

	"github.com/Jusonex/RELang/pkg/parser"
)

type RELangGenerator struct {
	*parser.BaseRELangListener
}

func NewRELangGenerator() *RELangGenerator {
	return &RELangGenerator{}
}

func (s *RELangGenerator) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	fmt.Println("declare class: ", ctx.Name())
}

func (s *RELangGenerator) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	fmt.Println("declare function: ", ctx.Name())
}
