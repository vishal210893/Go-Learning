package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, eachCard := range d {
		fmt.Println(i, " = "+eachCard)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	deckAsStringSlice := []string(d)
	deckAsString := strings.Join(deckAsStringSlice, ",")
	return deckAsString
}

func (d deck) saveToFIle(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	deckAsString := string(file)
	deckAsStringSlice := strings.Split(deckAsString, ",")
	return deck(deckAsStringSlice)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	for index := range d {
		newPos := random.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]
	}
}
