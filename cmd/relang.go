package main

import (
	"os"

	"github.com/Jusonex/RELang/pkg/generator/cpp"
	"github.com/Jusonex/RELang/pkg/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	flags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

type cliOptions struct {
	InputPath  string `short:"p" long:"path" description:"Path to source file" required:"true"`
	OutputPath string `short:"o" long:"out" description:"Path to output file" required:"true"`
	Lint       bool   `short:"l" long:"lint" description:"Lint-only mode"`
}

func main() {
	// Init logger
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
	})

	// Parse arguments
	options := cliOptions{}
	_, err := flags.ParseArgs(&options, os.Args)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("invalid arguments")
	}

	// Open file as stream
	input, err := antlr.NewFileStream(options.InputPath)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("could not open input file")
	}

	// Create lexer and token stream
	lexer := parser.NewRELangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser and parse
	p := parser.NewRELangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	// Walk AST
	chunkReader := parser.NewChunkReader()
	antlr.ParseTreeWalkerDefault.Walk(chunkReader, p.Init())

	// Generate C++ code
	if !options.Lint {
		generator := cpp.NewCodeGenerator(options.OutputPath)
		defer generator.Close()
		generator.Generate(chunkReader.State.Chunk)
	}

	log.Info("compiled successfully")
}
