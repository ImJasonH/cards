// Package cards provides methods for common operations on a deck of cards
package cards

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// The possible suits to use. By default this is {'♠', '♡', '♢', '♣'}
var Suits = []rune{'♠', '♡', '♢', '♣'}

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
	return d[h:], d[:h]
}

// Combines two decks using a strictly mechanical riffle, alternating a card from each deck
func Riffle(d1, d2 Deck) Deck {
	var d Deck
	for !d1.Empty() && !d2.Empty() {
		d.Add(*d1.Top(), *d2.Top())
	}
	d.Add(d1...)
	d.Add(d2...)
	return d
}

// Top deals the top card from the deck, removing it from the deck
func (dp *Deck) Top() *Card {
	d := *dp
	if d == nil {
		return nil
	}
	c := d[0]
	if len(d) == 1 {
		d = nil
	} else {
		d = d[1:]
	}
	*dp = d
	return &c
}

// TopN deals the top n cards from the deck
//
// If the deck does not have n cards, it will return all its cards and become empty
func (dp *Deck) TopN(n int) []Card {
	c := []Card{}
	for i := 0; i < n && !dp.Empty(); i++ {
		c = append(c, *dp.Top())
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

// Sort sorts the deck into the same order it was in when it was created
func (dp *Deck) Sort() {
	var is []int
	for _, c := range *dp {
		is = append(is, int(c))
	}
	sort.Ints(is)
	var cs []Card
	for _, i := range is {
		cs = append(cs, Card(i))
	}
	*dp = Deck(cs)
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

var faces = []string{"J", "Q", "K", "A"}

// String returns a string representation of the card, e.g., "4♡", "A♠"
//
// With the default value of Ranks, the rank of the card will be returned starting at 2 until 10, then J, Q, K, A.
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
		s = faces[v-9]
	}
	return fmt.Sprintf("%s%s", s, c.Suit())
}

// Suit represents the Suit of the card, one of Suits
type Suit int

// String returns the string representation of the card, e.g., "Hearts"
func (s Suit) String() string {
	return string(Suits[s])
}
