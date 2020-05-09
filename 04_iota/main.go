package main

import "fmt"

const (
	a=2017 +iota
	b=2017+ iota
	c
	d

)
func main()  {
	fmt.Println(a,b,c,d)

}