package main

import "fmt"

const (
	message = "Hello World! The answer is %d\n"
	one     = iota
	two
	three
)

var (
	answer = 42
)

func main() {
	fmt.Printf(message, answer)
	// answer = iota
	// answer = iota
	fmt.Printf(message, answer)
	fmt.Printf(message, one)
	fmt.Printf(message, two)
	fmt.Printf(message, three)
}
