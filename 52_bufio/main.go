package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s:=`Everything is fucked !`
	scanner :=bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)
	}
}