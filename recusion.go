package main

import "fmt"

func fact(x int) int {
	if x <= 1 {
		return 1
	}
	return x * fact(x-1)
}

func main() {
	fmt.Println(fact(5))

	var fib func(int)int 

	fib = func(x int) int {
		if x < 2 {
			return x
		}
		return fib(x-1) + fib(x-2)
	}
	fmt.Println(fib(7))
}
