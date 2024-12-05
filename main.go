package main

import "fmt"

func main() {
	/*
		var cards string = "Ace of Spades"
		var cardsWithoutType = "Cards"
		cardsWithoutVar := "Without_var_card"

		card := newCards()
		fmt.Println(cards, cardsWithoutType, cardsWithoutVar)
		fmt.Println(card)

		multipleCards := deck{"Ace of Diamond", "Ace of Spade"}
		multipleCards = append(multipleCards, "Six of Spades")
		multipleCards.print()
	*/

	cards := newDeck()
	//cards.print()

	hand, remainingCards := deal(cards, 2)
	hand.print()
	fmt.Println("------------------")

	remainingCards.print()
	fmt.Println("------------------")

	// IMPLEMENT TOSTRING TO CONVERT DECK TO STRING
	fmt.Println(hand.toString())

	// WRITE TO FILE
	err := cards.saveToFIle("my_cards")
	if err != nil {
		return
	}
	fmt.Println("------------------")

	cardsFromFile := newDeckFromFile("my_cards")
	cardsFromFile.print()
	fmt.Println("------------------")

	cardsFromFile.shuffle()
	cardsFromFile.print()

}

func newCards() string {
	return "Five of Diamonds"
}
