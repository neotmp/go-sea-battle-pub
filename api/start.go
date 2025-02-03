package api

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-sea-battle/game"
)

// Start initiates a new game, deployes computers fleet, stores data in DB
// and returns JSON with
// This end-point is invoked everytime the user chooses
// to start a new game or there's no game to resume
func Start(c *fiber.Ctx) error {

	g := new(game.Game)

	// Initialize a New Game
	g, err := g.StartGame()
	if err != nil {
		return err
	}

	// convert uint into string to store in cookie cookie
	id := strconv.Itoa(int(g.Id))

	// sett a cookie w/ new id from DB
	c.Cookie(&fiber.Cookie{
		Name:     "gameId",
		Value:    id,
		Expires:  time.Now().Add(24 * time.Hour * 30), // 30 days
		HTTPOnly: true,
		Secure:   true,
		//Domain:   "http://localhost",
		SameSite: "lax",
	})

	g.MaskShips()

	return c.JSON(ServerResponse{Code: 2, Message: "Computer has deployed its fleet. You are advised to deploy yours.", GameId: id, Data: *g})

}
