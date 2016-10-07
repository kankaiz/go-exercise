// https://www.hackerrank.com/challenges/2d-array
package main

import (
	"fmt"
	"math"
)

type Context [6][6]int

func CalHourGlass(array Context, x int, y int) int {
	if (x+3) <= 6 && (y+3) <= 6 {
		sum := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				sum += array[y+j][x+i] //pay attention to the matix index order
			}
		}
		sum = sum - array[y+1][x] - array[y+1][x+2]
		return sum
	}
	return math.MinInt32

}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// array := make([][]int, [6,6])
	var array Context
	for i := range array {
		for j := range array[i] {
			fmt.Scanf("%d", &array[i][j])
		}
	}
	var max, temp int
	max = math.MinInt32

	for i := 0; i < 6-2; i++ {
		for j := 0; j < 6-2; j++ {
			temp = CalHourGlass(array, i, j)
			if temp > max {
				max = temp
			}
		}
	}
	fmt.Println(max)
}
