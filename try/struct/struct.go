package main

import (
	"fmt"
	"math"
)

// Capped Foo is public struct
type Foo struct {
	A int
	b string
}

func main() {
	// Strusts are not Objects!
	f := Foo{
		A: 20,
	}
	//f2 full copy of f
	var f2 Foo
	f2 = f
	f2.A = 100
	fmt.Println(f2.A, f.A)

	//f3 points to the same memory of f
	var f3 *Foo = &f
	f3.A = 200
	fmt.Println(f3.A, f.A)

	rec1 := Rectangle{10, 5}
	cir := Circle{7}

	// fmt.Println(rec1.height)
	fmt.Println(rec1.rarea(), rec1.area())
	// fmt.Println(getArea(rec1))
	// cannot use rec1 (variable of type Rectangle) as Shape value in argument to getArea:
	// missing method area (area has pointer receiver)
	fmt.Println(getArea(&rec1))
	fmt.Println(getArea(cir))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (rec *Rectangle) area() float64 {
	return rec.height * rec.width
}

func (rec *Rectangle) rarea() float64 {
	return rec.height * rec.width
}

type Circle struct {
	radius float64
}

func (cir Circle) area() float64 {
	return math.Pi * math.Pow(cir.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}
