package main

import "fmt"

const (
	message = "Hello World! The answer is %d\n"
)

var (
	answer = 42
)

func main() {
	fmt.Printf(message, answer)
}
