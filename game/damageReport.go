package game

import "fmt"

// addLostShips create set of unique values to ensure that
// sunk ships are added only once to the array when we check them in Damage Report
func (g *Game) addLostShip(num int) {

	var found bool

	if g.WhoseTurn == 1 {

		// empty array
		if len(g.LostComp) == 0 {
			g.LostComp = append(g.LostComp, int32(num))
			return
		} else {
			for _, v := range g.LostComp {

				if v == int32(num) {

					found = true
					return
				}
			}

			if !found {
				g.LostComp = append(g.LostComp, int32(num))
			}
		}

	} else {

		// empty array
		if len(g.LostComp) == 0 {
			g.LostComp = append(g.LostComp, int32(num))
			return
		} else {
			for _, v := range g.LostComp {

				if v == int32(num) {

					found = true
					return
				}
			}

			if !found {
				g.LostComp = append(g.LostComp, int32(num))
			}
		}

	}
}

// DamageReport loops over GridHuman or GridComp
// and add sunk ships to lost slice
// Since we take turns to takek shots we deliver DamageR Report
// only for one player, the other player cannot lose any ships during this cycle
func (g *Game) DamageReport() (*Game, error) {

	var carrier uint8    // 6 -num, l - 5
	var battleship uint8 // 5, -4
	var submarine uint8  // 4, -3
	var cruiser uint8    // 3, -3
	var destroyer uint8  //2, -2

	if g.WhoseTurn == 0 { // computer fires

		for _, v := range g.GridHuman {

			if v == 6 {
				carrier += 1
			}

			if v == 5 {
				battleship += 1
			}

			if v == 4 {
				submarine += 1
			}

			if v == 3 {
				cruiser += 1
			}

			if v == 2 {
				destroyer += 1
			}

		}

		if carrier == 0 {
			g.addLostShip(6)
		}
		if battleship == 0 {
			g.addLostShip(5)
		}
		if submarine == 0 {
			g.addLostShip(4)
		}
		if cruiser == 0 {
			g.addLostShip(3)
		}
		if destroyer == 0 {
			g.addLostShip(2)
		}

	} else {

		for _, v := range g.GridComp {
			if v == 6 {
				carrier += 1
			}

			if v == 5 {
				battleship += 1
			}

			if v == 4 {
				submarine += 1
			}

			if v == 3 {
				cruiser += 1
			}

			if v == 2 {
				destroyer += 1
			}

		}

		if carrier == 0 {
			g.addLostShip(6)
		}
		if battleship == 0 {
			g.addLostShip(5)
		}
		if submarine == 0 {
			g.addLostShip(4)
		}
		if cruiser == 0 {
			g.addLostShip(3)
		}
		if destroyer == 0 {
			g.addLostShip(2)
		}

	}

	fmt.Printf("Player %d Hits Carrier: %d, Battleship: %d, Submarine: %d, Cruiser: %d, Destroyer: %d", g.WhoseTurn, carrier, battleship, submarine, cruiser, destroyer)

	fmt.Println(g.LostComp, g.LostHuman, "LOSES COMP HUMAN")

	return g, nil

}
