package main

import "fmt"

var stack []int

func EmptyStack() []int {
	return stack
}

func Push(inp ...int) {
	stack = append(stack, inp...)
}

func Pop(inp int) {
	temp := len(stack) - 1
	stack = stack[1:temp]
}

func main() {

	EmptyStack := EmptyStack()

	fmt.Println("Trying to commit :", EmptyStack)

}
