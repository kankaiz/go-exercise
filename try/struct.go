package main

import (
	"fmt"
	"math"
)

func main() {
	rec1 := Rectangle{10, 5}
	cir := Circle{7}

	// fmt.Println(rec1.height)
	fmt.Println(rec1.rarea())
	fmt.Println(rec1.area())
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
