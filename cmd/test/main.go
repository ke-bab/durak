package main

import (
	. "durak"
	"fmt"
	"log"
)

func main() {
	g := NewGame()
	err := g.JoinPlayer(Player{ID: 123})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", g)
}
