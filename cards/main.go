package main

func main() {
	//var card = "ace of spades"
	//or
	//var cards deck
	cards := newDeck()
	cards.print()

	d1, d2 := deal(cards, 2)

	d1.print()
	d2.print()

	cards.writeTOFile("tempfile.txt")
	readFromFile("tempfile.txt")
}

func newCard() string {
	return "ace of spades"
	//return 4
}
