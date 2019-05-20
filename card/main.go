package main

func main() {
	//cards := newDeck()

	//hand, remainCards := deal(cards, 5)

	//hand.print()
	//remainCards.print()

	//fmt.Println(cards.toString())
	//cards.toSaveFile("My_file")

	//cards := newDeckFile("My_file")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
