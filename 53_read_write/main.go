package main

import (
	"fmt"
	"bufio"
	"net"
	"log"
)

func main()  {
	li , err :=net.Listen("tcp",":8080")//Listen when phone rings
	if err!=nil{										// Is there a line intruption ?
		log.Fatalln(err)								// Yeah it is ..
	}
	defer li.Close()									// Hang up the call

	for {
		conn,err := li.Accept()							// Accept the call
		if err!=nil{
			fmt.Println("not connected !",err)		//Connected but noone talks ?
			continue									// Stay on call
		}
		go handle(conn)									// Handle this call anyway
	}

}

func handle(conn net.Conn){
	scanner := bufio.NewScanner(conn)					//Check if there is any noise or sound ?
	for scanner.Scan(){
		ln:=scanner.Text()								//Got one.
		fmt.Println(ln)
	}
	defer conn.Close()
	fmt.Println("Code got here !")
}