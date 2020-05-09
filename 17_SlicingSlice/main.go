package main

import "golang.org/src/fmt"

func main() {
	x:=[]int{1,2,3,4,5,6,42}
	// Position of array with values
	// {1   2   3   4   5   6   42} values
	// {0   1   2   3   4   5   6}  Postion
	// [1:2] means start on 1:which is 2 and end and leave on 2 which is 3 * 3 is excluded

	fmt.Println(x[1:2])
	fmt.Println(x[:3]) // prints [1 2 3] which means it counts from pos. 0 for : and ends and excludes on 3rd pos. which is 4

	fmt.Println(x[4:]) //Prints [5 6 42 ] because its started on pos. 4 which is 5 and ends and exlude till last

}
