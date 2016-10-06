package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt ...
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",
		float64(e))
}

// Sqrt ...
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var z = float64(5)
	var p = float64(0)
	for math.Abs(p-z) > 0.00001 {
		p = z
		z = z - ((z*z - x) / (2 * z))
		//fmt.Println(z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2)) //print error object will try to find obj.Error()
}
