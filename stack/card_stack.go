package stack

import "fmt"

type CardStack interface {
	StackCards()
	UnstackCards()
	ContainsCard(card string)
	GoToCard(card string)
	DeckSize()
}
type cardStack struct {
	stack BasicStack[string]
}

func (c *cardStack) StackCards() {
	c.stack.Push("Ace of Spades")
	c.stack.Push("2 of Spades")
	c.stack.Push("3 of Spades")
	c.stack.Push("4 of Spades")
	c.stack.Push("5 of Spades")
	c.stack.Push("6 of Spades")
	c.stack.Push("7 of Spades")
	c.stack.Push("8 of Spades")
	c.stack.Push("9 of Spades")
	c.stack.Push("10 of Spades")
	c.stack.Push("Jack of Spades")
	c.stack.Push("Queen of Spades")
	c.stack.Push("King of Spades")

	//c.stack the diamond suit
	c.stack.Push("Ace of Diamonds")
	c.stack.Push("2 of Diamonds")
	c.stack.Push("3 of Diamonds")
	c.stack.Push("4 of Diamonds")
	c.stack.Push("5 of Diamonds")
	c.stack.Push("6 of Diamonds")
	c.stack.Push("7 of Diamonds")
	c.stack.Push("8 of Diamonds")
	c.stack.Push("9 of Diamonds")
	c.stack.Push("10 of Diamonds")
	c.stack.Push("Jack of Diamonds")
	c.stack.Push("Queen of Diamonds")
	c.stack.Push("King of Diamonds")

	//c.stack the club suit
	c.stack.Push("Ace of Clubs")
	c.stack.Push("2 of Clubs")
	c.stack.Push("3 of Clubs")
	c.stack.Push("4 of Clubs")
	c.stack.Push("5 of Clubs")
	c.stack.Push("6 of Clubs")
	c.stack.Push("7 of Clubs")
	c.stack.Push("8 of Clubs")
	c.stack.Push("9 of Clubs")
	c.stack.Push("10 of Clubs")
	c.stack.Push("Jack of Clubs")
	c.stack.Push("Queen of Clubs")
	c.stack.Push("King of Clubs")

	//c.stack the heart suit
	c.stack.Push("Ace of Hearts")
	c.stack.Push("2 of Hearts")
	c.stack.Push("3 of Hearts")
	c.stack.Push("4 of Hearts")
	c.stack.Push("5 of Hearts")
	c.stack.Push("6 of Hearts")
	c.stack.Push("7 of Hearts")
	c.stack.Push("8 of Hearts")
	c.stack.Push("9 of Hearts")
	c.stack.Push("10 of Hearts")
	c.stack.Push("Jack of Hearts")
	c.stack.Push("Queen of Hearts")
	c.stack.Push("King of Hearts")
}

func (c *cardStack) UnstackCards() {
	for c.stack.Size() > 0 {
		fmt.Println(c.stack.Pop())
	}
}

func (c *cardStack) ContainsCard(card string) {
	fmt.Println(c.stack.Contains(card))
}

func (c *cardStack) GoToCard(card string) {
	fmt.Println(c.stack.Access(card))
}

func (c *cardStack) DeckSize() {
	fmt.Println(c.stack.Size())
}

func NewCardStack() CardStack {
	return &cardStack{
		stack: NewStack[string](1000),
	}
}
