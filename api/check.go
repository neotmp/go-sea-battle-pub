package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-sea-battle/database"
	"github.com/neotmp/go-sea-battle/game"
)

func Check(c *fiber.Ctx) error {

	i := c.Cookies("gameId")

	fmt.Println(i, "Got ID in cookie")

	//g := 1
	r := new(game.Game)

	if i == "" {
		return c.JSON(ServerResponse{Code: 0, Message: "Game Id not found", GameId: i})
	} else {
		if err := database.DB.QueryRow("SELECT id, won_by FROM games WHERE id = $1", i).Scan(&r.Id, &r.WonBy); err != nil {
			// put a logger here?
			return c.JSON(ServerResponse{Code: 0, Message: "Game Id not found", GameId: i})
		}

	}

	// if id found and status is not 2, i.e. game is not in progress

	if r.WonBy != 2 {
		return c.JSON(ServerResponse{Code: 1, Message: "Game is finished. You need to start new game.", GameId: i})
	}

	return c.JSON(ServerResponse{Code: 2, Message: "Game may be resumed.", GameId: i})

}
