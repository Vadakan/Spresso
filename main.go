package main

import "fmt"

var stack []int

func EmptyStack() []int {
	return stack
}

func main() {

	EmptyStack := EmptyStack()

	fmt.Println(EmptyStack)

}
