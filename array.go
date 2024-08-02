package main

import "fmt"

func main() {
	var a[5] int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println("len ", len(a))

	b := [4]int {1,2,3,4}
	fmt.Println(b)

	c := [...]int {4,4,4,4}
	fmt.Println(c)

	d := [...]int {1, 2:9, 3, 4}
	fmt.Println(d)

	var e [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			e[i][j] = i + j
		}
	}
	fmt.Println(e)

	f := [3][3]int {
		{1, 2, 3},
		{4, 5, 6},
		{8, 9, 0},
	}
	fmt.Println(f)
}

