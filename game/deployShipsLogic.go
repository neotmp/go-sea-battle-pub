package game

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func generateRandomInt(limit int) int {

	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano() * 6599384432)))
	//r = rand.Seed(uint64(time.Now().UnixNano()))
	// TO DO MAke sure there's a check on lenght of the slice, else it will PANIC
	//randomIndex := rand.Intn(len(s))
	//pick := s[randomIndex]

	return int(r.Int31n(int32(limit)))
}

// Deploy places a ship into the slice if all cells it should occupy are empty
func (g *Game) placeShip(pos []int, number int) (*Game, bool, error) {

	//fmt.Println("GOT NUM", int32(len(pos)))

	if len(pos) == 0 {
		return g, false, errors.New("empty slice")
	}

	// We need to check all cells to make sure they're not occupied
	// and only then insert
	// h holds "cleared" indeces, lenght is the VALUE we insert
	var h []int
	for _, v := range pos {

		if g.GridComp[v] != 0 {
			return g, false, errors.New("cell is occupied")
		}

		h = append(h, v)

	}

	for _, vv := range h {
		g.GridComp[vv] = int32(number)

	}

	return g, true, nil

}

func (g *Game) checkCellsAvailability(lenght, number, cell, axis, direction int) ([]int, int) {

	var cells []int

	if axis == 0 { // we go up

		if direction == 1 { // we add 10

			var c = cell

			for i := 1; i <= lenght; i++ {
				c += 10
				if c > 99 {
					return []int{}, number
				}
				cells = append(cells, c)

			}

		} else {

			var c = cell

			for i := 1; i <= lenght; i++ {
				c -= 10
				if c < 0 {
					return []int{}, number
				}
				cells = append(cells, c)

			}

		}

	} else {

		if direction == 0 { // we add 1

			var c = cell

			for i := 1; i <= lenght; i++ {
				c += 1
				if c > 99 {
					return []int{}, number
				}
				cells = append(cells, c)

			}

		} else { // we subtract 1

			var c = cell

			for i := 1; i <= lenght; i++ {
				c -= 1
				if c < 0 {
					return []int{}, number
				}
				cells = append(cells, c)

			}

		}

	}

	return cells, number
}

func whereTo() (int, int, int) {

	var cell int
	var axis int      // 0 - Y, 1 - x
	var direction int // 0 - up/left, 1 - down/right

	cell = generateRandomInt(100)
	if cell%2 == 0 {
		if cell > 50 {
			axis = 1
			direction = 0
		} else {
			axis = 0
			direction = 1

		}
	}

	return cell, axis, direction

}

// 1. DeployShipsLogic limits possible positions for each ship being deployed
// 2. Randomly pick a Cell out of 99
// 3. Randomly pick the axis: X or Y
// 4. Randomly pick the direction: up or down for Y, and left or right for X
// 5. Randomly pick a ship from the [5]uint8{1,2,3,4,5}, where VALUE != 0
// 6. Get the indeces for a ship based on its lenght, if their all values are zeros,
// deploy the ship, set value of the ship in the array to ZERO
// 7. if step 6 returns NONE-ZERO values, repeat step 1 thru 6 till step 6 returns ZERO-lenght array
// 8. return pointer to Game struct
func (g *Game) DeployShipsLogic() (*Game, error) {

	ships := []Ship{
		{Type: "Carrier",
			Lenght: 5,
			Number: 6},

		{Type: "Battleship",
			Lenght: 4,
			Number: 5},

		{Type: "Cruiser",
			Lenght: 3,
			Number: 4},

		{Type: "Submarine",
			Lenght: 3,
			Number: 3},

		{Type: "Destroyer",
			Lenght: 2,
			Number: 2},
	}

	for _, v := range ships {

		var do bool = true

		for do {

			cell, axis, direction := whereTo()

			// WE need to check here
			_, ok, err := g.placeShip(g.checkCellsAvailability(int(v.Lenght), int(v.Number), int(cell), int(axis), int(direction)))
			if err != nil {
				fmt.Println(err)
			}

			if ok {
				do = false
			}

		}

	}

	return g, nil
}
