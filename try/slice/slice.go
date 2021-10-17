package main

import "fmt"

func main() {
	slice := []int{5, 4, 3, 2, 1}

	fmt.Println("try slice copy")
	old := slice[:]
	fmt.Println(old)

	slice[3] = 9
	fmt.Println(slice, old)

	fmt.Println("try copy")
	// var cp []int  this will make copy result to empty list
	cp := make([]int, len(slice))
	// copy() will match the dst slice length
	copy(cp, slice[:])
	slice[0] = 8
	fmt.Println(slice, cp)

	fmt.Println("*try first append copy")
	var appd []int
	appd = append(appd, slice[:]...)
	fmt.Println(slice, appd)
	slice[4] = 7
	fmt.Println(slice, appd)
	fmt.Println("try second append copy")
	appd2 := append(slice)
	fmt.Println(slice, appd2)
	slice[1] = 2
	fmt.Println(slice, appd2)
	fmt.Println("try third append copy")
	var appd3 []int
	appd3 = append(slice[:0], slice...)
	// appd3 := append(slice[:0], slice...)
	fmt.Println(slice, appd3)
	slice[2] = 5
	fmt.Println(slice, appd3)
	fmt.Println("*try forth append copy")
	appd4 := append([]int(nil), slice...)
	fmt.Println(slice, appd4)
	slice[0] = 6
	fmt.Println(slice, appd4)

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
