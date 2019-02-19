grammar RELang;

// Parser rules
init : expression*;

expression : (classDeclaration | functionDeclaration | variableDeclaration | importExpression) Separator;

variableType : primitiveType | pointer;
pointer : Name '*'+;
primitiveType : 'bool'|'int8'|'int16'|'int32'|'int64'|'uint8'|'uint16'|'uint32'|'uint64'|'float32'|'float64';
memoryAddress : '@' HexInteger;

functionDeclaration : functionReturnType Name '(' functionParamList ')' memoryAddress;
functionReturnType : variableType | 'void';
functionParameter : variableType Name;
functionParamList : | functionParameter (',' functionParameter)*;

variableDeclaration : variableType Name memoryAddress?;

classDeclaration : 'class' Name (':' Name  (',' Name)*)? '{' classExpression* '}';
classExpression : (functionDeclaration | variableDeclaration) Separator;

importExpression : 'import' NormalString;

// Lexer rules
Name : [a-zA-Z_][a-zA-Z_0-9]*;
HexInteger : '0x'(('0'..'9')|('A'..'F'))+;
NormalString : '"' (~('\\'|'"'))* '"';
Separator: ';';

Whitespace : [ \t\u000C\r\n]+ -> skip;
LineComment : '//' ~[\r\n]* -> skip;
BlockComment : '/*' .*? '*/' -> skip;
