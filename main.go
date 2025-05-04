package main

func main() {
	cards := newDeck()

	cards.print()
	cards.shuffle()
	println("...........")
	cards.print()

}
