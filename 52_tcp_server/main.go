package main

import (
	"golang.org/src/fmt"
	"net"
	"io"
	"log"
)

func main() {
	li, err := net.Listen("tcp",":8080")
	if err!=nil{
		log.Fatalln(err)
	}
	defer li.Close()

	for{
		conn,err := li.Accept()
		if err!=nil{
			fmt.Println(err)
			continue
		}
	 	io.WriteString(conn,"\n Hello this is a message for you sent over TCP via my newly created server\n")
		fmt.Fprintln(conn,"Also i want you to kill the server once you are done")
		fmt.Fprintln(conn,"Use lsof and kill commands to do same, happy day to you ba bye")
	}
}