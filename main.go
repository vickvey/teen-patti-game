package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Suit represents the suit of a playing card
type Suit int

// Rank represents the rank of a playing card
type Rank int

// Card represents a playing card with a suit and a rank
type Card struct {
	Suit
	Rank
}

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Map to store the image stickers or emojis for each suit
var suitImages = map[Suit]string{
	Spades:   "♠️",
	Hearts:   "♥️",
	Diamonds: "♦️",
	Clubs:    "♣️",
}

func (card *Card) DisplayCard() {
	// Determine text color based on the card's suit
	var c *color.Color
	switch card.Suit {
	case Hearts, Diamonds:
		c = color.New(color.FgRed)
	default:
		c = color.New(color.FgHiBlack)
	}

	// Print the card with the appropriate color
	c.Printf("%s %s %s\n", card.Suit, suitImages[card.Suit], card.Rank)
}

func (s Suit) String() string {
	suits := [...]string{"Spades", "Hearts", "Diamonds", "Clubs"}
	if s < Spades || s > Clubs {
		return "Unknown"
	}
	return suits[s]
}

func (r Rank) String() string {
	ranks := [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	if r < Ace || r > King {
		return "Unknown"
	}
	return ranks[r-1]
}

func GenerateDeck() []Card {
	var deck []Card
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for _, rank := range []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King} {
			card := Card{Suit: suit, Rank: rank}
			deck = append(deck, card)
		}
	}
	return deck
}

func ShuffleDeck(deck []Card) []Card {
	rand.NewSource(time.Now().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

// DealCard deals a single card from the deck and returns it along with the remaining deck.
// It shuffles the deck to ensure randomness before dealing the card.
func DealCard(deck []Card) (Card, []Card) {
	deck = ShuffleDeck(deck)
	return deck[0], deck[1:]
}

func main() {
	fmt.Println("Hello World!!")

	deck := GenerateDeck()
	card, _ := DealCard(deck)
	card.DisplayCard()
}
