package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x, "vs", &x)

	noChange(x)
	fmt.Println(x)

	change(&x)
	fmt.Println(x)
}

func noChange(x int) {
	x = 7
}

func change(x *int) {
	*x = 7
}
