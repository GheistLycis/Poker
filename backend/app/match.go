package app

type Message[T any] struct {
	Type    string
	Payload T
}

type DirectMessage struct {
	player  string
	message *Message[any]
}

type Match struct {
	players      map[string]*Player
	register     chan *Player
	unregister   chan string
	broadcastMsg chan *Message[any]
	directMsg    chan *DirectMessage
	playerTurn   chan string
}

func (m *Match) handleNextTick() {
	for {
		select {
		case p := <-m.register:
			m.players[p.name] = p

		case pn := <-m.unregister:
			player := m.players[pn]

			delete(m.players, pn)
			close(player.send)

		case m := <-m.broadcastMsg:
			for _, p := range m.players {
				p.send <- m
			}

		case dm := <-m.directMsg:
			player := m.players[dm.player]

			player.send <- dm.message

		case pn := <-m.playerTurn:
			player := m.players[pn]
			msg := &Message[any]{}

			msg.Type = "m.seat-turn"
			msg.Payload = map[string]int{
				"seatIndex": player.seatIndex,
			}
			m.broadcastMsg <- msg
		}
	}
}
