// Code generated from grammar\RELang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // RELang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRELangListener is a complete listener for a parse tree produced by RELangParser.
type BaseRELangListener struct{}

var _ RELangListener = &BaseRELangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRELangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRELangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRELangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRELangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInit is called when production init is entered.
func (s *BaseRELangListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseRELangListener) ExitInit(ctx *InitContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRELangListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRELangListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVariableType is called when production variableType is entered.
func (s *BaseRELangListener) EnterVariableType(ctx *VariableTypeContext) {}

// ExitVariableType is called when production variableType is exited.
func (s *BaseRELangListener) ExitVariableType(ctx *VariableTypeContext) {}

// EnterPointer is called when production pointer is entered.
func (s *BaseRELangListener) EnterPointer(ctx *PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *BaseRELangListener) ExitPointer(ctx *PointerContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseRELangListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseRELangListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterMemoryAddress is called when production memoryAddress is entered.
func (s *BaseRELangListener) EnterMemoryAddress(ctx *MemoryAddressContext) {}

// ExitMemoryAddress is called when production memoryAddress is exited.
func (s *BaseRELangListener) ExitMemoryAddress(ctx *MemoryAddressContext) {}

// EnterCallingConvention is called when production callingConvention is entered.
func (s *BaseRELangListener) EnterCallingConvention(ctx *CallingConventionContext) {}

// ExitCallingConvention is called when production callingConvention is exited.
func (s *BaseRELangListener) ExitCallingConvention(ctx *CallingConventionContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseRELangListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseRELangListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionReturnType is called when production functionReturnType is entered.
func (s *BaseRELangListener) EnterFunctionReturnType(ctx *FunctionReturnTypeContext) {}

// ExitFunctionReturnType is called when production functionReturnType is exited.
func (s *BaseRELangListener) ExitFunctionReturnType(ctx *FunctionReturnTypeContext) {}

// EnterFunctionParameter is called when production functionParameter is entered.
func (s *BaseRELangListener) EnterFunctionParameter(ctx *FunctionParameterContext) {}

// ExitFunctionParameter is called when production functionParameter is exited.
func (s *BaseRELangListener) ExitFunctionParameter(ctx *FunctionParameterContext) {}

// EnterFunctionParamList is called when production functionParamList is entered.
func (s *BaseRELangListener) EnterFunctionParamList(ctx *FunctionParamListContext) {}

// ExitFunctionParamList is called when production functionParamList is exited.
func (s *BaseRELangListener) ExitFunctionParamList(ctx *FunctionParamListContext) {}

// EnterFunctionModifier is called when production functionModifier is entered.
func (s *BaseRELangListener) EnterFunctionModifier(ctx *FunctionModifierContext) {}

// ExitFunctionModifier is called when production functionModifier is exited.
func (s *BaseRELangListener) ExitFunctionModifier(ctx *FunctionModifierContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseRELangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseRELangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseRELangListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseRELangListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassExpression is called when production classExpression is entered.
func (s *BaseRELangListener) EnterClassExpression(ctx *ClassExpressionContext) {}

// ExitClassExpression is called when production classExpression is exited.
func (s *BaseRELangListener) ExitClassExpression(ctx *ClassExpressionContext) {}

// EnterImportExpression is called when production importExpression is entered.
func (s *BaseRELangListener) EnterImportExpression(ctx *ImportExpressionContext) {}

// ExitImportExpression is called when production importExpression is exited.
func (s *BaseRELangListener) ExitImportExpression(ctx *ImportExpressionContext) {}

// EnterRawExpression is called when production rawExpression is entered.
func (s *BaseRELangListener) EnterRawExpression(ctx *RawExpressionContext) {}

// ExitRawExpression is called when production rawExpression is exited.
func (s *BaseRELangListener) ExitRawExpression(ctx *RawExpressionContext) {}
