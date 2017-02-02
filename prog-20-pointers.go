package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 7
	y := new(int)
	zero(&x)
	zero(y)
	fmt.Println(x)
	fmt.Println(*y)
}
