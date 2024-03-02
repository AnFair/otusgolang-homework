package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	sb := strings.Builder{}
	var char rune = -1 // используем отрицательный инт в качестве "пустышки".

	for _, c := range str {
		if unicode.IsDigit(c) {
			if char != -1 {
				// '0' - руна, литерал которой равен 48, используем ее для приведения рун цифр к интовому значению
				sb.WriteString(strings.Repeat(string(char), int(c-'0')))
				char = -1
			} else {
				return "", ErrInvalidString
			}
		} else {
			if char != -1 {
				sb.WriteRune(char)
			}
			char = c
		}
	}
	if char != -1 {
		sb.WriteRune(char)
	}

	return sb.String(), nil
}
