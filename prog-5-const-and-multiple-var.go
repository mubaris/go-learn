package main

import "fmt"

func main() {
	const x string = "Hello, Go!!"
	var (
		a = 1
		b = 4.5
		c = "Cool"
		d = true || false
	    )
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
