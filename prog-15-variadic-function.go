package main

import "fmt"

func stringGenerator(args ...string) string {
	newString := ""
	for _, x := range args {
		newString += x + " "
	}
	return newString
}

func main() {
	fmt.Println(stringGenerator("Barry", "Allen", "is", "the", "Flash"))
}
