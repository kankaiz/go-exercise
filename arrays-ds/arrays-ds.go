package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scanln(&n)
	array := make([]int, n)
	for i := range array {
		fmt.Scanf("%d", &array[i])
	}
	for ; n > 0; n-- {
		fmt.Printf("%v", array[n-1])
		fmt.Print(" ")
	}
	fmt.Printf("\n")

}
