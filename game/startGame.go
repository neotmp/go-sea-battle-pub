package game

import (
	"fmt"
	"time"
)

func (g *Game) StartGame() (*Game, error) {

	// Initializing two arrays with zeros
	gOne := new([100]int32)
	gTwo := new([100]int32)

	g.WhoseTurn = 1 // Human always fires a first shot
	g.WonBy = 2     // 0 - computer, 1 - human, 2 - In Progress
	g.StartedAt = time.Now()
	g.GridHuman = gOne[0:] // we recieve this grid from Front-End
	g.GridComp = gTwo[0:]
	g.LostHuman = []int32{}
	g.LostComp = []int32{}

	// Computer Deploys its fleet

	deployed, err := g.DeployShipsLogic()
	if err != nil {
		return g, err
	}

	fmt.Println("Deployed", deployed)

	// After computer deployed its fleet we write into datebase

	write, err := deployed.DbWrite()
	if err != nil {
		return g, err
	}

	return write, nil

}
