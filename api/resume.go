package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-sea-battle/game"
)

func Resume(c *fiber.Ctx) error {

	id := c.Cookies("gameId")

	ni, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	g := new(game.Game)
	g.Id = uint32(ni)
	g.ResumeGame()
	g.MaskShips()

	return c.JSON(ServerResponse{Code: 2, Message: "The game is resumed. Your turn to fire.", GameId: id, Data: *g})
}
