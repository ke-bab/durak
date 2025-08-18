package main

import (
	. "durak"
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	gm, err := NewGameManager()
	fatal(err)
	game, err := gm.CreateGame()
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
