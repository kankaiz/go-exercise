// https://www.hackerrank.com/challenges/array-left-rotation
package main

import (
	"fmt"
	"strings"
)

// rotate ...
func rotate(array []int, r int) []int {
	length := len(array)
	rl := r % length
	return append(array[rl:], array[:rl]...)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var len, r int
	fmt.Scanf("%v %v", &len, &r)
	// array := []int{}
	array := make([]int, len)
	for i := range array {
		fmt.Scanf("%v", &array[i])
	}
	// fmt.Println(rotate(array, r))
	// this way to trim brackets
	fmt.Println(strings.Trim(fmt.Sprint(rotate(array, r)), "[]"))
	// for i := 0; i < len; i++ {
	// 	fmt.Printf("%v ", array[(i+r)%len])
	// }
}
