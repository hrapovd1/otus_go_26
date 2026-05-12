package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	s1 := string("Hello, OTUS!")
	fmt.Println(reverse.String(s1))
}
