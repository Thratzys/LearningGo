package main

import "fmt"

func main() {
	fmt.Println(sum(100,20,30,40,50))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
