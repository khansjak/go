package main

import "golang.org/src/fmt"

func main() {

	c:=make(chan int)
	fmt.Println("%T",c)
	fmt.Println(c)
}