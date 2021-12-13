package main

import "fmt"

var stack []int

func EmptyStack() []int {
	return stack
}

func Push(inp ...int) {
	stack = append(stack, inp...)
}

func main() {

	EmptyStack := EmptyStack()

	fmt.Println("Trying to commit :", EmptyStack)
	fmt.Println("Try to pushh...")
}
