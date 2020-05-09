package main

import (
	"golang.org/src/fmt"
	"time"
	"runtime"
)

func main()  {
	fmt.Println("Number of Goroutine running:",runtime.NumGoroutine())
	go f()

	fmt.Println("Number of Goroutine running:",runtime.NumGoroutine())
	bar()
	fmt.Println("Number of Goroutine running:",runtime.NumGoroutine())

}

func f(){
	for i := 0; i < 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("n:",i)
		fmt.Println("Number of Goroutine running:",runtime.NumGoroutine())
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("m:",i)
		fmt.Println("Number of Goroutine running:",runtime.NumGoroutine())
	}
}