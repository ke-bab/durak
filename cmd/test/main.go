package main

import (
	. "durak"
	"fmt"
	"log"
)

func main() {
	manager, err := NewGameManager()
	if err != nil {
		log.Fatal(err)
	}
	game, err := manager.CreateGame()
	if err != nil {
		log.Fatal(err)
	}
	_, err = game.JoinPlayer()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", game)
}
