package blackjack

import (
	"fmt"
	"github.com/gophercises/deck"
)

type Player struct {
	Name string
	Hand []deck.Card
	Dealer bool
}

func (p Player) String() string {
	s := fmt.Sprintf("%v:\n", p.Name)
	for i, c := range p.Hand {
		if p.Dealer && i != 0 {
			s += "   hidden card\n"
		} else {
			s += fmt.Sprintf("   %v\n", c)
		}
	}
	s += "\n"
	return s
}

func (p *Player) Draw(c deck.Card) {
	p.Hand = append(p.Hand, c)
}

func (p Player) PrintDealer() {
	s := fmt.Sprintf("%v:\n", p.Name)
	for i, c := range p.Hand {
		if i == 0 {
			s += fmt.Sprintf("   %v\n", c)
		} else {
			s += "   hidden card\n"
		}
	}
	s += "\n"
	fmt.Print(s)
}
