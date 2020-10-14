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
	s := fmt.Sprintf("%v: %v", p.Name, len(p.Hand))
	for _, c := range p.Hand {
		s += fmt.Sprintf("poop")
		s += fmt.Sprintf("%v, ", c)
	}
	s += "\n"
	return s
}
