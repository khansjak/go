package main

import "golang.org/src/fmt"

func main() {
	x:=[]string{"james","Bond","Chocolate","Martini"}
	y:=[]string{"Miss","Moneypenny","Vanilla","Rum"}

	xp:=[][]string{x,y} // have to write two square brackets for multi dimensional array print
	fmt.Println(xp)

}
