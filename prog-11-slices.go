package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	newSlice1 := []int{5, 6, 7, 8, 9}
	newSlice2 := make([]int, 3)
	copy(newSlice2, newSlice1)
	fmt.Println(newSlice1, newSlice2)
}
