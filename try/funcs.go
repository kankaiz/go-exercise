package main

import "fmt"

func main() {
	// local anonymous func to modify local value
	b := 2

	// local func called closures
	myAddOne := func(a int) {
		b = a + b
	}

	myAddOne(1)
	fmt.Println(b)
}
