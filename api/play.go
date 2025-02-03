package api

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-sea-battle/game"
)

func Play(c *fiber.Ctx) error {

	id := c.Cookies("game_id")

	fmt.Println(id, "Got ID in cookie")

	type Shot struct {
		Cell int16 `json:"cell"`
	}

	s := new(Shot)

	if err := c.BodyParser(&s); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}

	fmt.Println(s.Cell, "Got Cell")

	// Shot Cannot be smaller than 0 and larger than 99
	if s.Cell < 0 || s.Cell > 99 {

		c.JSON(fiber.Map{"Error Message": "Your shot should be within the range: from 0 (inclusive) to 99 (inclusive)",
			"Response code": 406})

		return c.SendStatus(406)
		//errors.New("your shot should be within the range: from 0 (inclusive) to 99 (inclusive)")
	}

	// we got game id and shot

	g := new(game.Game)

	if id == "" {
		return errors.New("no game id given")
	}
	num, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	g.Id = uint32(num)

	read, err := g.DbRead()
	if err != nil {
		fmt.Println(err, "Err")

	}

	// Game logic is here
	play, err := read.PlayGame(int(s.Cell))
	if err != nil {
		fmt.Println(err, "Err")

	}

	// human takes their shot

	fmt.Println(play.LostComp, play.LostHuman, "FINAL LOST COMP HUMAN")

	return c.JSON(play)
}
