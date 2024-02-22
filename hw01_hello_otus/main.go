package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

var s = "Hello, OTUS!"

func main() {
	s = reverse.String(s)
	fmt.Printf("%s\n", s)
}
