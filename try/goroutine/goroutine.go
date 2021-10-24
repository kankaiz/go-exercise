package main

import (
	"fmt"
	"sync"
)

func runMe(name string) {
	fmt.Println("Hello to", name, "from a goroutine")
}

func main() {
	/*
		goroutines are lightweight processes managed by Go runtime scheduler
		cf threads are managed by OS
		goroutines are faster to create, with smaller stack sizes
		and faster to switch
		can have tens of thounsands in a single process
	*/

	// this will not print anything as main goroutine not wait other to finish
	// go runMe()

	// we should use wait group rather than pause the main goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func(name string) {
		runMe(name)
		wg.Done()
	}("AAA")
	// pause the main goroutine to wait until
	// the added count in wait group reach zero

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// go func() {
		// 	// all the goroutines are sharing the same var i
		//  // so will print out 10 for most times
		// 	fmt.Println(i)
		// 	wg.Done()
		// }()
		go func(localI int) {
			fmt.Println(localI)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
