package sharememo

import (
	"fmt"
	"sync"
	"time"
)

var mp sync.Map

func rwGlobalMemory() {
	if val, ok := mp.Load("mykey"); ok {
		fmt.Println(val)
	} else {
		mp.Store("mykey", "myval")
	}
}

func main1() {
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()

	time.Sleep(time.Second)
	// may print 2 or 3. unexpected behavior
	// bcuz run rwGlobalMemory concurrently
}
