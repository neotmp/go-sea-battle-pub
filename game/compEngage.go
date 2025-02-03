package game

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

// FireAtWill returns all slice indeces that are not Hit && Miss, i.e. whose value is not 8 OR 9
// From this array we randomly pick a number to fire a shot at
func (g *Game) FireAtWill() ([]int32, error) {
	// read, err := g.DbRead()

	// if err != nil {
	// 	return []int32{}, err
	// }

	var shotsAvailable []int32

	// We look at GridHuman to see what Cells are available for shots, any other
	// those w/ values of either 8 or 9
	for i, v := range g.GridHuman {

		if v != 8 && v != 9 {
			//fmt.Println(v, "Value of Cell", i)
			shotsAvailable = append(shotsAvailable, int32(i))

		}

	}

	return shotsAvailable, nil

}

func randomInteger(s []int32) (int32, error) {

	// s, err := g.FireAtWill()
	// if err != nil {
	// 	return 0, err
	// }

	// TO DO Make it look better
	rand.Seed(uint64(time.Now().UnixNano() * 99))
	// TO DO MAke sure there's a check on lenght of the slice, else it will PANIC
	randomIndex := rand.Intn(len(s))
	pick := s[randomIndex]

	return pick, nil

}

// Engage accepts two params: player id and coordinate, both integers
// Coordinate, or Cell, is an index of the grid, i.e. slice
func (g *Game) ComputerEngage() (*Game, error) {

	shots, err := g.FireAtWill()
	if err != nil {
		return g, err
	}

	aim, err := randomInteger(shots)
	if err != nil {
		fmt.Println(err)
	}

	if g.WhoseTurn == 0 { //computer turn

		if g.GridHuman[aim] != 0 {
			p, c, d := g.CalculateDamage(int(aim))
			g.GridHuman[aim] = int32(d)
			g.DbWriteShot(p, d, c)
			g.WhoseTurn = 1
			//g.DbUpdate()
			//fmt.Println(g.GridHuman[aim], "Got cell-1")
			//fmt.Println("Direct Hit")
		} else {
			p, c, d := g.CalculateDamage(int(aim))
			g.GridHuman[aim] = int32(d)
			g.DbWriteShot(p, d, c)
			g.WhoseTurn = 1
			//g.DbUpdate()
			//fmt.Println(g.GridHuman[aim], "Got Cell-0")
			//fmt.Println("A Miss")

		}
	}

	return g, nil
}
