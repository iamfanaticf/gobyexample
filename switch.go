package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fmt.Println("two 2nd time")
	case 3:
		fmt.Println("three")
	}
	
	fmt.Println(time.Now().Weekday())
	fmt.Println()
	
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its week end")
	case 9: 
		fmt.Println("its a case of Mondaysss")
	case time.Monday:
		fmt.Println("it works")
	default: 
		fmt.Println("hhehhee")
	}
	fmt.Println()

	// time.now
	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("Morning")
	default:
		fmt.Println("Not Morning")
	}
	fmt.Println()

	// type switch
	var r interface{}
	r = uint8(1) 
	fmt.Println(r)

	switch  r.(type) {
	case uint8: 
		fmt.Println("uint8")
	case int8:
		fmt.Println("int8")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("default")
	}

	whoAmI := func(i interface{}) {
	switch t := i.(type) {
	case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
	}
	}
	
	whoAmI(true)
	whoAmI(1)
	whoAmI('x')
	whoAmI("x")
}	
