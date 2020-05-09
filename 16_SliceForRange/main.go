package main

import "golang.org/src/fmt"

func main() {
	x:=[]int{1,2,3,4,5,6,7,8,9.0}

	fmt.Println(x[1])
	for _,v :=range x{
		fmt.Print("\t","\t",v)
	}
}
