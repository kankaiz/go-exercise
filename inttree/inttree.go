package main

import "fmt"

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func main() {
	// need to define as pointer as the value will be updated
	var it *IntTree
	fmt.Println(it)
	it = it.Insert(3)
	it = it.Insert(6)
	fmt.Println(it)
}
