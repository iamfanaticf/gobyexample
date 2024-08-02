package main

import (
	"fmt"
	"maps"
)

func main() {
	var m map[int]string
	fmt.Println(m, m==nil)
	_, ok := m[0]
	fmt.Println("is ok?", ok)

	m2 := make(map[string]int)
	m2["k1"] = 7
	m2["k2"] = 9

	fmt.Println(m2)

	v1 := m2["k1"]
	fmt.Println(v1) 

	v3 := m2["k3"]
	fmt.Println(v3)

	fmt.Println(m2)
	fmt.Println(len(m2))

	delete(m2, "k2")
	fmt.Println(m2)
	fmt.Println(len(m2))

	clear(m2)
	fmt.Println(m2)
	fmt.Println(len(m2))
	
	_,ok = m2["k1"]
	fmt.Println("is ok?", ok)

	n := map[string]int { "k1":1, "k2":2, }
	n2 := map[string]int { "k1":1, "k2":2}

	fmt.Println(maps.Equal(n, n2))


}

