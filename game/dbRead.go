package game

import (
	"github.com/lib/pq"
	"github.com/neotmp/go-sea-battle/database"
)

func (g *Game) DbRead() (*Game, error) {

	qr := `SELECT id, grid_h, grid_c, whose_turn, won_by, started_at, lost_h, lost_c, deployed, modified_at FROM games WHERE id = $1`

	if err := database.DB.QueryRow(qr,
		g.Id,
	).Scan(
		&g.Id,
		pq.Array(&g.GridHuman),
		pq.Array(&g.GridComp),
		&g.WhoseTurn,
		&g.WonBy,
		&g.StartedAt,
		pq.Array(&g.LostHuman),
		pq.Array(&g.LostComp),
		&g.Deployed,
		&g.ModifiedAt,
	); err != nil {
		return g, err
	}

	return g, nil
}
