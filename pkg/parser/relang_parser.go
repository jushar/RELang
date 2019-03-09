// Code generated from grammar\RELang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // RELang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 146,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 7, 2, 38, 10, 2, 12, 2, 14, 2, 41, 11, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 48, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 54, 10, 4, 3, 5,
	3, 5, 6, 5, 58, 10, 5, 13, 5, 14, 5, 59, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 72, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 82, 10, 9, 3, 10, 3, 10, 5, 10, 86, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 95, 10, 12, 12,
	12, 14, 12, 98, 11, 12, 5, 12, 100, 10, 12, 3, 13, 3, 13, 3, 13, 5, 13,
	105, 10, 13, 3, 14, 3, 14, 3, 14, 5, 14, 110, 10, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 7, 15, 118, 10, 15, 12, 15, 14, 15, 121, 11, 15,
	5, 15, 123, 10, 15, 3, 15, 3, 15, 7, 15, 127, 10, 15, 12, 15, 14, 15, 130,
	11, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 5, 16, 137, 10, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 2, 2, 19, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 3, 3, 2, 4, 14, 2, 151,
	2, 39, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 53, 3, 2, 2, 2, 8, 55, 3, 2, 2,
	2, 10, 61, 3, 2, 2, 2, 12, 63, 3, 2, 2, 2, 14, 71, 3, 2, 2, 2, 16, 73,
	3, 2, 2, 2, 18, 85, 3, 2, 2, 2, 20, 87, 3, 2, 2, 2, 22, 99, 3, 2, 2, 2,
	24, 104, 3, 2, 2, 2, 26, 106, 3, 2, 2, 2, 28, 111, 3, 2, 2, 2, 30, 136,
	3, 2, 2, 2, 32, 140, 3, 2, 2, 2, 34, 143, 3, 2, 2, 2, 36, 38, 5, 4, 3,
	2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40,
	3, 2, 2, 2, 40, 3, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 48, 5, 28, 15, 2,
	43, 48, 5, 16, 9, 2, 44, 48, 5, 26, 14, 2, 45, 48, 5, 32, 17, 2, 46, 48,
	5, 34, 18, 2, 47, 42, 3, 2, 2, 2, 47, 43, 3, 2, 2, 2, 47, 44, 3, 2, 2,
	2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50,
	7, 34, 2, 2, 50, 5, 3, 2, 2, 2, 51, 54, 5, 10, 6, 2, 52, 54, 5, 8, 5, 2,
	53, 51, 3, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 7, 3, 2, 2, 2, 55, 57, 7, 31,
	2, 2, 56, 58, 7, 3, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57,
	3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 9, 3, 2, 2, 2, 61, 62, 9, 2, 2, 2,
	62, 11, 3, 2, 2, 2, 63, 64, 7, 15, 2, 2, 64, 65, 7, 32, 2, 2, 65, 13, 3,
	2, 2, 2, 66, 72, 3, 2, 2, 2, 67, 72, 7, 16, 2, 2, 68, 72, 7, 17, 2, 2,
	69, 72, 7, 18, 2, 2, 70, 72, 7, 19, 2, 2, 71, 66, 3, 2, 2, 2, 71, 67, 3,
	2, 2, 2, 71, 68, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72,
	15, 3, 2, 2, 2, 73, 74, 5, 24, 13, 2, 74, 75, 5, 18, 10, 2, 75, 76, 5,
	14, 8, 2, 76, 77, 7, 31, 2, 2, 77, 78, 7, 20, 2, 2, 78, 79, 5, 22, 12,
	2, 79, 81, 7, 21, 2, 2, 80, 82, 5, 12, 7, 2, 81, 80, 3, 2, 2, 2, 81, 82,
	3, 2, 2, 2, 82, 17, 3, 2, 2, 2, 83, 86, 5, 6, 4, 2, 84, 86, 7, 22, 2, 2,
	85, 83, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2, 86, 19, 3, 2, 2, 2, 87, 88, 5,
	6, 4, 2, 88, 89, 7, 31, 2, 2, 89, 21, 3, 2, 2, 2, 90, 100, 3, 2, 2, 2,
	91, 96, 5, 20, 11, 2, 92, 93, 7, 23, 2, 2, 93, 95, 5, 20, 11, 2, 94, 92,
	3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2,
	97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 90, 3, 2, 2, 2, 99, 91, 3,
	2, 2, 2, 100, 23, 3, 2, 2, 2, 101, 105, 3, 2, 2, 2, 102, 105, 7, 24, 2,
	2, 103, 105, 7, 25, 2, 2, 104, 101, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104,
	103, 3, 2, 2, 2, 105, 25, 3, 2, 2, 2, 106, 107, 5, 6, 4, 2, 107, 109, 7,
	31, 2, 2, 108, 110, 5, 12, 7, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3, 2,
	2, 2, 110, 27, 3, 2, 2, 2, 111, 112, 7, 26, 2, 2, 112, 122, 7, 31, 2, 2,
	113, 114, 7, 27, 2, 2, 114, 119, 7, 31, 2, 2, 115, 116, 7, 23, 2, 2, 116,
	118, 7, 31, 2, 2, 117, 115, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117,
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 123, 3, 2, 2, 2, 121, 119, 3, 2,
	2, 2, 122, 113, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2,
	124, 128, 7, 28, 2, 2, 125, 127, 5, 30, 16, 2, 126, 125, 3, 2, 2, 2, 127,
	130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 131,
	3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 29, 2, 2, 132, 29, 3, 2,
	2, 2, 133, 137, 5, 16, 9, 2, 134, 137, 5, 26, 14, 2, 135, 137, 5, 34, 18,
	2, 136, 133, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 135, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 139, 7, 34, 2, 2, 139, 31, 3, 2, 2, 2, 140, 141,
	7, 30, 2, 2, 141, 142, 7, 33, 2, 2, 142, 33, 3, 2, 2, 2, 143, 144, 7, 38,
	2, 2, 144, 35, 3, 2, 2, 2, 17, 39, 47, 53, 59, 71, 81, 85, 96, 99, 104,
	109, 119, 122, 128, 136,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'*'", "'bool'", "'int8'", "'int16'", "'int32'", "'int64'", "'uint8'",
	"'uint16'", "'uint32'", "'uint64'", "'float32'", "'float64'", "'@'", "'__cdecl'",
	"'__stdcall'", "'__thiscall'", "'__fastcall'", "'('", "')'", "'void'",
	"','", "'virtual'", "'static'", "'class'", "':'", "'{'", "'}'", "'import'",
	"", "", "", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "Name", "HexInteger", "NormalString",
	"Separator", "Whitespace", "LineComment", "BlockComment", "RawBlock",
}

var ruleNames = []string{
	"init", "expression", "variableType", "pointer", "primitiveType", "memoryAddress",
	"callingConvention", "functionDeclaration", "functionReturnType", "functionParameter",
	"functionParamList", "functionModifier", "variableDeclaration", "classDeclaration",
	"classExpression", "importExpression", "rawExpression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RELangParser struct {
	*antlr.BaseParser
}

func NewRELangParser(input antlr.TokenStream) *RELangParser {
	this := new(RELangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "RELang.g4"

	return this
}

// RELangParser tokens.
const (
	RELangParserEOF          = antlr.TokenEOF
	RELangParserT__0         = 1
	RELangParserT__1         = 2
	RELangParserT__2         = 3
	RELangParserT__3         = 4
	RELangParserT__4         = 5
	RELangParserT__5         = 6
	RELangParserT__6         = 7
	RELangParserT__7         = 8
	RELangParserT__8         = 9
	RELangParserT__9         = 10
	RELangParserT__10        = 11
	RELangParserT__11        = 12
	RELangParserT__12        = 13
	RELangParserT__13        = 14
	RELangParserT__14        = 15
	RELangParserT__15        = 16
	RELangParserT__16        = 17
	RELangParserT__17        = 18
	RELangParserT__18        = 19
	RELangParserT__19        = 20
	RELangParserT__20        = 21
	RELangParserT__21        = 22
	RELangParserT__22        = 23
	RELangParserT__23        = 24
	RELangParserT__24        = 25
	RELangParserT__25        = 26
	RELangParserT__26        = 27
	RELangParserT__27        = 28
	RELangParserName         = 29
	RELangParserHexInteger   = 30
	RELangParserNormalString = 31
	RELangParserSeparator    = 32
	RELangParserWhitespace   = 33
	RELangParserLineComment  = 34
	RELangParserBlockComment = 35
	RELangParserRawBlock     = 36
)

// RELangParser rules.
const (
	RELangParserRULE_init                = 0
	RELangParserRULE_expression          = 1
	RELangParserRULE_variableType        = 2
	RELangParserRULE_pointer             = 3
	RELangParserRULE_primitiveType       = 4
	RELangParserRULE_memoryAddress       = 5
	RELangParserRULE_callingConvention   = 6
	RELangParserRULE_functionDeclaration = 7
	RELangParserRULE_functionReturnType  = 8
	RELangParserRULE_functionParameter   = 9
	RELangParserRULE_functionParamList   = 10
	RELangParserRULE_functionModifier    = 11
	RELangParserRULE_variableDeclaration = 12
	RELangParserRULE_classDeclaration    = 13
	RELangParserRULE_classExpression     = 14
	RELangParserRULE_importExpression    = 15
	RELangParserRULE_rawExpression       = 16
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_init
	return p
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InitContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *RELangParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RELangParserRULE_init)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RELangParserT__1)|(1<<RELangParserT__2)|(1<<RELangParserT__3)|(1<<RELangParserT__4)|(1<<RELangParserT__5)|(1<<RELangParserT__6)|(1<<RELangParserT__7)|(1<<RELangParserT__8)|(1<<RELangParserT__9)|(1<<RELangParserT__10)|(1<<RELangParserT__11)|(1<<RELangParserT__19)|(1<<RELangParserT__21)|(1<<RELangParserT__22)|(1<<RELangParserT__23)|(1<<RELangParserT__27)|(1<<RELangParserName))) != 0) || _la == RELangParserRawBlock {
		{
			p.SetState(34)
			p.Expression()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Separator() antlr.TerminalNode {
	return s.GetToken(RELangParserSeparator, 0)
}

func (s *ExpressionContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *ExpressionContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *ExpressionContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ExpressionContext) ImportExpression() IImportExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportExpressionContext)
}

func (s *ExpressionContext) RawExpression() IRawExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRawExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRawExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RELangParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RELangParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(40)
			p.ClassDeclaration()
		}

	case 2:
		{
			p.SetState(41)
			p.FunctionDeclaration()
		}

	case 3:
		{
			p.SetState(42)
			p.VariableDeclaration()
		}

	case 4:
		{
			p.SetState(43)
			p.ImportExpression()
		}

	case 5:
		{
			p.SetState(44)
			p.RawExpression()
		}

	}
	{
		p.SetState(47)
		p.Match(RELangParserSeparator)
	}

	return localctx
}

// IVariableTypeContext is an interface to support dynamic dispatch.
type IVariableTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableTypeContext differentiates from other interfaces.
	IsVariableTypeContext()
}

type VariableTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableTypeContext() *VariableTypeContext {
	var p = new(VariableTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_variableType
	return p
}

func (*VariableTypeContext) IsVariableTypeContext() {}

func NewVariableTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableTypeContext {
	var p = new(VariableTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_variableType

	return p
}

func (s *VariableTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *VariableTypeContext) Pointer() IPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *VariableTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterVariableType(s)
	}
}

func (s *VariableTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitVariableType(s)
	}
}

func (p *RELangParser) VariableType() (localctx IVariableTypeContext) {
	localctx = NewVariableTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RELangParserRULE_variableType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RELangParserT__1, RELangParserT__2, RELangParserT__3, RELangParserT__4, RELangParserT__5, RELangParserT__6, RELangParserT__7, RELangParserT__8, RELangParserT__9, RELangParserT__10, RELangParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.PrimitiveType()
		}

	case RELangParserName:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Pointer()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPointerContext is an interface to support dynamic dispatch.
type IPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointerContext differentiates from other interfaces.
	IsPointerContext()
}

type PointerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerContext() *PointerContext {
	var p = new(PointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_pointer
	return p
}

func (*PointerContext) IsPointerContext() {}

func NewPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerContext {
	var p = new(PointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_pointer

	return p
}

func (s *PointerContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerContext) Name() antlr.TerminalNode {
	return s.GetToken(RELangParserName, 0)
}

func (s *PointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterPointer(s)
	}
}

func (s *PointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitPointer(s)
	}
}

func (p *RELangParser) Pointer() (localctx IPointerContext) {
	localctx = NewPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RELangParserRULE_pointer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(RELangParserName)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RELangParserT__0 {
		{
			p.SetState(54)
			p.Match(RELangParserT__0)
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_primitiveType
	return p
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (p *RELangParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RELangParserRULE_primitiveType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RELangParserT__1)|(1<<RELangParserT__2)|(1<<RELangParserT__3)|(1<<RELangParserT__4)|(1<<RELangParserT__5)|(1<<RELangParserT__6)|(1<<RELangParserT__7)|(1<<RELangParserT__8)|(1<<RELangParserT__9)|(1<<RELangParserT__10)|(1<<RELangParserT__11))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMemoryAddressContext is an interface to support dynamic dispatch.
type IMemoryAddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemoryAddressContext differentiates from other interfaces.
	IsMemoryAddressContext()
}

type MemoryAddressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemoryAddressContext() *MemoryAddressContext {
	var p = new(MemoryAddressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_memoryAddress
	return p
}

func (*MemoryAddressContext) IsMemoryAddressContext() {}

func NewMemoryAddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemoryAddressContext {
	var p = new(MemoryAddressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_memoryAddress

	return p
}

func (s *MemoryAddressContext) GetParser() antlr.Parser { return s.parser }

func (s *MemoryAddressContext) HexInteger() antlr.TerminalNode {
	return s.GetToken(RELangParserHexInteger, 0)
}

func (s *MemoryAddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemoryAddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemoryAddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterMemoryAddress(s)
	}
}

func (s *MemoryAddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitMemoryAddress(s)
	}
}

func (p *RELangParser) MemoryAddress() (localctx IMemoryAddressContext) {
	localctx = NewMemoryAddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RELangParserRULE_memoryAddress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(RELangParserT__12)
	}
	{
		p.SetState(62)
		p.Match(RELangParserHexInteger)
	}

	return localctx
}

// ICallingConventionContext is an interface to support dynamic dispatch.
type ICallingConventionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallingConventionContext differentiates from other interfaces.
	IsCallingConventionContext()
}

type CallingConventionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallingConventionContext() *CallingConventionContext {
	var p = new(CallingConventionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_callingConvention
	return p
}

func (*CallingConventionContext) IsCallingConventionContext() {}

func NewCallingConventionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallingConventionContext {
	var p = new(CallingConventionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_callingConvention

	return p
}

func (s *CallingConventionContext) GetParser() antlr.Parser { return s.parser }
func (s *CallingConventionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallingConventionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallingConventionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterCallingConvention(s)
	}
}

func (s *CallingConventionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitCallingConvention(s)
	}
}

func (p *RELangParser) CallingConvention() (localctx ICallingConventionContext) {
	localctx = NewCallingConventionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RELangParserRULE_callingConvention)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RELangParserName:
		p.EnterOuterAlt(localctx, 1)

	case RELangParserT__13:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Match(RELangParserT__13)
		}

	case RELangParserT__14:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Match(RELangParserT__14)
		}

	case RELangParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.Match(RELangParserT__15)
		}

	case RELangParserT__16:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.Match(RELangParserT__16)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) FunctionModifier() IFunctionModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionModifierContext)
}

func (s *FunctionDeclarationContext) FunctionReturnType() IFunctionReturnTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionReturnTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionReturnTypeContext)
}

func (s *FunctionDeclarationContext) CallingConvention() ICallingConventionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallingConventionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallingConventionContext)
}

func (s *FunctionDeclarationContext) Name() antlr.TerminalNode {
	return s.GetToken(RELangParserName, 0)
}

func (s *FunctionDeclarationContext) FunctionParamList() IFunctionParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParamListContext)
}

func (s *FunctionDeclarationContext) MemoryAddress() IMemoryAddressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryAddressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryAddressContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *RELangParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RELangParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.FunctionModifier()
	}
	{
		p.SetState(72)
		p.FunctionReturnType()
	}
	{
		p.SetState(73)
		p.CallingConvention()
	}
	{
		p.SetState(74)
		p.Match(RELangParserName)
	}
	{
		p.SetState(75)
		p.Match(RELangParserT__17)
	}
	{
		p.SetState(76)
		p.FunctionParamList()
	}
	{
		p.SetState(77)
		p.Match(RELangParserT__18)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RELangParserT__12 {
		{
			p.SetState(78)
			p.MemoryAddress()
		}

	}

	return localctx
}

// IFunctionReturnTypeContext is an interface to support dynamic dispatch.
type IFunctionReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionReturnTypeContext differentiates from other interfaces.
	IsFunctionReturnTypeContext()
}

type FunctionReturnTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionReturnTypeContext() *FunctionReturnTypeContext {
	var p = new(FunctionReturnTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_functionReturnType
	return p
}

func (*FunctionReturnTypeContext) IsFunctionReturnTypeContext() {}

func NewFunctionReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionReturnTypeContext {
	var p = new(FunctionReturnTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_functionReturnType

	return p
}

func (s *FunctionReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionReturnTypeContext) VariableType() IVariableTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *FunctionReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterFunctionReturnType(s)
	}
}

func (s *FunctionReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitFunctionReturnType(s)
	}
}

func (p *RELangParser) FunctionReturnType() (localctx IFunctionReturnTypeContext) {
	localctx = NewFunctionReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RELangParserRULE_functionReturnType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RELangParserT__1, RELangParserT__2, RELangParserT__3, RELangParserT__4, RELangParserT__5, RELangParserT__6, RELangParserT__7, RELangParserT__8, RELangParserT__9, RELangParserT__10, RELangParserT__11, RELangParserName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.VariableType()
		}

	case RELangParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(RELangParserT__19)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionParameterContext is an interface to support dynamic dispatch.
type IFunctionParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParameterContext differentiates from other interfaces.
	IsFunctionParameterContext()
}

type FunctionParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterContext() *FunctionParameterContext {
	var p = new(FunctionParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_functionParameter
	return p
}

func (*FunctionParameterContext) IsFunctionParameterContext() {}

func NewFunctionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterContext {
	var p = new(FunctionParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_functionParameter

	return p
}

func (s *FunctionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterContext) VariableType() IVariableTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *FunctionParameterContext) Name() antlr.TerminalNode {
	return s.GetToken(RELangParserName, 0)
}

func (s *FunctionParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterFunctionParameter(s)
	}
}

func (s *FunctionParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitFunctionParameter(s)
	}
}

func (p *RELangParser) FunctionParameter() (localctx IFunctionParameterContext) {
	localctx = NewFunctionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RELangParserRULE_functionParameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.VariableType()
	}
	{
		p.SetState(86)
		p.Match(RELangParserName)
	}

	return localctx
}

// IFunctionParamListContext is an interface to support dynamic dispatch.
type IFunctionParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParamListContext differentiates from other interfaces.
	IsFunctionParamListContext()
}

type FunctionParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParamListContext() *FunctionParamListContext {
	var p = new(FunctionParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_functionParamList
	return p
}

func (*FunctionParamListContext) IsFunctionParamListContext() {}

func NewFunctionParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParamListContext {
	var p = new(FunctionParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_functionParamList

	return p
}

func (s *FunctionParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParamListContext) AllFunctionParameter() []IFunctionParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem())
	var tst = make([]IFunctionParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionParameterContext)
		}
	}

	return tst
}

func (s *FunctionParamListContext) FunctionParameter(i int) IFunctionParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterContext)
}

func (s *FunctionParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterFunctionParamList(s)
	}
}

func (s *FunctionParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitFunctionParamList(s)
	}
}

func (p *RELangParser) FunctionParamList() (localctx IFunctionParamListContext) {
	localctx = NewFunctionParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RELangParserRULE_functionParamList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RELangParserT__18:
		p.EnterOuterAlt(localctx, 1)

	case RELangParserT__1, RELangParserT__2, RELangParserT__3, RELangParserT__4, RELangParserT__5, RELangParserT__6, RELangParserT__7, RELangParserT__8, RELangParserT__9, RELangParserT__10, RELangParserT__11, RELangParserName:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.FunctionParameter()
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RELangParserT__20 {
			{
				p.SetState(90)
				p.Match(RELangParserT__20)
			}
			{
				p.SetState(91)
				p.FunctionParameter()
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionModifierContext is an interface to support dynamic dispatch.
type IFunctionModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionModifierContext differentiates from other interfaces.
	IsFunctionModifierContext()
}

type FunctionModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionModifierContext() *FunctionModifierContext {
	var p = new(FunctionModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_functionModifier
	return p
}

func (*FunctionModifierContext) IsFunctionModifierContext() {}

func NewFunctionModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionModifierContext {
	var p = new(FunctionModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_functionModifier

	return p
}

func (s *FunctionModifierContext) GetParser() antlr.Parser { return s.parser }
func (s *FunctionModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterFunctionModifier(s)
	}
}

func (s *FunctionModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitFunctionModifier(s)
	}
}

func (p *RELangParser) FunctionModifier() (localctx IFunctionModifierContext) {
	localctx = NewFunctionModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RELangParserRULE_functionModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RELangParserT__1, RELangParserT__2, RELangParserT__3, RELangParserT__4, RELangParserT__5, RELangParserT__6, RELangParserT__7, RELangParserT__8, RELangParserT__9, RELangParserT__10, RELangParserT__11, RELangParserT__19, RELangParserName:
		p.EnterOuterAlt(localctx, 1)

	case RELangParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Match(RELangParserT__21)
		}

	case RELangParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.Match(RELangParserT__22)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VariableType() IVariableTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *VariableDeclarationContext) Name() antlr.TerminalNode {
	return s.GetToken(RELangParserName, 0)
}

func (s *VariableDeclarationContext) MemoryAddress() IMemoryAddressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryAddressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryAddressContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *RELangParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RELangParserRULE_variableDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.VariableType()
	}
	{
		p.SetState(105)
		p.Match(RELangParserName)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RELangParserT__12 {
		{
			p.SetState(106)
			p.MemoryAddress()
		}

	}

	return localctx
}

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_classDeclaration
	return p
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) AllName() []antlr.TerminalNode {
	return s.GetTokens(RELangParserName)
}

func (s *ClassDeclarationContext) Name(i int) antlr.TerminalNode {
	return s.GetToken(RELangParserName, i)
}

func (s *ClassDeclarationContext) AllClassExpression() []IClassExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassExpressionContext)(nil)).Elem())
	var tst = make([]IClassExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassExpressionContext)
		}
	}

	return tst
}

func (s *ClassDeclarationContext) ClassExpression(i int) IClassExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassExpressionContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (p *RELangParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RELangParserRULE_classDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(RELangParserT__23)
	}
	{
		p.SetState(110)
		p.Match(RELangParserName)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RELangParserT__24 {
		{
			p.SetState(111)
			p.Match(RELangParserT__24)
		}
		{
			p.SetState(112)
			p.Match(RELangParserName)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RELangParserT__20 {
			{
				p.SetState(113)
				p.Match(RELangParserT__20)
			}
			{
				p.SetState(114)
				p.Match(RELangParserName)
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(122)
		p.Match(RELangParserT__25)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RELangParserT__1)|(1<<RELangParserT__2)|(1<<RELangParserT__3)|(1<<RELangParserT__4)|(1<<RELangParserT__5)|(1<<RELangParserT__6)|(1<<RELangParserT__7)|(1<<RELangParserT__8)|(1<<RELangParserT__9)|(1<<RELangParserT__10)|(1<<RELangParserT__11)|(1<<RELangParserT__19)|(1<<RELangParserT__21)|(1<<RELangParserT__22)|(1<<RELangParserName))) != 0) || _la == RELangParserRawBlock {
		{
			p.SetState(123)
			p.ClassExpression()
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
		p.Match(RELangParserT__26)
	}

	return localctx
}

// IClassExpressionContext is an interface to support dynamic dispatch.
type IClassExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassExpressionContext differentiates from other interfaces.
	IsClassExpressionContext()
}

type ClassExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassExpressionContext() *ClassExpressionContext {
	var p = new(ClassExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_classExpression
	return p
}

func (*ClassExpressionContext) IsClassExpressionContext() {}

func NewClassExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassExpressionContext {
	var p = new(ClassExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_classExpression

	return p
}

func (s *ClassExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassExpressionContext) Separator() antlr.TerminalNode {
	return s.GetToken(RELangParserSeparator, 0)
}

func (s *ClassExpressionContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *ClassExpressionContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ClassExpressionContext) RawExpression() IRawExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRawExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRawExpressionContext)
}

func (s *ClassExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterClassExpression(s)
	}
}

func (s *ClassExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitClassExpression(s)
	}
}

func (p *RELangParser) ClassExpression() (localctx IClassExpressionContext) {
	localctx = NewClassExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RELangParserRULE_classExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(131)
			p.FunctionDeclaration()
		}

	case 2:
		{
			p.SetState(132)
			p.VariableDeclaration()
		}

	case 3:
		{
			p.SetState(133)
			p.RawExpression()
		}

	}
	{
		p.SetState(136)
		p.Match(RELangParserSeparator)
	}

	return localctx
}

// IImportExpressionContext is an interface to support dynamic dispatch.
type IImportExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportExpressionContext differentiates from other interfaces.
	IsImportExpressionContext()
}

type ImportExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportExpressionContext() *ImportExpressionContext {
	var p = new(ImportExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_importExpression
	return p
}

func (*ImportExpressionContext) IsImportExpressionContext() {}

func NewImportExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportExpressionContext {
	var p = new(ImportExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_importExpression

	return p
}

func (s *ImportExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportExpressionContext) NormalString() antlr.TerminalNode {
	return s.GetToken(RELangParserNormalString, 0)
}

func (s *ImportExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterImportExpression(s)
	}
}

func (s *ImportExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitImportExpression(s)
	}
}

func (p *RELangParser) ImportExpression() (localctx IImportExpressionContext) {
	localctx = NewImportExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RELangParserRULE_importExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(RELangParserT__27)
	}
	{
		p.SetState(139)
		p.Match(RELangParserNormalString)
	}

	return localctx
}

// IRawExpressionContext is an interface to support dynamic dispatch.
type IRawExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRawExpressionContext differentiates from other interfaces.
	IsRawExpressionContext()
}

type RawExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawExpressionContext() *RawExpressionContext {
	var p = new(RawExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RELangParserRULE_rawExpression
	return p
}

func (*RawExpressionContext) IsRawExpressionContext() {}

func NewRawExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawExpressionContext {
	var p = new(RawExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RELangParserRULE_rawExpression

	return p
}

func (s *RawExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RawExpressionContext) RawBlock() antlr.TerminalNode {
	return s.GetToken(RELangParserRawBlock, 0)
}

func (s *RawExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.EnterRawExpression(s)
	}
}

func (s *RawExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RELangListener); ok {
		listenerT.ExitRawExpression(s)
	}
}

func (p *RELangParser) RawExpression() (localctx IRawExpressionContext) {
	localctx = NewRawExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RELangParserRULE_rawExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(RELangParserRawBlock)
	}

	return localctx
}
