package main

import (
	"os"
)

func main() {
	file, err := os.Create("created.txt")

	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("test")
}
