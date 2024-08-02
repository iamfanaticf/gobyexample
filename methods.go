package main

import "fmt"

type rect struct {
	Length		int
	Width		int
}

func (r *rect) area() int {
	return r.Length * r.Width
}

func (r rect) perim() int {
	return r.Length + r.Width
}

func main() {
	r := rect { Width: 20, Length: 10 }
	fmt.Println(r.area())
	fmt.Println(r.perim())
}
