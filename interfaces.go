package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	l, b int
}

type circle struct {
	r int
}

func (r *rect) area() float64 {
	return float64(r.l * r.b)
}

func (r *rect) perim() float64 {
	return  float64(2 * r.l + 2 * r.b)
}

func (c *circle) area() float64 {
	return float64( c.r * c.r) * math.Pi
}

func (c *circle) perim() float64 {
	return float64( 2 * c.r) * math.Pi
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{3, 4}
	c := circle{5}
	
	measure(&r)
	measure(&c)
}

