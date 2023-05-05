package token

import (
	"fmt"
)

// Weekday - Custom type to hold value for weekday ranging from 1-7
type TokenType int

// Declare related constants for each weekday starting with index 1
const (
	// one-char delimeters
	Comma TokenType = iota
	ParenOpen
	ParenClose
	Period
	Colon
	Equals
	Asterisk
	Plus
	ForwardSlash

	// two-char delimeters
	LessThan
	GreaterThan
	NotEquals
	GreaterThanEquals
	LessThanEquals

	// delimeters needing look ahead
	Minus

	// seperators
	Space
	NewLine
	Comment
)

func (tt TokenType) String() string {
	return [...]string{"Comma", "ParenOpen", "ParenClose", "Period", "Colon", "Equals", "Asterisk", "Plus", "ForwardSlash", "LessThan", "GreaterThan", "NotEquals", "GreaterThanEquals", "LessThanEquals", "Minus", "Space", "NewLine", "Comment"}[tt]
}

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
	return fmt.Sprintf("%v", t.ttype)
}
