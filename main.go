package main

import "github.com/olusamimaths/datastructures/stack"

func main() {
	myStack := stack.NewCardStack()

	myStack.StackCards()

	myStack.UnstackCards()
}
