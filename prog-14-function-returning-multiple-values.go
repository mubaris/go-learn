package main

import "fmt"

func myFunction(flash int) (string, string) {
	if flash == 2 {
		return "Barry", "Allen"
	} else if flash == 3 {
		return "Wally", "West"
	} else {
		return "Jessy", "Quick"
	}
}

func main() {
	x := 2
	firstName, lastName := myFunction(x)
	fmt.Println("The name is " + firstName + " " + lastName)
}
