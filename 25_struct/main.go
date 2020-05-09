package main

import "golang.org/src/fmt"

type person struct {
	fname string
	lname string
	favorites []string
}
func main()  {
	p1:=person{
		fname:"James",
		lname:"Bond",
		favorites:[]string{
			"Cat",
			"Dogs",
			"Parrots",
		},
	}

	p2:=person{
		"Miss",
		"Penny",
		[]string{
			"Dragons",
			"Unicorns",
			"Dwrfs",
			"Sabertooths",
		},
	}



	for _ , v := range p1.favorites{
		fmt.Println(p1.fname,p1.lname,"Likes\n","\t\t",v)
	}

	for _ , v := range p2.favorites{
		fmt.Println(p2.fname,p2.lname,"Likes\n","\t\t",v)

	}
	fmt.Println(p2.fname,p2.lname,"Likes",p1.fname,p1.lname,"who now likes",p1.favorites,"but she like ",p2.favorites)
}