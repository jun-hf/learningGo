package main

import "fmt"

type Vertex struct {
	area int
}

func (v *Vertex) scale() {
	v.area = 90
	fmt.Println(v.area)
}

type Circle struct {
	radius int
}

func (c Circle) scale() {
	c.radius = 90
	fmt.Println(c.radius)
}

func Ver() {
	a := Vertex{80}
	ap := &Vertex{1}

	a.scale() // &a.scale compiler will assume this way
	ap.scale() 

	fmt.Println(a)
	fmt.Println(ap)

	// calling a value receiver function

	c1 := Circle{11}
	c2 := &Circle{89}

	c1.scale()
	c2.scale()

	fmt.Println(c1)
	fmt.Println(c2)

}