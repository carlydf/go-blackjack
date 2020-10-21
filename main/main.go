package main

import (
	"fmt"
	"github.com/gophercises/blackjack"
)

func main() {
	blackjack.StartGame()
	//blackjack.PrintFullStatus()
	blackjack.PrintStatus()
	noBusts, winner := blackjack.CheckScores()
	for noBusts {
		if winner != nil {
			fmt.Print("Game over! The winner is...\n")
			fmt.Print(winner)
			return
		}
		blackjack.TakeTurn()
		noBusts, winner = blackjack.CheckScores()
	}
}
