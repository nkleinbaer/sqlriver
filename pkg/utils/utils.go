package utils

import "errors"

func ErrIfNewLine(b byte) error {
	if b == '\n' {
		return errors.New("Unexpected EOL while scanning")
	}
	return nil
}

func IsAlpha(b byte) bool {
	return b == '_' || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func IsDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func IsIdentifierChar(b byte) bool {
	if IsDigit(b) || IsAlpha(b) || b == '$' {
		return true
	}
	return false
}
