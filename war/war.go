package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ImJasonH/cards"
)

var (
	games = flag.Int64("n", 100, "Number of games to play")
	suits = flag.Int("suits", 4, "Number of suits to use")
	ranks = flag.Int("ranks", 13, "Number of ranks to use")
)

func main() {
	flag.Parse()

	p1wins := int64(0)
	var wg sync.WaitGroup
	for i := int64(0); i < *games; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if War().p1wins {
				atomic.AddInt64(&p1wins, int64(1))
			}
		}()
	}
	wg.Wait()
	fmt.Println(p1wins, *games-p1wins)
}

type Result struct {
	p1wins      bool
	d           time.Duration
	turns, wars int
}

func War() Result {
	start := time.Now()
	d := cards.NewDeck()
	d.Shuffle()
	p1, p2 := d.Cut()

	turn := 0
	wars := 0
	var pile Pile
	for !p1.Empty() && !p2.Empty() {
		if len(p1)+len(p2)+len(pile.Deck) != *suits**ranks {
			panic("whoops")
		}
		c1, c2 := *p1.Top(), *p2.Top()
		pile.Add(c1, c2)
		pile.Shuffle()
		switch {
		case c1 > c2:
			p1.Add(pile.All()...)
		case c2 > c1:
			p2.Add(pile.All()...)
		case c1 == c2:
			wars++
			pile.Add(p1.TopN(3)...)
			pile.Add(p2.TopN(3)...)
			pile.Shuffle()
		}
		turn++
	}
	return Result{
		p1wins: p2.Empty(),
		d:      time.Since(start),
		turns:  turn,
		wars:   wars,
	}
}

type Pile struct {
	cards.Deck
}

func (p *Pile) All() []cards.Card {
	return p.TopN(len(p.Deck))
}
