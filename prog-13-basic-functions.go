package main

import "fmt"

func average(arr []float64) float64 {
	total := 0.0
	for _, value := range arr {
		total += value
	}
	return total / float64(len(arr))
}

func main() {
	random := []float64{32.23, 23.32, 64.46, 46.64}
	fmt.Println(average(random))
}
