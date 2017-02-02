package main

import "fmt"

func first() {
	fmt.Println("Print from 1st func")
}

func second() {
	fmt.Println("Print from 2nd func")
}

func main() {
	defer second()
	first()
}
