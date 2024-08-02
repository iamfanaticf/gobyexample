package main

import "fmt"

func nextInt() func()int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

func main() {
	myCounter := nextInt()
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")

	fmt.Println()
	myCounter = nextInt()
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
	fmt.Print(myCounter(), " ")
}
