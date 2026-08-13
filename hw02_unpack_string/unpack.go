package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}
	if unicode.IsDigit([]rune(s)[0]) {
		return s, ErrInvalidString
	}
	out := strings.Builder{}
	slash := false
	dig := false
	prev := ""

	for _, r := range []rune(s) {
		if unicode.IsDigit(r) {
			if slash {
				prev = string(r)
				slash = false
				continue
			}
			if dig {
				return out.String(), ErrInvalidString
			}
			dig = true
			count, err := strconv.Atoi(string(r))
			if err != nil {
				return out.String(), ErrInvalidString
			}
			if _, err := out.WriteString(strings.Repeat(prev, count)); err != nil {
				return out.String(), ErrInvalidString
			}
			prev = ""
			continue
		} else {
			dig = false
		}
		if slash {
			if slash {
				prev = "\\"
			} else {
				prev += string(r)
			}
			slash = false
			continue
		}
		if r == '\\' {
			slash = true
		}
		if _, err := out.WriteString(prev); err != nil {
			return out.String(), ErrInvalidString
		}
		prev = string(r)
	}
	if _, err := out.WriteString(prev); err != nil {
		return out.String(), ErrInvalidString
	}
	return out.String(), nil
}
