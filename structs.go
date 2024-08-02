package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name:name, age:age}
	return &p
}

func main() {
	fmt.Println(*newPerson("Jon", 21))

	anon := struct {
		name int
		age int
	} {
		0,
		1,
	}
	fmt.Println(anon)
}
