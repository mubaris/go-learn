package main

import "fmt"

func main() {
	for i := 1; i <= 25; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Divisible by 2")
		} else if i%3 == 0 {
			fmt.Println(i, "Divisible by 3")
		} else if i%5 == 0 {
			fmt.Println(i, "Divisible by 5")
		} else {
			fmt.Println("Foo, this doesn't work for me!! LOLzz")
		}
	}
}
