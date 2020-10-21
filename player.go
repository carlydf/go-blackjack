package blackjack

import (
	"fmt"
	"github.com/gophercises/deck"
	"strconv"
)

type Player struct {
	Name string
	Hand []deck.Card
	Dealer bool
}

func (p Player) String() string {
	s := fmt.Sprintf("%v:\n", p.Name)
	for i, c := range p.Hand {
		if p.Dealer && i == 0 {
			s += "   hidden card\n"
		} else {
			s += fmt.Sprintf("   %v\n", c)
		}
	}
	s += "\n"
	return s
}

func (p Player) PrintHand() {
	for i, c := range p.Hand {
		if p.Dealer && i == 0 {
			fmt.Print("   hidden card\n")
		} else {
			fmt.Printf("   %v\n", c)
		}
	}
}

func (p *Player) Draw(c deck.Card) {
	p.Hand = append(p.Hand, c)
}

func (p Player) ScoreHand() (s1, s11 int) {
	s1, s11 = 0, 0
	for _, c := range p.Hand {
		t1, t11 := scoreCard(c)
		s1 += t1
		s11 += t11
	}
	return s1, s11
}

func scoreCard(c deck.Card) (s1, s11 int) {
	if c.Rank == "A" {
		s1 = 1
		s11 = 11
	} else if c.Rank == "J" || c.Rank == "Q" || c.Rank == "K" {
		s1, s11 = 10, 10
	} else {
		s1, _ = strconv.Atoi(c.Rank) // cards 2-10
		s11 = s1
	}
	return s1, s11
}
