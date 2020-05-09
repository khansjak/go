package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	p1 := person{"James", "Bond", 33}
	p2 := person{"Ian", "Flaming", 55}

	people := []person{p1, p2}
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
