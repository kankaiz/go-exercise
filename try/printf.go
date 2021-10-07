package main

import "fmt"

func main() {
	var sss string = "This is a string"
	const pi float64 = 3.14159265354
	ok := true
	x := 5

	fmt.Println(len(sss))
	fmt.Println(sss + " print")
	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)
	// print type
	fmt.Printf("%T \n", pi)
	// print boolean
	fmt.Printf("%t \n", ok)
	// print int
	fmt.Printf("%d \n", x)
	// print binary
	fmt.Printf("%b \n", 25)
	// chr
	fmt.Printf("%c \n", 34)
}
