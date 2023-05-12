package lexer

import (
	"errors"
	"log"
	"sqlriver/pkg/token"
	"sqlriver/pkg/utils"

	"golang.org/x/exp/slices"
)

type Gobbler interface {
	isDone() (bool, error)
}

type stringGobbler struct {
	l *Lexer
}

func (s stringGobbler) isDone() (bool, error) {
	c := s.l.getChar(0)
	log.Printf("%c", c)
	if err := utils.ErrIfNewLine(c); err != nil {
		return true, err
	}
	s.l.advance()
	if c == '\'' {
		if n := s.l.getChar(0); n == '\'' {
			s.l.advance()
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

type bitStringGobbler struct {
	l *Lexer
}

func (b bitStringGobbler) isDone() (bool, error) {
	c := b.l.getChar(0)
	if err := utils.ErrIfNewLine(c); err != nil {
		return true, err
	}
	b.l.advance()

	if c == '\'' {
		return true, nil
	} else if c != '0' && c != '1' {
		return true, errors.New("Bit String Constants cannot contain characters besides 0 or 1")
	}
	return false, nil
}

type dollarStringGobbler struct {
	l *Lexer
}

func (d dollarStringGobbler) isDone() (bool, error) {
	// TODO implement tags
	c := d.l.getChar(0)
	if err := utils.ErrIfNewLine(c); err != nil {
		return true, err
	}
	d.l.advance()

	if c == '$' {
		// TODO this attr could be on the gobbler instead of main lexer
		d.l.n_delims++
	}

	if d.l.n_delims == 3 {
		return true, nil
	}
	return false, nil
}

type numericGobbler struct {
	l *Lexer
}

func (n numericGobbler) isDone() (bool, error) {
	c := n.l.getChar(0)
	if utils.IsDigit(c) {
		n.l.advance()
		return false, nil
	}
	if c == '.' {
		if !n.l.has_radix && !n.l.isAtEnd() {
			if c = n.l.getChar(1); utils.IsDigit(c) {
				n.l.advance()
				n.l.has_radix = true
				return false, nil
			} else {
				return true, errors.New("Error parsing numeric literal")
			}
		} else {
			return true, errors.New("Error parsing numeric literal")
		}
	}
	return true, nil
}

type argGobbler struct {
	l *Lexer
}

func (a argGobbler) isDone() (bool, error) {
	c := a.l.getChar(0)
	if c >= '0' && c <= '9' {
		a.l.advance()
		return false, nil
	}

	return true, nil
}

type operatorGobbler struct {
	l *Lexer
}

func (o operatorGobbler) isDone() (bool, error) {
	c := o.l.getChar(0)
	if slices.Contains([]byte(token.OPERATORS), c) {
		o.l.advance()
		return false, nil
	} else {
		return true, nil
	}
}

type commentGobbler struct {
	l *Lexer
}

func (cg commentGobbler) isDone() (bool, error) {

	if c := cg.l.getChar(0); c != '\n' {
		cg.l.advance()
		return false, nil
	}

	return true, nil

}

type identifierGobbler struct {
	l *Lexer
}

func (i identifierGobbler) isDone() (bool, error) {
	if c := i.l.getChar(0); utils.IsIdentifierChar(c) {
		i.l.advance()
		return false, nil
	}

	return true, nil

}
