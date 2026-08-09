package main

import (
	"backend/src/cmd/ws"
)

func main() {
	ws.InitServer("/", 3000)
}
