package main

import "golang.org/src/fmt"

type vehicle struct {
	door int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}
type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	v1 := truck{
		vehicle{
			2,
			"White",
		},
		false,
	}

	v2:=sedan{
		vehicle{
			4,
			"Black",
		},
		true,
	}

	fmt.Println(v1.color)
	fmt.Println(v2)


}
