package hw02unpackstring

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")
var b strings.Builder
var c strings.Builder

func isDigit(r rune) bool {
	_, err := strconv.Atoi(string(r))
	if err != nil {
		return false
	}
	return true
}

func Unpack(str string) (string, error) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if len(runes) > i {
			if runes[i] == '\\' {
				if isDigit(runes[i+1]) && runes[i+2] == '\\' {
					fmt.Fprintf(&c, "%c", runes[i+1])
					i = i + 2
					continue
				}
				if isDigit(runes[i+1]) && isDigit(runes[i+2]) {
					num, _ := strconv.Atoi(string(runes[i+2]))
					fmt.Fprintf(&c, "%s", strings.Repeat(string(runes[i+1]), num))
					i = i + 2
					continue
				}
				if runes[i+1] == '\\' && isDigit(runes[i+2]) {
					num, _ := strconv.Atoi(string(runes[i+2]))
					fmt.Fprintf(&c, "%s", strings.Repeat(string(runes[i+1]), num))
					i = i + 2
					continue
				}
				if !isDigit(runes[i+1]) && !isDigit(runes[i+2]) {
					os.Exit(1)
				}
			}
			if isDigit(runes[i]) && i == 0 {
				os.Exit(1)
			}
			if isDigit(runes[i]) && isDigit(runes[i+1]) {
				os.Exit(1)
			}

			if !isDigit(runes[i]) && i < len(runes)-1 {
				if isDigit(runes[i+1]) && !isDigit(runes[i+2]) {
					num, _ := strconv.Atoi(string(runes[i+1]))
					fmt.Fprintf(&c, "%s", strings.Repeat(string(runes[i]), num))
					i = i + 1
					continue
				}
			}

			fmt.Fprintf(&c, "%c", runes[i])
		}
	}
	fmt.Printf("%s => %s\n", str, c.String())
	return c.String(), nil
}
