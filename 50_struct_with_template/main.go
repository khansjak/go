package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi:=sage{
		Name:"Mohandas",
		Motto:"Live and let live",
	}


	sages := []sage{buddha, gandhi}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
