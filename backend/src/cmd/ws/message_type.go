package ws

/*
const (

	MATCH_POT_AMOUNT  MessageType = "match.pot-amount"
	MATCH_SEAT_TURN   MessageType = "match.seat-turn"
	MATCH_SEATS       MessageType = "match.seats"
	MATCH_TABLE_CARDS MessageType = "match.table-cards"
	MATCH_WINNERS     MessageType = "match.winners"

	OPPONENTS_ACTION MessageType = "opponents.action"
	OPPONENTS_INFO   MessageType = "opponents.info"

	USER_ACTION MessageType = "user.action"
	USER_EMOTE  MessageType = "user.emote"
	USER_INFO   MessageType = "user.info"
	USER_LOGIN  MessageType = "user.login"

)
*/
type MessageType string

const (
	MATCH_POT_AMOUNT  MessageType = "match.pot-amount"
	MATCH_SEAT_TURN   MessageType = "match.seat-turn"
	MATCH_SEATS       MessageType = "match.seats"
	MATCH_TABLE_CARDS MessageType = "match.table-cards"
	MATCH_WINNERS     MessageType = "match.winners"

	OPPONENTS_ACTION MessageType = "opponents.action"
	OPPONENTS_INFO   MessageType = "opponents.info"

	USER_ACTION MessageType = "user.action"
	USER_EMOTE  MessageType = "user.emote"
	USER_INFO   MessageType = "user.info"
	USER_LOGIN  MessageType = "user.login"
)
