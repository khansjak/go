package main

import "golang.org/src/fmt"

func main() {
	x:=map[int]string{
		1:"Javed",
		2:"Parvej",
		3:"Tabrez",
		4:"Shahina",
		5:"Firoz",
	}
	fmt.Println(x,"\n")
	fmt.Println(x[1])
	fmt.Println(x[10])
	v,ok :=x[5]
	fmt.Println(ok)
	fmt.Println(v)

	if v,ok := x[1];ok{
		fmt.Println("The big brother is and will be always ",v)
	}

	y:=map[string]int{
		"Javed":38,
		"Parvej":37,
		"Tabrez":33,
		"Shahina":30,
		"Firoz":28,
	}

	if v , ok := y["Javed"];ok{
		fmt.Println("The Big brother age is ",v)
	}

}
