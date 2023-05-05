package lexer

import (
	"log"
	"sqlriver/pkg/token"
)

type Lexer struct {
	source    string
	sourceLen int
	tokens    []token.Token
	start     int
	current   int
	line      int
}

func NewLexer(source string) *Lexer {
	l := Lexer{source: source, sourceLen: len(source), start: 0, current: 0, line: 1}
	log.Println("Initializing lexer")
	return &l
}

func (l *Lexer) isAtEnd() bool {

	log.Println("Checking for remaining input, current position ", l.current, " of ", l.sourceLen)
	return l.current >= len(l.source)
}

func (l *Lexer) advance() {
	l.current++
}

func (l *Lexer) getNextChar() byte {
	defer l.advance()
	return l.source[l.current]
}

func (l *Lexer) matchNextChar(expected byte) bool {
	if l.isAtEnd() {
		return false
	}
	if l.source[l.current] != expected {
		return false
	}

	l.advance()
	return true
}

func (l *Lexer) peekNextChar() byte {
	return l.source[l.current]
}

func (l *Lexer) addToken(t token.TokenType) {
	lexeme := l.source[l.start:l.current]
	token := token.NewToken(t, lexeme, l.line, l.start)
	l.tokens = append(l.tokens, *token)
	log.Println("Added token of type", t)
}

func (l *Lexer) Lex() []token.Token {
	for !l.isAtEnd() {
		log.Println("Getting next token ...")
		l.nextToken()
	}
	return l.tokens
}

func (l *Lexer) nextToken() {

	switch c := l.getNextChar(); c {
	case ',':
		l.addToken(token.Comma)
	case '(':
		l.addToken(token.ParenOpen)
	case ')':
		l.addToken(token.ParenClose)
	case '.':
		l.addToken(token.Period)
	case ':':
		l.addToken(token.Colon)
	case '=':
		l.addToken(token.Equals)
	case '*':
		l.addToken(token.Asterisk)
	case '+':
		l.addToken(token.Plus)

	case '/':
		l.addToken(token.ForwardSlash)
	case '>':
		if l.matchNextChar('=') {
			l.addToken(token.GreaterThanEquals)
		} else {
			l.addToken(token.GreaterThan)
		}
	case '<':
		if l.matchNextChar('>') {
			l.addToken(token.NotEquals)
		} else if l.matchNextChar('=') {
			l.addToken(token.LessThanEquals)
		} else {
			l.addToken(token.LessThan)
		}
	case '-':
		if l.matchNextChar('-') {
			for !l.isAtEnd() && l.peekNextChar() != '\n' {
				l.advance()
			}
			l.addToken(token.Comment)
		} else {
			l.addToken(token.Minus)
		}
	case ' ':
		l.addToken(token.Space)
	case '\n':
		l.addToken(token.NewLine)
	case '\'':
		l.stringLiteral()
	}
}

func (*Lexer) stringLiteral() {
}
