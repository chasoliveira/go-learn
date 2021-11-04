package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	hand, remainDeck := deal(cards, 5)
	hand.print()
	remainDeck.print()

	fmt.Println(cards.toString())
	cardsFileName := "cards_file"
	errors := cards.saveToFile(cardsFileName)
	if errors != nil {
		fmt.Println(errors)
	}
	deckFromFile := newDeckFromFile(cardsFileName)
	deckFromFile.print()

	deckFromFile.shuffle()
	deckFromFile.print()
}
