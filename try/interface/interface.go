package main

import "fmt"

type tester interface {
	// list all methods
	// method name(parameter types) return type
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

type rangeTest struct {
	min int
	max int
}

// rangeTest implement test method
// !!!no explicit interface declaration!!!
func (rt rangeTest) test(i int) bool {
	return rt.min <= i && i <= rt.max
}

type divTest int

// divTest implement test method
func (dt divTest) test(i int) bool {
	return i%int(dt) == 0
}

// declare a testerFunc which is a func
// that takes int and returns bool
type testerFunc func(int) bool

// as test(i int) bool takes int and returns bool
// it implement the interface tester
// also as testerFunc is an interface returns bool
// can call testerFunc
func (tf testerFunc) test(i int) bool {
	return tf(i)
}

func main() {
	// rangeTest against 10
	// and divTest against 10
	result := runTests(10, []tester{
		rangeTest{min: 5, max: 20},
		divTest(5),
	})
	fmt.Println(result)

	result2 := runTests(10, []tester{
		testerFunc(func(i int) bool {
			return i%2 == 0
		}),
		testerFunc(func(i int) bool {
			return i < 20
		}),
	})
	fmt.Println(result2)

	fmt.Println("empty interface")
	var i interface{}
	i = "Hello"
	// type assertion
	// can only do type assertion on interface
	j := i.(string)
	k, ok := i.(int)
	fmt.Println(j, k, ok)
	m := i.(int)
	fmt.Println(m)
}
