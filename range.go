package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _,val := range nums {
		sum += val
	}
	fmt.Println(sum)

	for i, val := range nums {
		if i == 3 {
			fmt.Println(val)
		}
	}

	m := map[int]string {1:"one", 2:"two", 3:"three"}
	for k, v := range m {
		fmt.Println("key ", k, " value ", v)
	}

	fmt.Println()
	for k := range m {
		fmt.Printf("%v ", k)
	}
	fmt.Println()

	for _,c := range "golang" {
		fmt.Printf(" %v ", c)
	}
}
