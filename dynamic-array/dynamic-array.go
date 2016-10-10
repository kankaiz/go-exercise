// https://www.hackerrank.com/challenges/dynamic-array
package main

import "fmt"

func main() {
	var n, queryCount int
	fmt.Scanf("%v %v", &n, &queryCount)
	// make the dynamica array
	array := make([][]int, n)

	// do not use make to initialise array otherwise it will assign 0 to it
	// for i := range array {
	// 	// array[i] = make([]int, size)
	// }

	// use declaration will not specify the array index which will cause index out of range
	// array := [][]int{}
	// var array [][]int

	// initialise lastAns
	lastAns := 0
	// get Query
	var q, x, y int
	var tempN int
	for i := 0; i < queryCount; i++ {
		fmt.Scanf("%v %v %v", &q, &x, &y)
		if q == 1 {
			// deal with append
			// tempN = math.Mod(x^lastAns, n)
			tempN = (x ^ lastAns) % n
			array[tempN] = append(array[tempN], y)
		} else if q == 2 {
			tempN = (x ^ lastAns) % n
			size := len(array[tempN])
			lastAns = array[tempN][y%size]
			fmt.Println(lastAns)
		}
	}
}
