package main

import (
	"golang.org/src/fmt"
	"runtime"
)

func greet(c chan string)  {
	fmt.Println("Hello"+<-c)
	fmt.Println("## of goroutines",runtime.NumGoroutine())
}



func main() {
	c:=make(chan string)
	go greet(c)
	c<-" John"
	go greet(c)
	c<-" Rambo"
	go greet(c)
	//close(c)
	c<-" Ian"
	fmt.Println("# of goroutines",runtime.NumGoroutine())
	close(c)
	fmt.Println("? of goroutines",runtime.NumGoroutine())
}