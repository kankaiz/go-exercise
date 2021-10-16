package main

import "fmt"

func main() {
	rec1 := Rectangle{10, 5}

	fmt.Println(rec1.height)
	fmt.Println(rec1.area())
}

type Rectangle struct {
	height float64
	width  float64
}

func (rec *Rectangle) area() float64 {
	return rec.height * rec.width
}
