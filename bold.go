package main

import "fmt"

func bold() {
	fmt.Printf("\x1b[1m")
}
func unbold() {
	fmt.Printf("\x1b[22m")
}
