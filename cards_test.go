package cards

import (
	"reflect"
	"testing"
)

func TestShuffleSort(t *testing.T) {
	d := NewDeck()
	if !reflect.DeepEqual(d, NewDeck()) {
		t.Errorf("expected new deck")
	}
	if d[0] != Card(0) {
		t.Errorf("expected first card to be %q, got %q", Card(0), d[0])
	}
	d.Shuffle()
	// Technically possible though impossibly unlikely that the shuffled deck will be the sorted deck...
	if reflect.DeepEqual(d, NewDeck()) {
		t.Errorf("expected shuffled, got %v", d)
	}
	d.Sort()
	if !reflect.DeepEqual(d, NewDeck()) {
		t.Errorf("expected sorted deck, got %v", d)
	}
}

func TestCut(t *testing.T) {
	d := NewDeck()
	d1, d2 := d.Cut()
	if len(d1) != len(d2) || len(d1)+len(d2) != len(d) {
		t.Errorf("unexpected cut, got %v\n%v", d1, d2)
	}
	d1.Add(d2.TopN(len(d2))...)
	d1.Sort()
	if !reflect.DeepEqual(d1, d) {
		t.Errorf("unexpected result rejoining cut deck,\n got %v\nwant %v", d1, d)
	}
}

func TestEmpty(t *testing.T) {
	d := Deck(nil)
	if !d.Empty() {
		t.Errorf("expected empty deck")
	}
	if got := d.Top(); got != nil {
		t.Errorf("expected nil top card from empty deck")
	}

	d = NewDeck()
	d.TopN(len(d))
	if d != nil {
		t.Errorf("expected nil deck")
	}
}

func TestRiffle(t *testing.T) {
	d1 := Deck([]Card{0, 1, 2})
	d2 := Deck([]Card{3, 4, 5})
	d := Riffle(d1, d2)
	exp := Deck([]Card{0, 3, 1, 4, 2, 5})
	if !reflect.DeepEqual(d, exp) {
		t.Errorf("unexpected riffle\n got %s\nwant %s", d, exp)
	}
}
