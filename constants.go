package main

import (
	"fmt"
	"math"
)

func main() {
	const a = 1
	fmt.Printf("%T \n", a)
	const b = 3e20 / 50000
	fmt.Println(b)
	fmt.Printf("%T \n", b) // a numeric constant has no type. type is decided by context
	fmt.Println(math.Cos(a)) // here cos expected a float64 type but int was given. cos a const has no type actually it is being decided by the context in which it is placed
}
