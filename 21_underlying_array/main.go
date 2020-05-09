package main

import "golang.org/src/fmt"

func main() {
	x:=[]int{1,2,3,4,5,42}
	fmt.Println(x)

	y:=append(x,44,44,44,44)
	fmt.Println(x)
	fmt.Println(y)

}
