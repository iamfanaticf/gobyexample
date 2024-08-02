package main

import "fmt"

func plus(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _,num := range nums {
		total += num
	}
	return total
}

func  main() {
	fmt.Println(plus(1,2,3,4,5))

	sl := []int {1,2,3,4,5,6}
	fmt.Println(plus(sl...))

}
