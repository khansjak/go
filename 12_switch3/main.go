package main

import "golang.org/src/fmt"

func main() {
	switch "Bond" {
	case "MoneyPenny":
		fmt.Println("Miss Moneypenny")
	case "Bond":
		fmt.Println("Mr. James Bond")
	case "Dr. No":
		fmt.Println("Dr. No")
	case "Q":
		fmt.Println("Dr. Q")
	}
}
