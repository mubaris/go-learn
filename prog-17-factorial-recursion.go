package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func main() {
	fmt.Println("Factorial of 14 = ", factorial(14))
}
