package main

import (
	"fmt"
	"sync"
)

func runMe(name string) {
	fmt.Println("Hello to", name, "from a goroutine")
}

func main() {
	var wg sync.WaitGroup
	// tell WaitGroup to add 1 goroutine to run
	wg.Add(1)
	go func(name string) {
		runMe(name)
		wg.Done()
	}("AAA")

	wg.Wait()

	// wrong way
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(localI int) {
			fmt.Println(localI)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
