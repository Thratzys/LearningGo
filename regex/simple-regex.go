package main

import (
	"regexp"
	"fmt"
)

func main() {

	a := regexp.MustCompile("AB")

	matches := a.FindAllStringIndex("AB C B A", -1)
	fmt.Println(len(matches))
}
