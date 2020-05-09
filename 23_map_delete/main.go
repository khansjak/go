package main

import "golang.org/src/fmt"

func main() {
	x:= map[string]int{
		"James":33,
		"Penny":27,
	}
	fmt.Println(x)

	x["Todd"]=56

	fmt.Println(x)

	for k , v := range x{
		fmt.Println(k,v)
	}

	if v , ok := x["James"];ok{
		fmt.Println("James age is :",v)
	}

	if v , ok := x["Penny"];ok{
		fmt.Println("Penny age is :",v)
	}
	
}
