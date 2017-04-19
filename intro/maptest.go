package main

import "fmt"

func main() {
	x := make(map[string]string)
	x["Paulo"] = "Alves"

	fmt.Println(x["Paulo"])
}
