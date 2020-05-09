package main

import "golang.org/src/fmt"

func main() {
	defer foo()
	bar()

}

func foo()  {
	fmt.Println("Hi from foo")
}

func bar()  {
	fmt.Println("Hi from bar")
}