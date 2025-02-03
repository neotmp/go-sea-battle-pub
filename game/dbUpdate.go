package game

import (
	"github.com/lib/pq"
	"github.com/neotmp/go-sea-battle/database"
)

func (g *Game) DbUpdate() (*Game, error) {

	q := `UPDATE games SET
	whose_turn = $1,
	won_by = $2,
	finished_at = $3,
	grid_h = $4,
	grid_c = $5,
	lost_h = $6,
	lost_c = $7,
	deployed = $8,
	modified_at = now()
	WHERE id = $9 RETURNING id`

	if err := database.DB.QueryRow(q,
		&g.WhoseTurn,
		&g.WonBy,
		&g.FinishedAt,
		pq.Array(g.GridHuman),
		pq.Array(g.GridComp),
		pq.Array(g.LostHuman),
		pq.Array(g.LostComp),
		&g.Deployed,
		&g.Id,
	).Scan(&g.Id); err != nil {
		return g, err
	}

	return g, nil
}
