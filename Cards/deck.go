package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck--> []string
//this is a slice of type string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
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

func (d deck) tostring() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.tostring()), 0666)

}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		//op1: Log the error & return the call
		//op2: log the error & entirely quit the program
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")

	return deck(s)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
