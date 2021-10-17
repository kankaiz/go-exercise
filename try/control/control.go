package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	defer firstRun()
	defer secondRun()

	s := "💪 🔥"
	// the offset is measured in bytes: 0,4,5
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
	// RuneCountInString counts characters
	fmt.Println(s, len(s), utf8.RuneCountInString(s))
}

func firstRun()  { fmt.Println("first") }
func secondRun() { fmt.Println("second") }
