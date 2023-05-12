package lexer

import (
	"errors"
	"log"
	"sqlriver/pkg/token"
)

func (l *Lexer) consumeOperator() error {
	log.Println("Attempting to consume Operator")
	gobbler := operatorGobbler{l: l}
	if err := l.gobbleChars(gobbler); err != nil {
		return err
	}
	l.addToken(token.Operator)
	return nil
}

func (l *Lexer) consumeStringConstant() error {
	log.Println("Attempting to consume StringConstant")
	gobbler := delimitedGobbler{l: l, d: '\''}
	if err := l.gobbleChars(gobbler); err != nil {
		return err
	}
	l.addToken(token.StringConstant)
	return nil
}

func (l *Lexer) consumeEscapedStringConstant() error {
	log.Println("Attempting to consume EscapedStringConstant")
	gobbler := delimitedGobbler{l: l, d: '\''}
	if err := l.gobbleChars(gobbler); err != nil {
		return err
	}
	l.addToken(token.EscapedStringConstant)
	return nil
}

func (l *Lexer) consumeUnicodeStringConstant() error {
	log.Println("Attempting to consume EscapedStringConstant")
	gobbler := delimitedGobbler{l: l, d: '\''}
	if err := l.gobbleChars(gobbler); err != nil {
		return err
	}
	l.addToken(token.UnicodeStringConstant)
	return nil
}

func (l *Lexer) consumerDollarStringConstant() error {
	log.Println("Attempting to consume DollarStringConstant")
	gobbler := dollarStringGobbler{l: l}
	if err := l.gobbleChars(gobbler); err != nil {
		return err
	}
	l.addToken(token.DollarStringConstant)

	return nil
}

func (l *Lexer) consumerBitStringConstnat() error {
	log.Println("Attempting to consume BitStringConstant")
	gobbler := bitStringGobbler{l: l}
	if err := l.gobbleChars(gobbler); err != nil {
		return err
	}
	l.addToken(token.BitStringConstant)

	return nil
}

func (l *Lexer) consumeNumericConstant() error {
	log.Println("Attempting to consume NumericConstant")
	gobbler := numericGobbler{l: l}
	if err := l.gobbleChars(gobbler); err != nil {
		return err
	}
	l.addToken(token.NumericConstant)

	return nil
}

func (l *Lexer) consumeArg() error {
	log.Println("Attempting to consume FunctionArg")
	gobbler := argGobbler{l: l}
	l.gobbleChars(gobbler)
	l.addToken(token.FunctionArg)
	if l.start == l.next-1 {
		return errors.New("Not an arg")
	}
	return nil
}

func (l *Lexer) consumeComment() error {
	log.Println("Attempting to consume Comment")
	gobbler := commentGobbler{l: l}
	l.gobbleChars(gobbler)
	l.addToken(token.Comment)
	return nil
}

func (l *Lexer) consumeIdentifier() error {
	log.Println("Attempting to consume Identifier")
	gobbler := identifierGobbler{l: l}
	l.gobbleChars(gobbler)
	l.addToken(token.Identifier)
	return nil
}

func (l *Lexer) consumeDelimitedIdentifier() error {
	log.Println("Attempting to consume DelimitedIdentifier")
	gobbler := delimitedGobbler{l: l, d: '"'}
	l.gobbleChars(gobbler)
	l.addToken(token.DelimitedIdentifier)
	return nil
}
