package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// Capped Foo is public struct
type Foo struct {
	A int
	b string
}

type Bar struct {
	C string
	F Foo
}

type Baz struct {
	D string
	// embedding
	Foo
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Strusts are not Objects!
	f := Foo{
		A: 20, b: "Hello",
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

	// try embedding structs
	fmt.Println("try embedding structs")
	b1 := Bar{C: "Fred", F: f}
	b2 := Baz{D: "Nancy", Foo: f}
	fmt.Println(b1.F.A, b2.A)

	f2 = b2.Foo
	fmt.Println(f2)

	rec1 := Rectangle{10, 5}
	cir := Circle{7}

	// fmt.Println(rec1.height)
	fmt.Println(rec1.rarea(), rec1.area())
	// fmt.Println(getArea(rec1))
	// cannot use rec1 (variable of type Rectangle) as Shape value in argument to getArea:
	// missing method area (area has pointer receiver)
	fmt.Println(getArea(&rec1))
	fmt.Println(getArea(cir))

	// try json marshal and unmarshal
	fmt.Println("try json marshal")
	bob := `{ "name": "Bob", "age": 30}`
	var b Person
	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b)
	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
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
