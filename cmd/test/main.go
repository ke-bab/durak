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

	err = game.DoAction(&JoinAction{PlayerId: 1})
	if err != nil {
		log.Fatal(err)
	}
	err = game.DoAction(&JoinAction{PlayerId: 2})
	if err != nil {
		log.Fatal(err)
	}

	err = game.DoAction(&ReadyAction{PlayerId: 1})
	if err != nil {
		log.Fatal(err)
	}
	err = game.DoAction(&ReadyAction{PlayerId: 2})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", game)
}
