package main

import "fmt"

func main() {
	// basic if
	if 7 % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// if without else
	if 7 % 2 == 0 {
		fmt.Println("odd")
	}

	// if with init statement
	if n:=10; n<100 {
		fmt.Println("less than 100")
	}
	// if with paren
	if (1 == 1) {
		fmt.Println("works")
	}
	// for with paren
}	
