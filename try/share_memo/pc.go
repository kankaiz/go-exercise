package sharememo

import (
	"fmt"
	"time"
)

type Task struct {
	A int
	B int
}

var ch = make(chan Task, 100)

func producer() {
	for i := 0; i < 10; i++ {
		ch <- Task{A: i + 3, B: i - 1}
	}
}

func consumer() {
	for i := 0; i < 10; i++ {
		task := <-ch
		sum := task.A + task.B //2i+2
		fmt.Println(sum)
	}
}

func main2() {
	go producer()           // producer routine
	go consumer()           // consumer routine
	time.Sleep(time.Second) // expected 2i+2
}
