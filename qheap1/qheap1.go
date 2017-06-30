// https://www.hackerrank.com/challenges/qheap1
package main

import "fmt"

func main() {
	var n int
	var t, v int
	fmt.Scanln(&n)
	q := make([]int, n)
	// q :=
	for i := 0; i < n; i++ {
		fmt.Scanln(&t, &v)
		if t == 1 {

		} else if t == 2 {

		} else {
			fmt.Println(q)
		}
	}
}
