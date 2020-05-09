package main

import (
	"log"
	"io"
	"os"
	"strings"
)

func main(){
	//name:="Javed"

	str:=`
		<!DOCTYPE html>
		<html>
			<title>Test Page</title>
			<body>
				<h1>This is from go...!</h1>
			</body>

		</html>`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating the file",err)

	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))

}