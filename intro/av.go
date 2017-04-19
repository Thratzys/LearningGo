package main

import "fmt"

func main() {
	input := []float64{100, 200, 300, 400, 500, 600, 700, 800, 900}
	fmt.Println(av(input))
}

func av(input []float64) float64 {
	total := 0.0
	for _, value := range input { // foreach
		total += value
	}
	return total/float64(len(input))
}
