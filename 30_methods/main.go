package main

import "golang.org/src/fmt"

type vehicles struct {
	model string
	make int
	color string
	fuelType string
}

type sedan struct {
	vehicles
	colorsAvailabe []string
	luxury bool
}

type truck struct {
	sedan
	heavyDuty bool
}

func (f truck) goestoIndia()  {
	if f.heavyDuty{
		fmt.Println("This vehciel is heavy duty")
	}else {
		fmt.Println("This is a luxury car")
	}
}

func main()  {
	v1:=truck{
		sedan{
			vehicles{
				"GM",
				2017,
				"White",
				"Petrol",
			},
			[]string{"White","Blue"},
			true,
		},
		false,
	}

	fmt.Println(v1)
	if v1.heavyDuty {
		v1.goestoIndia()
	}


	for _ , v := range v1.colorsAvailabe{
		fmt.Println("This one is",v1.model,v1.make,v1.color,"and the colors available ", v)
	}
}