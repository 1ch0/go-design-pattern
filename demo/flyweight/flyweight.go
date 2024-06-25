package main

import "fmt"

var portCards = map[int]*Card{
	1: {
		Name:  "A",
		Value: "红",
	},
	2: {
		Name:  "A",
		Value: "黑",
	},
	3: {
		Name:  "B",
		Value: "红",
	},
}

type Card struct {
	Name  string
	Value string
}

type CardGame struct {
	Cards map[int]*Card
}

func NewGame(game map[int]*Card) *CardGame {
	ins := &CardGame{Cards: map[int]*Card{}}
	for i := range game {
		ins.Cards[i] = game[i]
	}
	return ins
}

func main() {
	game1 := NewGame(portCards)
	game2 := NewGame(portCards)

	fmt.Printf("game1:%#v\ngame2:%#v", game1.Cards[1], game2.Cards[1])

}
