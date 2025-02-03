package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/neotmp/go-sea-battle/database"
	"github.com/neotmp/go-sea-battle/routes"
)

func healthStatus(g *[100]uint8) {
	var carrier uint8
	var battleship uint8
	var submarine uint8
	var cruiser uint8
	var destroyer uint8

	for _, v := range *g {
		if v == uint8(5) {
			carrier += 1
		}
		if v == uint8(4) {
			battleship += 1
		}

		if v == uint8(3) {
			submarine += 1
		}

	}

	fmt.Printf("Carrier health is %d of 5\n", carrier)
	fmt.Printf("Battleship health is %d of 4\n", battleship)
	fmt.Printf("Submarine health is %d of 3\n", submarine)
	fmt.Printf("Cruiser health is %d of 3\n", cruiser)
	fmt.Printf("Dstroyer health is %d of 2\n", destroyer)

}

// 2 grids
// my grid
// enemy's grid

// based on the grid I calculate damage to the ships and show to players
// history of the hits
// whose turn to fire a shot?

// grid is stored as a 100-lenght string in the table
// 0 - empty cell
// 9 - hit
// 8 - miss
// carrier - 5
// battleship - 4
// cruiser - 3
// submarine - 2
// destroyer - 1

// Human deploys their fleet
// once all ships are in place we send a fetch request with human's grid
// new game is initiated, data is stored
// computer deployes its assets
// status of teh game changes to active and whosturn becomes 1, human always commences the war
// human pick cell and hit fire button
// we send request
// we record the shot as a miss or a hit
// we change whosTurn to that of computer's
// computer fires its shot
// we change the turn
// we return two updated grids
// on client side we debounce showing the response to simulate a little bit of delay lest it be too quick for human palyer
// to enjoy the echange
// we repeat all previous steps til one of teh sides looses their fleet
// when grid returns no  surviving ships left we declare the winner and change status of teh game to inactive

// var num string = "0000000000000000000009"

// grid := new([100]uint8)
// grid[0] = 5
// grid[1] = 5
// grid[2] = 5
// grid[3] = 5
// grid[4] = 99
// grid[34] = 4
// grid[44] = 4
// grid[54] = 4
// grid[64] = 4
// grid[15] = 3
// grid[16] = 99
// grid[17] = 3

// for _, v := range num {
// 	num, err := convRune(v)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(num)
// }

// fmt.Println(num)

// healthStatus(grid)

func main() {

	//gOne := new([100]int32)
	//gTwo := new([100]int32)

	//grid_h := gOne
	//grid_c := gTwo

	// g := game.Game{
	// 	Id:        7,
	// 	GridHuman: gOne[0:],
	// 	GridComp:  gTwo[0:],
	// 	WhoseTurn: 1,
	// 	WonBy:     2,
	// 	StartedAt: time.Now(),
	// }

	// gg := game.Game{Id: 7}

	// read, err := gg.DbRead()
	// if err != nil {
	// 	fmt.Println(err, "Error while readin/writing DB")
	// }

	//masked, _ := read.MaskShips()

	//fmt.Println(masked.GridComp, "Comp Grid")

	// engage, err := read.Engage(1, 2)
	// if err != nil {
	// 	fmt.Println(err, "Error while readin/writing DB")
	// }

	// deploy, err := read.DeployShipsLogic()
	// if err != nil {
	// 	fmt.Println(err, "Error while readin/writing DB")
	// }

	// fmt.Println(deploy, "deploy")

	// gg := game.Game{}

	// newG, err := gg.StartGame()
	// newG.Id = 1
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(newG.WhoseTurn, "New Game TURN")

	//p, d, c := gg.CalculateDamage(0, 0)

	//fmt.Println(p, d, c, "Player, Damage, Cell")

	// play, err := newG.PlayGame()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(play.WonBy, "Winner+++")

	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     os.Getenv("ALLOW_ORIGINS"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	routes.Setup(app)

	if os.Getenv("ENV") == "dev" {
		app.Listen(":4003")
	} else {
		log.Fatal(app.ListenTLS(":4003", "./cert.pem", "./cert.key"))
	}

}
