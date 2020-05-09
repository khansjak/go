package main

import (
	//"fmt"
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl=template.Must(template.ParseFiles("tpl.gohtml"))
}

type sage struct {
	Name string
	Quote string
}

type car struct {
	Manufacturer string
	Year int
	Model string
}

type items struct {
	Wisdom []sage
	Transport []car
}

func main() {
	a := sage{"Xun Yan", "Human Nature and delibrate effort murt unit"}
	b := sage{"Tom Douglus", "Another way i like to barbeque the king salmon is whole stuffed"}
	c := sage{"Dianne Reaves", "I knew what the story behind my dreams was"}

	d := car{"Ford", 2007, "Tracker"}
	e := car{"Mustang", 2019, "Power"}
	f := car{"CAT", 2018, "Mover"}

	sages :=[]sage{a,b,c}
	cars := []car{d,e,f}

	data:=items{sages,
		cars}

	//fmt.Println(data)

	err := tpl.Execute(os.Stdout,data)
	if err!=nil{
		log.Fatalln("Something went wrong",err)
	}


}