package main

import "fmt"

func main() {
	/*
		channel for transafer data between goroutines
		one or more goroutines write to a channel
		one or more goroutines read back from the same channel
		in the ORDER which was written
	*/
	/*
		the data in the channel is typed
		by default, reads and writes in the channel are synchronous
		be careful when passing reference types over a channel
	*/

	in1 := make(chan string)
	out1 := make(chan string)
	go func() {
		name := <-in1 // read from in channel
		out1 <- fmt.Sprintf("Hello, " + name)
	}()
	go func() {
		name := <-in1 // read from in channel
		out1 <- fmt.Sprintf("Again, " + name)
	}()
	in1 <- "Bob"
	// in1 <- "another"
	// above will cause fatal error: all goroutines are asleep - deadlock!
	// as unbuffered channel pause further write
	close(in1)
	message := <-out1
	fmt.Println(message)

	// cf this block
	// go func() {
	// 	name := <-in1 // read from in channel
	// 	out1 <- fmt.Sprintf("Hello, " + name)
	// }()
	// go func() {
	// 	name := <-in1 // read from in channel
	// 	out1 <- fmt.Sprintf("Again, " + name)
	// }()
	// in1 <- "Bob"
	// in1 <- "another"
	// // also working if not close an unbuffered channel
	// // close(in1)
	// for i := 0; i < 2; i++ {
	// 	v := <-out1
	// 	fmt.Println(v)
	// }

	// By default, channels are unbuffered
	// every write to an open, unbuffered channel causes the writing goroutine
	// to PAUSE
	// if don't want to wait for a channel to be read
	// after we write to it
	// use buffer channel
	out2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(localI int) {
			out2 <- localI * 2
		}(i)
	}
	var result2 []int
	for i := 0; i < 10; i++ {
		val := <-out2
		result2 = append(result2, val)
	}
	// get the doubled i in a RANDOM order
	fmt.Println(result2)

	in3 := make(chan int, 10)
	out3 := make(chan int)

	for i := 0; i < 10; i++ {
		in3 <- i
	}
	// close means the channel will not have more values written to it
	// close does not wipe it's contents
	// if a buffer channel is closed, any value in the channel is available
	// if read from a closed channel, the read returns zero value for the type of channel
	close(in3)

	go func() {
		for {
			// check if read returns zero value or zero of channel type
			i, ok := <-in3
			if !ok {
				close(out3)
				break
			}
			out3 <- i * 2
		}
	}()

	// the loop continues until the channel is closed
	for v := range out3 {
		fmt.Println(v)
	}

	// the zero value for a channel is nil just like other reference types
}
