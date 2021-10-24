package main

import (
	"fmt"
	"time"
)

func main() {
	in1 := make(chan int)
	out1 := make(chan int, 1)
	// program not block as out is buffered
	out1 <- 1
	// program blocks as no goroutine reads from in
	// in <- 2
	// fmt.Println("wrote 2 to in")
	// v := <-out
	// fmt.Println("read", v, "from out")

	// program not blocked as select pick up the unblocked case
	// will always run  v := <-out:
	select {
	case in1 <- 2:
		fmt.Println("wrote 2 to in1")
	case v := <-out1:
		fmt.Println("read", v, "from out1")
	}

	in2 := make(chan int, 1)
	out2 := make(chan int, 1)

	out2 <- 1
	// neither case will block
	// select random pick up one to run
	select {
	case in2 <- 2:
		fmt.Println("wrote 2 to in2")
	case v := <-out2:
		fmt.Println("read", v, "from out2")
	}

	in3 := make(chan int)
	out3 := make(chan int)
	// both are blocked so only default case will be picked up
	// if any case can run, default will not be picked up
	select {
	case in3 <- 2:
		fmt.Println("wrote 2 to in3")
	case v := <-out3:
		fmt.Println("read", v, "from out3")
	default:
		fmt.Println("nothing else works")
	}

	// recycle channel when it is finished
	// assign multiples(2) to twosCh
	// assign empty struct to done channel
	twosCh, done := multiples(2)
	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}
	close(done)

	//do more stuff
	time.Sleep(1 * time.Second)
}

func multiples(i int) (chan int, chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	curVal := 0
	go func() {
		for {
			select {
			case out <- curVal * i:
				curVal++
			case <-done:
				close(out) // need that?
				fmt.Println("goroutine shutting down")
				return
			}
		}
	}()
	return out, done
}
