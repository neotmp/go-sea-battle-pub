package game

// Because we use only two grids we shall not reveal the unmasked grid of the Computer player
// to Human, thus, before we send out json we mask the grid
// All values except for 9 - hit and 8 - miss would be changed to zeros
// this way Human player will only be seeing their hits and misses
func (g *Game) MaskShips() *Game {

	for i, v := range g.GridComp {
		if v != 0 && v != 8 && v != 9 {
			g.GridComp[i] = 0
		}
	}

	return g

}
