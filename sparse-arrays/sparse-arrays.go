//https://www.hackerrank.com/challenges/sparse-arrays
package main

import "fmt"

func parseArray(array []string, word string) int {
	var count int
	for i := range array {
		if array[i] == word {
			count++
		}
	}
	return count
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, q int
	fmt.Scanln(&n)
	array := make([]string, n)
	for i := range array {
		fmt.Scanln(&array[i])
	}
	fmt.Scanln(&q)
	for i := 0; i < q; i++ {
		var temp string
		fmt.Scanln(&temp)
		fmt.Println(parseArray(array, temp))
	}
}
