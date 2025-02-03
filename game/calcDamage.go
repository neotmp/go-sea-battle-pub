package game

import "fmt"

// CalculateDamage returns player id of the player who fired the shot, cell and damage (8 or 9)
func (g *Game) CalculateDamage(cell int) (int, int, int) {

	if g.WhoseTurn == 1 { // human shot - 1
		if g.GridComp[cell] != 0 {
			fmt.Println(cell, "CELL 91")
			return 1, cell, 9
		} else {
			fmt.Println(cell, "CELL 81")
			return 1, cell, 8
		}

	} else { // computer shot
		if g.GridHuman[cell] != 0 {
			fmt.Println(cell, "CELL 90")
			return 0, cell, 9
		} else {
			fmt.Println(cell, "CELL 80")
			return 0, cell, 8
		}

	}

}
