package game

import "time"

type Game struct {
	Id         uint32    `json:"id" db:"id"`
	GridHuman  []int32   `json:"gridHuman" db:"grid_h"`     // stores human ships on the grid
	GridComp   []int32   `json:"gridComp" db:"grid_c"`      // stores computer ships
	WhoseTurn  uint8     `json:"whoseTurn" db:"whose_turn"` // 0 - computer, 1 - human
	WonBy      uint8     `json:"wonBy" db:"won_by"`         // 0 - computer, 1 - human, 2 - in progress
	StartedAt  time.Time `json:"startedAt" db:"started_at"`
	FinishedAt string    `json:"finishedAt" db:"finished_at"`
	LostHuman  []int32   `json:"lostHuman" db:"lost_h"`  // human player losses
	LostComp   []int32   `json:"lostComp" db:"lost_c"`   // computer player losses
	Deployed   bool      `json:"deployed" db:"deployed"` // to lock Human players fleet in position to let 'im fire a first shot
	ModifiedAt string    `json:"modifiedAt" db:"modified_at"`

	//
	ComputerNationId uint16 `json:"computerNationId" db:"computer_nation_id"`
	HumanNationId    uint16 `json:"humanNationId" db:"human_nation_id"`
	NationHuman      string `json:"NationHuman" db:"nation_human"`
	NationComputer   string `json:"NationComp" db:"nation_computer"`
}

// Stored in Table for historical references and future training
type Shot struct {
	GameId    uint32    `json:"game_id"`
	PlayerId  uint16    `json:"player_id"` // 0 - computer, 1 - human
	Time      time.Time `json:"time"`
	Damage    uint8     `json:"damage"`     // 9 - hit, 8 - miss
	CellIndex uint8     `json:"cell_inxex"` // 0 - 99, since data is stored in [100]uint8
}

// To add different nations later, only contemporary fleet
type Ship struct {
	Id           uint16    `json:"id"`
	Type         string    `json:"type"`
	Lenght       uint8     `json:"lenght"` // cruiser, submarine -
	Number       uint8     `json:"number"` // submarine - 3, cruiser 2
	NationId     uint16    `json:"nation_id"`
	Nation       string    `json:"nation"`
	Name         string    `json:"name"`
	Commissined  time.Time `json:"commissioned"`
	Displacement uint32    `json:"displacement"`
	Class        string    `json:"class"`
}
