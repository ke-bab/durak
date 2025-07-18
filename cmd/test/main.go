package main

import (
	. "durak"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	manager, err := NewGameManager()
	fatal(err)
	game, err := manager.CreateGame()
	fatal(err)
	pool, err := NewIdPool(100)
	fatal(err)
	p1, err := NewPlayer(pool)
	fatal(err)
	p2, err := NewPlayer(pool)
	fatal(err)

	err = game.DoAction(&JoinAction{Player: p1})
	fatal(err)
	err = game.DoAction(&JoinAction{Player: p2})
	fatal(err)

	err = game.DoAction(&ReadyAction{Player: p1})
	fatal(err)
	err = game.DoAction(&ReadyAction{Player: p2})
	fatal(err)
	err = game.DoAction(&PlayCardAction{
		Player: p1,
		Card:   &Card{Suit: Hearts, Rank: Eight},
	})
	fatal(err)

	printGame(game)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printGame(g *Game) {
	str, err := json.MarshalIndent(g, "", "  ")
	fatal(err)
	fmt.Println(string(str))
}
