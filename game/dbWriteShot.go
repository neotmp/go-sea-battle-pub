package game

import (
	"fmt"

	"github.com/neotmp/go-sea-battle/database"
)

func (g *Game) DbWriteShot(player, damage, cell int) {

	fmt.Println(g.Id, player, damage, cell, "DB CONNECT")

	database.DB.Exec("INSERT INTO shots (game_id, player_id, damage, cell_index) VALUES($1, $2, $3, $4)", g.Id, player, damage, cell)

	//execute
}
