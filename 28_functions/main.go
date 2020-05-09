package main

import "golang.org/src/fmt"

func main() {
	foo()
	s:=bar(1)
	fmt.Println(s)
	p,q,r:=baz(2,3,4)
	fmt.Println(p,q,r)
	sum(1,2,3,4,5)
}

func foo() {
	fmt.Println("Call from foo")
}

func bar(a int)int  {
	fmt.Println("call form bar , taing argument of ",a)
	return a*100
}

func baz(a int,b int,c int) (int,int,int)  {
	return a*100,b*100,c*100
}

func sum(x ...int) int {
	fmt.Println(x)

	sum:=0

	for i , v := range x {
		sum += v
		fmt.Println("Index:",i,"Adding:",v,"to total:",sum)
	}
	fmt.Println("The total is ",sum)
	return sum

}

