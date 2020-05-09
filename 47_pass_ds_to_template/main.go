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

	sages :=[] string{"Gandhi","Buddha","Jesus","Mother Teresa"}

	err:=tpl.ExecuteTemplate(os.Stdout,"tpl.html",sages)
	if err!=nil{
		log.Fatalln("Error",err)
	}
}
