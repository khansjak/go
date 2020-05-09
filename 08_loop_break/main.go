package main

import "golang.org/src/fmt"

func main() {
	x:=1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("Done")
}