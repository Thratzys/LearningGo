package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "teste"

	fmt.Println(str) 
	fmt.Println(strings.Replace(str, "es", "ap", 1))
}
