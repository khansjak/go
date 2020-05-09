package main

import "golang.org/src/fmt"

func main() {
	s:=sum(1,2,3,4)
	fmt.Println(s)
}

func sum(x ...int) int {
	sum:=1
	for _ , v := range x {
		sum *= v

	}
	return sum
}
