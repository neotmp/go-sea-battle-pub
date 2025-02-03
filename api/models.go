package api

import "github.com/neotmp/go-sea-battle/game"

type ServerResponse struct {
	Code    uint8     `json:"code" db:"code"` // 0 - no game found, 1 - no game in progress, 2 game is available
	Message string    `json:"message" db:"message"`
	GameId  string    `json:"gameId" db:"game_id"`
	Data    game.Game `json:"data" db:",omitempty"`
}
