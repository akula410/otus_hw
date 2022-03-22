package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder
	var symbol string
	var repeat int
	var escape bool
	var err error

	for _, char := range str {
		switch {
		case escape:
			if unicode.IsDigit(char) || char == 92 {
				escape = false
			} else {
				return "", ErrInvalidString
			}
		case char == 92:
			builder.WriteString(symbol)
			symbol = ""
			escape = true
			continue
		case unicode.IsDigit(char):
			repeat, err = strconv.Atoi(string(char))
			if err != nil || len(symbol) == 0 {
				return "", ErrInvalidString
			}
			builder.WriteString(strings.Repeat(symbol, repeat))
			symbol = ""
			continue
		default:
			builder.WriteString(symbol)
		}

		symbol = string(char)
	}

	builder.WriteString(symbol)

	return builder.String(), nil
}
