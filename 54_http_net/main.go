package main

import (
	"golang.org/src/fmt"
	"net/http"
)

type hotdog string

func (m hotdog) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("This is net magic!")

}

func main()  {
	var d hotdog
	http.ListenAndServe(":8080",d)
}