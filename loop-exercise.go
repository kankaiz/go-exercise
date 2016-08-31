package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z = float64(5)
	var p = float64(0)
	for math.Abs(p-z)>0.00001 {
		p = z
		z = z - ((z*z-x)/(2*z))
		//fmt.Println(z)
	}	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
