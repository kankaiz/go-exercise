package main

import "fmt"

func main() {
	var a int = 5
	var b float32 = 23.235
	const pi float64 = 3.14159265354
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)

	x, y := 13, 14
	fmt.Println("x+y =", x+y)
	fmt.Println("x-y =", x-y)
	fmt.Println("x*y =", x*y)
	fmt.Println("x/y =", x/y)
	fmt.Println("x mod y =", x%y)

	// floor div and math div
	fmt.Println(3/4, "vs", float32(3/4), "vs", float32(3)/float32(4))
	// multi 2 same as Python
	fmt.Println(7, "*2 ops", 7<<1)

	TRUE, FALSE := true, false
	fmt.Println(TRUE && FALSE, TRUE || FALSE, !TRUE, !FALSE)
}
