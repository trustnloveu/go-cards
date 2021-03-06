/*
	Cards

	newDeck 		: Create a list of playing cards. Essentially an array of strings
	print			: Log out the contents of a deck of cards
	shuffle			: Shuffles all the cards in a deck
	deal			: Create a 'hand' of cards
	saveToFile		: Save a list of cards to a file on the local machine
	newDeckFromFile	: Load a list of cards from the local machine

*/

package main

func main() {
	//! String
	// var card string = "Ace of Spades"
	// card := newCard()

	//! Array & Slice
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	//! type
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards := newDeck()

	//! Iterate (for + slice)
	// for i, card := range cards {
	// 	fmt.Println(i, card);
	// }

	//! Byte Slice
	// greeting := "Hello, World"
	// fmt.Println([]byte(greeting))
	cards := newDeck()

	//! shuffle
	cards.shuffle()

	//! WriteFile
	cards.saveToFile("my_cards")

	//! ReadFile
	cards = newDeckFromFile("my_cards")

	//! Print
	// fmt.Println(cards)
	cards.print()
	// fmt.Println(cards.toString())
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
