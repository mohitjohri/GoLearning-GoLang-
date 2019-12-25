package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// fmt.Println("----------------------------")
	// deckOne, deckTwo := deal(cards, 5)

	// deckOne.print()
	// fmt.Println("----------------------------")

	// deckTwo.print()
	// s := deckTwo.tostring()
	// fmt.Println("----------Card to string-----------------")

	// fmt.Println(s)

	// fmt.Println("----------Save to file-----------------")

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

}
