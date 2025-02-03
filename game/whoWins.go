package game

import "fmt"

func (g *Game) WhoWins() (*Game, error) {

	//if g.WhoseTurn == 0 { // it means that human just fired a shot we need to check
	// if COMPUTER still has ships afloat

	if len(g.LostComp) == 5 {
		g.WonBy = 1
	}

	if len(g.LostHuman) == 5 {
		g.WonBy = 0
	}

	fmt.Println(g.WonBy, "Winner-----")
	fmt.Println(g.WhoseTurn, "Turn")

	return g, nil

}
