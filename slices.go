package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println(s==nil, len (s))
	
	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	fmt.Println(s, len(s), cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	s[0] = "z"
	fmt.Println(c)
	fmt.Println(s)

	l := s[:2]
	fmt.Println(l)

	t := []string {"g", "h", "i"}
	t2 := []string {"g", "h", "i"}

	if slices.Equal(t, t2) {
		fmt.Println("slices are equal")
	}

	twoD := make([][]int, 3)
	
	for i := range 3 {
		twoD[i] = make([]int, 3)
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}	
