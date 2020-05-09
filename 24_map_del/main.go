package main

import "golang.org/src/fmt"

func main() {
	x:= map[string]int{
		"James":33,
		"Penny":27,
	}

	delete(x,"James")
	fmt.Println(x)

	delete(x,"Penny")
	fmt.Println(x)
	delete(x,"foo")
	fmt.Println(x)

}
