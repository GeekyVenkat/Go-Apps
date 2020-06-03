package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type rectangle struct {
	width  float32
	length float32
}

type shape interface {
	area() float32
}

func (c *circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r *rectangle) area() float32 {
	r.length = 6
	return r.length * r.width
}

func area(s shape) {
	fmt.Println("Area :", s.area())
}
func main() {

	c := circle{
		radius: 6.5,
	}

	r := rectangle{
		length: 5,
		width:  4,
	}

	//c1 := &c
	// fmt.Println(c1.radius)
	//fmt.Println(r.length)

	area(&c)
	area(&r)
	fmt.Println(r.length)

}
