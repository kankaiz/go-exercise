package main

import "fmt"

func main() {
	var a int = 5
	var b float32 = 23.235
	const pi float64 = 3.14159265354

	x, y := 13, 14

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)

	fmt.Println(x, y)
	fmt.Println("x+y =", x+y)
}
