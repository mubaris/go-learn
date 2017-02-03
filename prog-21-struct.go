package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Wally", 21})

	fmt.Println(person{name: "Jesse", age: 19})

	fmt.Println(person{name: "Jay"})

	fmt.Println(&person{name: "Max", age: 34})

	s := person{name: "Barry", age: 27}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 29
	fmt.Println(sp.age)
}
