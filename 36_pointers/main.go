package main

import "golang.org/src/fmt"

func main() {
	a:=42
	fmt.Println(a)
	fmt.Println(&a)

	b:=&a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 44

	fmt.Println(a)

}
