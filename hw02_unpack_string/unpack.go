package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

type pointer struct {
	currChar    rune
	currIsDigit bool

	nextChar    rune
	nextIsDigit bool

	prevChar    rune
	prevIsDigit bool
}

func pointerInit(i int, str []rune) (p pointer) {
	p = pointer{}
	if i == 0 {
		p.prevChar = '&'
		p.prevIsDigit = false
		p.nextChar = str[i+1]
		p.nextIsDigit = isDigit(str[i+1])
	}
	if i+1 == len(str) {
		p.prevChar = str[i-1]
		p.prevIsDigit = isDigit(str[i-1])
		p.nextChar = '&'
		p.nextIsDigit = false
	}
	if i < len(str)-1 && i > 0 {
		p.prevChar = str[i-1]
		p.prevIsDigit = isDigit(str[i-1])
		p.nextChar = str[i+1]
		p.nextIsDigit = isDigit(str[i+1])
	}

	p.currChar = str[i]
	p.currIsDigit = isDigit(str[i])
	return p
}

func isDigit(r rune) bool {
	_, err := strconv.Atoi(string(r))
	return err == nil
}

func Unpack(str string) (string, error) {
	runes := []rune(str)
	f1 := false
	f2 := false
	var c strings.Builder

	for i := 0; i < len(runes); i++ {
		p := pointerInit(i, runes)

		if runes[i] == '\\' {
			if p.nextIsDigit {
				fmt.Fprintf(&c, "%s", string(p.nextChar))
				i++
				f1 = true
				continue
			}
			if p.nextChar == '\\' {
				fmt.Fprintf(&c, "%s", string(p.nextChar))
				i++
				f2 = true
				continue
			}
		}

		if f1 && p.currIsDigit {
			g, _ := strconv.Atoi(string(runes[i]))
			fmt.Fprintf(&c, "%s", strings.Repeat(string(p.prevChar), g-1))
			f1 = false
			continue
		}

		if f2 && p.currIsDigit {
			g, _ := strconv.Atoi(string(runes[i]))
			fmt.Fprintf(&c, "%s", strings.Repeat(string(p.prevChar), g-1))
			f2 = false
			continue
		}

		if p.currIsDigit {
			if i == 0 {
				return "", ErrInvalidString
			}
			if p.prevIsDigit {
				return "", ErrInvalidString
			}
		}

		if !p.currIsDigit && p.nextIsDigit {
			d, _ := strconv.Atoi(string(p.nextChar))
			fmt.Fprintf(&c, "%s", strings.Repeat(string(p.currChar), d))
			i++
			continue
		}

		fmt.Fprintf(&c, "%s", string(runes[i]))
	}
	return c.String(), nil
}
