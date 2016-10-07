// https://www.hackerrank.com/challenges/arrays-ds
package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scanln(&n)
	array := make([]int, n)
	for i := range array {
		fmt.Scanf("%v", &array[i]) // or use %d
	}
	for ; n > 0; n-- {
		fmt.Printf("%d", array[n-1])
		fmt.Print(" ")
	}
	fmt.Printf("\n")

}
