// Code generated from grammar\RELang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // RELang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RELangListener is a complete listener for a parse tree produced by RELangParser.
type RELangListener interface {
	antlr.ParseTreeListener

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVariableType is called when entering the variableType production.
	EnterVariableType(c *VariableTypeContext)

	// EnterPointer is called when entering the pointer production.
	EnterPointer(c *PointerContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterMemoryAddress is called when entering the memoryAddress production.
	EnterMemoryAddress(c *MemoryAddressContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionReturnType is called when entering the functionReturnType production.
	EnterFunctionReturnType(c *FunctionReturnTypeContext)

	// EnterFunctionParameter is called when entering the functionParameter production.
	EnterFunctionParameter(c *FunctionParameterContext)

	// EnterFunctionParamList is called when entering the functionParamList production.
	EnterFunctionParamList(c *FunctionParamListContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassExpression is called when entering the classExpression production.
	EnterClassExpression(c *ClassExpressionContext)

	// EnterImportExpression is called when entering the importExpression production.
	EnterImportExpression(c *ImportExpressionContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVariableType is called when exiting the variableType production.
	ExitVariableType(c *VariableTypeContext)

	// ExitPointer is called when exiting the pointer production.
	ExitPointer(c *PointerContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitMemoryAddress is called when exiting the memoryAddress production.
	ExitMemoryAddress(c *MemoryAddressContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionReturnType is called when exiting the functionReturnType production.
	ExitFunctionReturnType(c *FunctionReturnTypeContext)

	// ExitFunctionParameter is called when exiting the functionParameter production.
	ExitFunctionParameter(c *FunctionParameterContext)

	// ExitFunctionParamList is called when exiting the functionParamList production.
	ExitFunctionParamList(c *FunctionParamListContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassExpression is called when exiting the classExpression production.
	ExitClassExpression(c *ClassExpressionContext)

	// ExitImportExpression is called when exiting the importExpression production.
	ExitImportExpression(c *ImportExpressionContext)
}
