package main

import "fmt"

func main() {
	defer firstRun()
	defer secondRun()
}

func firstRun()  { fmt.Println("first") }
func secondRun() { fmt.Println("second") }
