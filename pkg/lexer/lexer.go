package lexer

import (
	"errors"
	"fmt"
	"log"
	"sqlriver/pkg/token"
	"sqlriver/pkg/utils"

	"golang.org/x/exp/slices"
)

type LexerError struct {
	line int
	char int
	err  error
}

func (e *LexerError) Error() string {
	return fmt.Sprintf("Error on line %v at position %v: %v", e.line, e.char, e.err)
}

type Lexer struct {
	source string
	tokens []token.Token
	total  int // total number of chars in source
	start  int // offset of first char of current token being scanned
	next   int // offset of next token being scanned
	line   int // current line #
	//has_error bool
	n_delims  int
	has_radix bool
}

func NewLexer(source string) *Lexer {
	l := Lexer{source: source, total: len(source), start: 0, next: 0, line: 1}
	log.Println("Initializing lexer")
	return &l
}

func (l *Lexer) isAtEnd() bool {
	return l.next >= l.total
}

func (l *Lexer) advance() {
	log.Printf("Advancing")
	l.next++
}

func (l *Lexer) addToken(t token.TokenType) {

	lexeme := l.source[l.start:l.next]
	if t == token.Identifier {
		t = token.CheckKeyword(lexeme)
	}
	token := token.NewToken(t, lexeme, l.line, l.start)
	l.tokens = append(l.tokens, *token)
	log.Println("Added token of type", t)
}

func (l *Lexer) consumeChar() byte {
	defer l.advance()
	return l.getChar(0)
}

func (l *Lexer) matchChar(expected byte) bool {
	if l.isAtEnd() {
		return false
	}
	if l.getChar(0) != expected {
		return false
	}

	l.advance()
	return true
}

func (l *Lexer) getChar(peek int) byte {
	log.Print(l.next)
	return l.source[l.next+peek]
}

func (l *Lexer) gobbleChars(g Gobbler) error {
	// The isDoneGobbling function can have two flavors
	// Both types will return true if there are no more characters to gobble
	// In some cases an error will be returned (e.g. for string constants we must close the delimeters)

	for !l.isAtEnd() {
		if done, err := g.isDone(); err != nil {
			return err
		} else if done == true {
			return nil
		}
	}
	// reset gobble trackers
	l.n_delims = 0
	l.has_radix = false
	return errors.New("Out of input")
}

func (l *Lexer) consumeToken() error {
	switch c := l.consumeChar(); c {
	case '$':
		if err := l.consumeArg(); err == nil {
			break
		} else if err := l.consumerDollarStringConstant(); err != nil {
			return err
		}
	case '(':
		l.addToken(token.ParenOpen)
	case ')':
		l.addToken(token.ParenClose)
	case '[':
		l.addToken(token.BracketOpen)
	case ']':
		l.addToken(token.BracketClose)
	case ',':
		l.addToken(token.Comma)
	case ';':
		l.addToken(token.Semicolon)
	case ':':
		l.addToken(token.Colon)
	case '*':
		l.addToken(token.Asterisk)
	case '.':
		if !l.isAtEnd() {
			if n := l.getChar(0); utils.IsDigit(n) {
				if err := l.consumeNumericConstant(); err != nil {
					return err
				}
				l.has_radix = true
				break
			}
		}
		l.addToken(token.Period)
	case '"':
		if err := l.consumeDelimitedIdentifier(); err != nil {
			return err
		}
	case '\'':
		if err := l.consumeStringConstant(); err != nil {
			return err
		}
	case 'E':
		if l.matchChar('\'') {
			l.consumeEscapedStringConstant()
		}
	case 'U':
		if l.matchChar('\'') {
			l.consumeUnicodeStringConstant()
		}
	case 'B':
		if l.matchChar('\'') {
			if err := l.consumerBitStringConstnat(); err != nil {
				return err
			}
		}
	case '\n':
		//l.addToken(token.NewLine)
		l.line++
	case '-':
		if l.matchChar('-') {
			l.consumeComment()
		}
	default:
		if slices.Contains([]byte(token.OPERATORS), c) {
			if err := l.consumeOperator(); err != nil {
				return err
			}
		} else if utils.IsDigit(c) {
			if err := l.consumeNumericConstant(); err != nil {
				return err
			}
		} else if utils.IsAlpha(c) {
			if err := l.consumeIdentifier(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("Unexpected character: %c", c)
		}
	}
	return nil
}

func (l *Lexer) Lex() ([]token.Token, error) {
	for !l.isAtEnd() {
		l.start = l.next
		log.Println("Getting next token ...")
		err := l.consumeToken()
		if err != nil {
			return l.tokens, &LexerError{line: l.line, char: l.next, err: err}
		}
	}
	return l.tokens, nil
}
