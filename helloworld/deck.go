package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Creating a type called 'deck' which is a Slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return an call to newDeck()
		// fmt.Println("Error:", err)
		// return newDeck()

		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fileContent := string(fileBytes)
	s := strings.Split(fileContent, ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		randomIndex := rand.Intn(len(d))
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}