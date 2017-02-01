package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 32
	x[1] = 64
	x[2] = 12
	x[3] = 9
	x[4] = 13

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += float64(x[i])
	}

	fmt.Println("Average = ", total/float64(len(x)))
}
