package main

import "golang.org/src/fmt"

func main() {
	f:= func() {
		fmt.Println("Hello")
	}
	f()

	g := func(x int) {
		fmt.Println("Number entered is :",x)
	}
	g(12)
}
