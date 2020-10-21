package main

import (
	"github.com/gophercises/blackjack"
)

func main() {
	blackjack.StartGame()
	//blackjack.PrintFullStatus()
	blackjack.PrintStatus()
	blackjack.TakeTurn()
}
