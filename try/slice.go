package main

import "fmt"

func main() {
	slice := []int{5, 4, 3, 2, 1}

	old := slice[:]
	fmt.Println(old)

	slice[3] = 9
	fmt.Println(slice)
	fmt.Println(old)

	// var cp []int  this will make copy result to empty list
	cp := make([]int, len(slice))
	copy(cp, slice[:])
	slice[0] = 8
	fmt.Println(slice)
	fmt.Println(cp)

	var appd []int
	appd = append(appd, slice[:]...)
	fmt.Println(appd)
	slice[4] = 7
	fmt.Println(slice)
	fmt.Println(appd)
}
