package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	//Another Method
	for j := 11; j <= 20; j++ {
		fmt.Println(j * j)
	}
}
