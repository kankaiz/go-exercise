package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x, "vs", &x)

	noChange(x)
	fmt.Println(x)

	change(&x)
	fmt.Println(x)

	var b *int
	// no memory allocated
	fmt.Println(b == nil)
	fmt.Println(b)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(*b)
	c := new(int)
	// memory allocated
	fmt.Println(c == nil)
	fmt.Println(c)
	// no panic
	// prints 0
	fmt.Println(*c)
}

func noChange(x int) {
	x = 7
}

func change(x *int) {
	*x = 7
}
