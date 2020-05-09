package main

import "golang.org/src/fmt"

func main() {
	var x [5]int 			// Intializing array x with size 5 of integers
	fmt.Println(x) 			// Printing array [0 0 0 0 0]
	x[3]=42					// Assigning x[3] which will be 4rth element in array to 42
	fmt.Println(x)			// Prints [0 0 0 42 0]
	fmt.Println(len(x))		// Prints 5 which is length of the array
	fmt.Printf("%T\n",x) // Prints [5]int which is same as we have created in first step

}
