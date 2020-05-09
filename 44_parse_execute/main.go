package main

import (
	"log"
	"golang.org/x/html/template"
	"os"
)

func main() {
	tpl , err := template.ParseFiles("tpl.gohtml")
	if err!= nil {
		log.Fatalln("Error fetching files ",err)
	}

	nf , err := os.Create("inex.html")
	err=tpl.Execute(nf,nil)
	if err!=nil{
		log.Fatalln(err)
	}
}
