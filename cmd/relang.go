package main

import (
	"os"

	"github.com/Jusonex/RELang/pkg/generator"
	"github.com/Jusonex/RELang/pkg/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	flags "github.com/jessevdk/go-flags"
)

type CliOptions struct {
	Path string `short:"p" long:"path" description:"Path to source file" required:"true"`
}

func main() {
	// Parse arguments
	options := CliOptions{}
	_, err := flags.ParseArgs(&options, os.Args)
	if err != nil {
		panic(err)
	}

	// Open file as stream
	input, err := antlr.NewFileStream(options.Path)
	if err != nil {
		panic(err)
	}

	// Create lexer and token stream
	lexer := parser.NewRELangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser and parse
	p := parser.NewRELangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	// Walk AST
	antlr.ParseTreeWalkerDefault.Walk(generator.NewRELangGenerator(), p.Init())
}
