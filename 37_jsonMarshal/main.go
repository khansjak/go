package main



type person struct {
	fname string
	lname string
	age int
}

type isStudent struct {
	person
	goesSchool bool
}

func main()  {
	 p:=isStudent{
	 	person{
	 		"fname1",
	 		"lname1",
	 		22,
		},
		true,
	 }

	 p1:=isStudent{

	 }
}