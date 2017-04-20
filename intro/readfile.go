package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("arq.txt")
	if err != nil {
		return
	}
	defer file.Close()

	stat, err = file.Stat()
	if err != nil {
		return
	}

	text = make([]byte, stat.Size())

	_, err = file.Read(text)
	if err != nil {
		return
	}

	fmt.Println(string(text));

}
