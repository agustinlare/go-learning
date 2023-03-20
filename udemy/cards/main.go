package main

func main() {
	cards := newDeck()
	cards.saveToFile("cartas.txt")
	cards = newDeckFromFile("cartas.txt")
	cards.print()
}
