package main

import "golang.org/src/fmt"

func main() {
	a:=foo()
	//b:=bar()

	fmt.Println(a)
	fmt.Println(bar()())
}

func foo() int  {
	s:=42
	return s

}

func bar() func() int{
	return func() int {
		return 43
	}
}