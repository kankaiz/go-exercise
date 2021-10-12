package main

import "fmt"

func main() {
	slice := []int{5, 4, 3, 2, 1}

	fmt.Println("try slice copy")
	old := slice[:]
	fmt.Println(old)

	slice[3] = 9
	fmt.Println(slice)
	fmt.Println(old)

	fmt.Println("try copy")
	// var cp []int  this will make copy result to empty list
	cp := make([]int, len(slice))
	// copy() will match the dst slice length
	copy(cp, slice[:])
	slice[0] = 8
	fmt.Println(slice)
	fmt.Println(cp)

	fmt.Println("try append copy")
	var appd []int
	appd = append(appd, slice[:]...)
	fmt.Println(slice)
	fmt.Println(appd)
	slice[4] = 7
	fmt.Println(slice)
	fmt.Println(appd)

	// insert ele
	fmt.Println("try insertion")
	slice1 := append(slice[:2], 0)
	slice1 = append(slice1, slice[2:]...)
	// here slice HAS changed
	fmt.Println(slice)
	fmt.Println(slice1)
	// cmp
	slice2 := append(slice[:2], slice[1:]...)
	slice2[1] = 3
	// here slice hasn't changed
	fmt.Println(slice)
	fmt.Println(slice2)
}
