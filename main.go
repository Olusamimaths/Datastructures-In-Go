package main

import "github.com/olusamimaths/datastructures/stack"

func main() {
	cardStack := stack.NewCardStack()

	cardStack.StackCards()
	cardStack.ContainsCard("Ace of Spades")
	cardStack.ContainsCard("Spade of Spades")
	cardStack.DeckSize()
	// cardStack.UnstackCards()
}
