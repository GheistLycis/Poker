package app

/*
const (

	CHECK PlayerAction = "CHECK"
	CALL  PlayerAction = "CALL"
	FOLD  PlayerAction = "FOLD"
	BET   PlayerAction = "BET"
	RAISE PlayerAction = "RAISE"

)
*/
type PlayerAction string

const (
	CHECK PlayerAction = "CHECK"
	CALL  PlayerAction = "CALL"
	FOLD  PlayerAction = "FOLD"
	BET   PlayerAction = "BET"
	RAISE PlayerAction = "RAISE"
)

var ActionsWithAmount = []PlayerAction{BET, RAISE}
