package game

// HumanEngage receives a cell number (shot) and determines
// if the shot must be registered as a hit or a miss
// returns pointer to the game
func (g *Game) HumanEngage(shot int) (*Game, error) {

	// Human returns their shot
	if g.WhoseTurn == 1 {

		if g.GridComp[shot] != 0 {
			p, c, d := g.CalculateDamage(int(shot))
			g.GridComp[shot] = int32(d)
			g.DbWriteShot(p, d, c)
			g.WhoseTurn = 0
			//g.DbUpdate()
			//fmt.Println(g.GridComp[shot], "Got cell-1")
			//fmt.Println("Direct Hit")
		} else {
			p, c, d := g.CalculateDamage(int(shot))
			g.GridComp[shot] = int32(d)
			g.DbWriteShot(p, d, c)
			g.WhoseTurn = 0
			//g.DbUpdate()
			//fmt.Println(g.GridComp[aim], "Got Cell-0")
			//fmt.Println("A Miss")

		}

	}

	return g, nil
}
