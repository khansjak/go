package main

import "golang.org/src/fmt"

func main() {
	switch  {
	case 1==1:
		fmt.Println("1 is equals to 1")
	case 2 > 1:
		fmt.Println("2 is greater than 1")
	case 2 < 1:
		fmt.Println()
	}
}
