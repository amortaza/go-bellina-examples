package main

import (
	"fmt"
)

func foo(s string) func() string {
	return func() string {
		return s
	}
}
func main() {
	cb1 := foo("one")
	cb2 := foo("two")

	fmt.Println(cb1())
	fmt.Println(cb2())
}
