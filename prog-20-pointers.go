package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 7
	zero(&x)
	fmt.Println(x)
}
