package main

import "fmt"

func main() {
	defer firstRun()
	defer secondRun()

	s := "💪 🔥"
	// the offset is measured in bytes: 0,4,5
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}

func firstRun()  { fmt.Println("first") }
func secondRun() { fmt.Println("second") }
