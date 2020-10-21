package blackjack

import (
	"fmt"
	"github.com/gophercises/deck"
)

// i read that global vars are problematic, but idk why
// could replace this with a struct, but that seemed unnecessary
var players []*Player // dealing order matters
var dk deck.Deck
var dealToNext int

func StartGame() {
	fmt.Print("Starting game...\n")
	names := []string{"me"}
	for _, n := range names {
		players = append(players, &Player{Name: n, Dealer: false})
	}
	// dealer must be last
	players = append(players, &Player{Name: "dealer", Dealer: true})
	dk = deck.New(deck.WithShuffle())
	dealToNext = 0
	deal(2)
}

func deal(numCards int) {
	fmt.Printf("Dealing %v cards to all players\n\n", numCards)
	for i := 0; i < numCards; i++ {
		for j, _ := range players {
			p := players[(dealToNext + j) % len(players)]
			p.Draw(draw())
		}
	}
	dealToNext = (dealToNext + 1) % len(players)
}

func draw() deck.Card {
	c := dk[0]
	dk = dk[1:]
	return c
}

func PrintFullStatus() {
	fmt.Println("### GAME STATUS (FULL/DEBUG) ###")
	fmt.Printf("Cards in deck: %v\n", len(dk))
	fmt.Printf("Deal to next: %v(%v)\n", players[dealToNext].Name, dealToNext)
	fmt.Println("List of players:")
	for _, p := range players {
		fmt.Print(p)
	}
}

func PrintStatus() {
	fmt.Println("### GAME STATUS ###")
	for _, p := range players {
			fmt.Print(p)
	}
}
