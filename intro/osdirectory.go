package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	filesinfo, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fileinfo := range filesinfo {
		fmt.Println(fileinfo.Name())
	}
}
