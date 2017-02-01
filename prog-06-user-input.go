package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := ((input + 64) / (input - 64)) * 3
	fmt.Println(output)
}
