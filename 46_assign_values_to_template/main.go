package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl=template.Must(template.ParseFiles("tpl.html"))
}

func main(){
	err:=tpl.ExecuteTemplate(os.Stdout,"tpl.html",`Release self focus : embrass other people focus`)
	if err!=nil{
		log.Fatalln("Error",err)
	}
}
