package game

func (g *Game) ResumeGame() (*Game, error) {

	g.DbRead()

	return g, nil

}
