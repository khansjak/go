package main

import "golang.org/src/fmt"

func main() {
	n:="Bond"
	switch n {
	case "Moneypenny","Bond","Dodo":
		fmt.Println("Miss Money penny dodo or James Bond")
	case "Q":
		fmt.Println("This is for Q")
	default:
		fmt.Println("When nothing happens , this happens")


	}
}
