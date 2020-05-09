package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl=template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	loves:= map[string]string{

	}
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.html",loves)
	if err !=nil{
		log.Fatalln(err)
	}
}
