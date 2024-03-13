package main

import (
	"fmt"
	"os"
)

func main() {
	cards := newDeck()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println()
	remainingCards.print()
	cards.print()

	data, _ := os.ReadFile("deck.go")
	fmt.Printf("%T", data)

	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	newCards := newDeck()
	newCards.print()
	newCards.shuffle()
	newCards.print()

	newCards2 := newDeckFromFile("my_cards")
	newCards2.print()
}