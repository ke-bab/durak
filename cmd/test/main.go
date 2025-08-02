package main

import (
	. "durak"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	game := NewGame()
	pool, err := NewIdPool(100)
	fatal(err)
	factory := NewPlayerFactory(pool)
	p1, err := factory.CreatePlayer()
	fatal(err)
	stateCanJoin, err := NewStateCanJoin(game)
	fatal(err)
	stateCanJoin.Join(p1)

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
