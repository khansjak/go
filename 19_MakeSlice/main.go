package main

import "golang.org/src/fmt"

func main() {
	x:=make([]int,10,12)	//Slice of int with size 10 and capacity of 12
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x=append(x,43)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x=append(x,44)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x=append(x,45)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))	//capacity incresed with same set declared as cap after capacity exhausted

}

