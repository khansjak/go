package main

import "golang.org/src/fmt"

func main() {
	x:=[]int{1,2,3,4,5,6,42}
	fmt.Println(x)

	x=append(x,43,44,45,46,47,48)
	fmt.Println(x)

	y:=[]int{100,200,300,400,500}
	x=append(x,y...)
	fmt.Println(x)
}
