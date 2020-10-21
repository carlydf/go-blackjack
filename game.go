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
var whoseTurn int
var roundNum int

func StartGame() {
	fmt.Print("Starting game...\n")
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)
	names := []string{name} // change for multiplayer
	for _, n := range names {
		players = append(players, &Player{Name: n, Dealer: false})
	}
	// dealer must be last
	players = append(players, &Player{Name: "Dealer", Dealer: true})
	dk = deck.New(deck.WithShuffle())
	dk = deck.New(deck.WithShuffle()) // added because shuffle wasnt working?
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

func TakeTurn() {
	p := players[whoseTurn]
	fmt.Printf("%v's turn!\n", p.Name)
	if p.Dealer {
		DealerTurn(p)
		return
	}
	fmt.Print("   Enter H to hit or S to stand: ")
	var choice string
	fmt.Scanln(&choice)
	fmt.Printf("   %v chose " + choice + ".\n", p.Name)
	switch choice {
	case "S":
		stand()
		return
	case "H":
		busted := hit(p)
		if busted {
			fmt.Print("Game over! The winner is...") // change this for multiplayer
			fmt.Print(players[len(players) - 1]) // dealer is always last
			return
		}
		TakeTurn()
	}
	whoseTurn = (whoseTurn + 1) % len(players)
	return
}

func DealerTurn(p *Player) {
	// In our second iteration the dealer will play with typical dealer rules -
	// if they have a score of 16 or less, or a soft 17, they will hit.
	s1, s11 := p.ScoreHand()
	if (s11 == 17 && s1 < 17) || s1 <= 16 || s11 <= 16 { // i think s11 <= 16 is redundant
		fmt.Printf("   %v chose H.\n", p.Name)
		busted := hit(p)
		if busted {
			fmt.Print("Game over! The dealer lost.\n")
			return
		}
	} else {
		fmt.Printf("   %v chose S.\n", p.Name)
		stand()
		roundNum++
	}
}

func hit(p *Player) bool {
	p.Draw(draw())
	fmt.Printf("   %v's hand is now:\n", p.Name)
	fmt.Print(p)
	s1, _ := p.ScoreHand()
	if s1 > 21 {
		fmt.Printf("%v busted with score %v :(\n", p.Name, s1)
		return true
	}
	return false
}

func stand() {
	whoseTurn = (whoseTurn + 1) % len(players)
	fmt.Print("   Turn over.\n\n")
}

func CheckScores() (noBusts bool, winner *Player) {
	noBusts, winner = true, nil
	for _, p := range players {
		s1, s11 := p.ScoreHand()
		// check for blackjack
		if roundNum == 0 && s11 == 21 {
			winner = p
		}
		if s1 > 21 {
			noBusts = false
		}
	}
	return noBusts, winner
}

func PrintStatus() {
	fmt.Print("### GAME STATUS ###\n")
	for _, p := range players {
			fmt.Print(p)
	}
	fmt.Print("##################\n\n")
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
