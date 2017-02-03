package main

import "fmt"

type Rectangle struct {
	l, w float64
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}

func main() {
	rect := Rectangle{4, 32.4}
	fmt.Println("Area = ", rect.area())
}
