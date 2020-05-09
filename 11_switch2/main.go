package main

import "golang.org/src/fmt"

func main() {
	switch  {
	case false:
		fmt.Println("This shoudnt not print")
	case true:
		fmt.Println("This should print")
	case (3 == 3):
		fmt.Println("This should also print since 3==3")
		fallthrough
	case (7 == 9):
		fmt.Println("Not printing because 7 isnt equal to 9")
		fallthrough
	default:
		fmt.Println("When nothing happens this happens")
	}
}
