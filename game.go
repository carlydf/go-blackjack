package blackjack

import (
	"fmt"
	"github.com/gophercises/deck"
)

// i read that global vars are problematic, but idk why
// could replace this with a struct, but that seemed unnecessary
var players []Player // dealing order matters
var dck deck.Deck
var dealToNext int

func startGame() {
	names := []string{"me", "dealer"} // dealer must be last
	for _, n := range names {
		players = append(players, Player{Name: n})
	}
	dck = deck.New(deck.WithShuffle())
	dealToNext = 0
	deal(2)
}

func deal(numCards int) {
	fmt.Printf("Dealing %v cards to all players", numCards)
	for i := 0; i < numCards; i++ {
		for j, _ := range players {
			p := players[(dealToNext + j) % len(players)]
			p.Hand = append(p.Hand, draw())
		}
	}
	dealToNext = (dealToNext + 1) % len(players)
}

func draw() deck.Card {
	c := dck[0]
	dck = dck[1:]
	return c
}
