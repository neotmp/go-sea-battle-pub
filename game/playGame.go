package game

// PlayGame realizes the main game logic
// it receives a CELL coor (shot),
func (g *Game) PlayGame(shot int) (*Game, error) {

	// we start with the human shot
	g, err := g.HumanEngage(shot)
	if err != nil {
		return g, err
	}

	// then we assess damage caused by the human player
	g, err = g.DamageReport()
	if err != nil {
		return g, err
	}

	// then we see if have a winner
	g, err = g.WhoWins()
	if err != nil {
		return g, err
	}

	if g.WonBy == 1 {
		// we update db
		g, err := g.DbUpdate()
		if err != nil {
			return g, err
		}
		return g, nil
	}

	// if human has not sunk all computer's ships computer has a chance
	// to respond w/ its shot

	g, err = g.ComputerEngage()
	if err != nil {
		return g, err
	}

	// then we assess damage caused by the computer player
	g, err = g.DamageReport()
	if err != nil {
		return g, err
	}

	// then we see if we have a winner
	g, err = g.WhoWins()
	if err != nil {
		return g, err
	}

	if g.WonBy == 0 {
		// we update db
		g, err := g.DbUpdate()
		if err != nil {
			return g, err
		}
		return g, nil
	}

	g, err = g.DbUpdate()
	// if err != nil {
	// 	return w, err
	// }

	//add masked
	//masked := w.MaskShips()

	return g, nil
}
