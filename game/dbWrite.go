package game

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/neotmp/go-sea-battle/database"
)

func (g *Game) DbWrite() (*Game, error) {

	fmt.Println(g, "GOT HERE")

	qr := `INSERT INTO games (grid_h, grid_c, whose_turn, won_by, started_at, lost_h, lost_c, deployed, modified_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	if err := database.DB.QueryRow(qr,
		pq.Array(g.GridHuman),
		pq.Array(g.GridComp),
		g.WhoseTurn,
		g.WonBy,
		g.StartedAt,
		pq.Array(g.LostHuman),
		pq.Array(g.LostComp),
		g.Deployed,
		g.ModifiedAt,
	).Scan(&g.Id); err != nil {
		return g, err
	}

	return g, nil
}
