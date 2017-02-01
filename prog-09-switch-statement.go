package main

import "fmt"

func main() {
	for i := 1; i < 7; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Some Other Number")
		}
	}
}
