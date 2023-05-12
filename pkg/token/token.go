package token

import (
	"fmt"
	"strings"
)

const OPERATORS = "+-*/<>=~!@#%^&|`?"

type TokenType string

// A dollar sign ($) followed by digits is used to represent a positional parameter in the body of a function definition or a prepared statement. In other contexts the dollar sign can be part of an identifier or a dollar-quoted string constant.

// Parentheses (()) have their usual meaning to group expressions and enforce precedence. In some cases parentheses are required as part of the fixed syntax of a particular SQL command.

// Brackets ([]) are used to select the elements of an array. See Section 8.15 for more information on arrays.

// Commas (,) are used in some syntactical constructs to separate the elements of a list.

// The semicolon (;) terminates an SQL command. It cannot appear anywhere within a command, except within a string constant or quoted identifier.

// The colon (:) is used to select “slices” from arrays. (See Section 8.15.) In certain SQL dialects (such as Embedded SQL), the colon is used to prefix variable names.

// The asterisk (*) is used in some contexts to denote all the fields of a table row or composite value. It also has a special meaning when used as the argument of an aggregate function, namely that the aggregate does not require any explicit parameter.

// The period (.) is used in numeric constants, and to separate schema, table, and column names.

const (
	// Special Characters
	DollarSign   = "DollarSign"
	ParenOpen    = "ParenOpen"
	ParenClose   = "ParenClose"
	BracketOpen  = "BracketOpen"
	BracketClose = "BracketClose"
	Comma        = "Comma"
	Semicolon    = "Semicolon"
	Colon        = "Colon"
	Asterisk     = "Asterisk"
	Period       = "Period"

	// All Opertors
	Operator = "Operator"

	// Constants
	StringConstant        = "StringConstant"
	EscapedStringConstant = "EscapedStringConstant"
	UnicodeStringConstant = "UnicodeStringConstant"
	DollarStringConstant  = "DollarStringConstant"
	BitStringConstant     = "BitStringConstant"
	NumericConstant       = "NumericConstant"
	// Other
	FunctionArg = "FunctionArg"
	Comment     = "Comment"

	// // delimeters needing look ahead
	// Minus

	// // seperators
	// Space
	NewLine = "NewLine"

	// Identifiers
	Identifier          = "Identifier"
	DelimitedIdentifier = "DelimitedIdentifier"
)

// delimiter ::= <alpha-literal> | , | ( | ) | < | > | . | : | =
//
//	| * | + | - | / | <> | >= | <=
//
// separator ::= <comment> | <space> | <newline>
// comment   ::= --[<character>]*<newline>
type Token struct {
	ttype  TokenType
	lexeme string
	line   int
	start  int
}

func NewToken(t TokenType, lexeme string, line int, start int) *Token {
	l := Token{ttype: t, lexeme: lexeme, line: line, start: start}
	return &l
}

// for printing tokens
func (t Token) String() string {
	if t.ttype == NewLine {
		return fmt.Sprintf("%v: \\n", t.ttype)
	}
	return fmt.Sprintf("%v: %v", t.ttype, t.lexeme)
}

func CheckKeyword(lexeme string) TokenType {
	lexeme = strings.ToUpper(lexeme)
	if t, ok := KEYWORDS[lexeme]; ok {
		return t
	} else {
		return Identifier
	}
}
