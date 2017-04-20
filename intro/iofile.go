package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bstring, err := ioutil.ReadFile("arq.txt")
	if err != nil {
		fmt.Println("Error")
		return
	}

	str := string(bstring)
	fmt.Println(str)
}	
