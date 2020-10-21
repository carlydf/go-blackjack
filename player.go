package blackjack

import (
	"fmt"
	"github.com/gophercises/deck"
)

type Player struct {
	Name string
	Hand []deck.Card
}

func (p Player) String() string {
	s := fmt.Sprintf("%v:", p.Name)
	for _, c := range p.Hand {
		s += fmt.Sprintf("%v, ", c)
	}
	s += "\n"
	return s
}

func (p *Player) Draw(c deck.Card) {
	p.Hand = append(p.Hand, c)
}
