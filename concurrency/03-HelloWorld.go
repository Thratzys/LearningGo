package main

import(
	"fmt"
)

func main() {
	c := make(chan string)

	go func (c chan string) {
		fmt.Println(<-c)
	}(c)

	for i := 0;  i <= 10; i++ {
		c <- "Hello World"
	}


	var input string
	fmt.Scanln(&input)
}