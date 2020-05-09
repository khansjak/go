package main

import "golang.org/src/fmt"

func main() {
	//x:= 	-->x is
	//[]int --> Slice of integer
	//{1,2,3,4,5,6,7,8,9,0} --> and the values are {1,2,3,4,5,6,7,8,9,0}

	x:=[]int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(x)

	//y:= -->y is
	//[]string --> slice of string
	//{"One","Two","Three","Four","Five"} --> and the vales are {"One","Two","Three","Four","Five"}

	y:=[]string{"One","Two","Three","Four","Five"}
	fmt.Println(y)
}
