package main

import "golang.org/src/fmt"

/*  1. Create a struct person
	2. Create a function with a method speack() for person
	3. Create a struct secretAgent with inline items from person as member of struct array
	4. Create a function with method speak() for secretAgent
	5. Create an Inteface with method speak()
	7. Create a function with interface as reciever 
 */


type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	wearsSuit bool
}

func (p person) speak()  {
	fmt.Println(p.fname,p.lname,"Speaks from person method")
}

func (s secretAgent) speak()  {
	fmt.Println(s.fname,s.lname,"Speaks from secretAgent method")
}

type humna interface {
	speak()
}

func bar(h humna)  {
	fmt.Println(h.(person).fname)
	fmt.Println(h.(secretAgent).fname)
}


func main() {
	s1:=secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	p:=person{
		"Ian",
		"Fleming",
	}

	fmt.Println(s1)
	s1.speak()
	p.speak()


	}
