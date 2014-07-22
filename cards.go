// Package cards provides methods for common operations on a deck of cards
package cards

import (
	"fmt"
	"math/rand"
	"time"
)

// The possible suits to use. By default this is "Spades", "Hearts", "Diamonds", "Clubs"
var Suits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}

// The number of ranks (values) per suit. By default this is 13
var Ranks = 13

// Deck represents a deck of cards
type Deck []Card

// NewDeck returns a new deck of cards, in order
//
// One card will be included of each suit and rank, for a total of Ranks*len(Suits) cards
func NewDeck() Deck {
	nCards := len(Suits) * Ranks
	d := make([]Card, nCards)
	for i := 0; i < nCards; i++ {
		d[i] = Card(i)
	}
	return d
}

// Shuffle randomizes the order of the cards in the deck
func (d Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		j := r.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

// Cut splits the deck into two decks of equal size
func (d Deck) Cut() (Deck, Deck) {
	// TODO: Support cutting into >2 decks
	h := len(d) / 2
	return d[:h], d[h:]
}

// Top deals the top card from the deck, removing it from the deck
func (dp *Deck) Top() Card {
	d := *dp
	c := d[0]
	if len(d) == 1 {
		d = nil
	} else {
		d = d[1:]
	}
	*dp = d
	return c
}

// TopN deals the top n cards from the deck
//
// If the deck does not have n cards, it will return all its cards and become empty
func (dp *Deck) TopN(n int) []Card {
	c := []Card{}
	for i := 0; i < n && !dp.Empty(); i++ {
		c = append(c, dp.Top())
	}
	return c
}

// Add adds card(s) to the bottom of the deck
func (dp *Deck) Add(cards ...Card) {
	d := *dp
	d = append(d, cards...)
	*dp = d
}

// Empty returns whether the deck is empty
func (d Deck) Empty() bool {
	return len(d) == 0
}

// Card represents a Card in a Deck
type Card int

// Suit returns the suit of the card
func (c Card) Suit() Suit {
	return Suit(int(c) / Ranks)
}

// Rank returns the value of the card, for comparison with the rank of other cards.
func (c Card) Rank() int {
	return int(c) % Ranks
}

// String returns a string representation of the card, e.g., "4 of Hearts", "Ace of Spades"
//
// With the default value of Ranks, the rank of the card will be returned starting at 2 until 10, then Jack, Queen, King, Ace.
//
// If Ranks has been modified from its default, the zero-indexed rank will be returned instead.
func (c Card) String() string {
	v := c.Rank()
	s := ""
	if Ranks != 13 {
		return fmt.Sprintf("%d", v)
	}
	if v < 9 {
		s = fmt.Sprintf("%d", v+2)
	} else {
		switch v {
		case 9:
			s = "Jack"
		case 10:
			s = "Queen"
		case 11:
			s = "King"
		case 12:
			s = "Ace"
		default:
			panic("bad value")
		}
	}
	return fmt.Sprintf("%s of %s", s, c.Suit())
}

// Suit represents the Suit of the card, one of Suits
type Suit int

// String returns the string representation of the card, e.g., "Hearts"
func (s Suit) String() string {
	return Suits[s]
}
