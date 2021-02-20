package main

import (
	"fmt"

	"github.com/premsmoi/smave/pkg/types"
)

func main() {
	boardA := types.Board{}
	boardA.AddPlayer(types.Player{Name: "Smoi", Role: types.People})
	boardA.AddPlayer(types.Player{Name: "Ake", Role: types.People})
	boardA.AddPlayer(types.Player{Name: "John", Role: types.People})
	boardA.AddPlayer(types.Player{Name: "Park", Role: types.People})

	deckA := boardA.Deck
	deckA.GenerateAllCards()
	deckA.Shuffle()

	for len(deckA) > 0 {
		i := len(deckA) % 4
		card := deckA.Draw()
		boardA.Players[i].CardList = append(boardA.Players[i].CardList, card)
		if (card == types.Card{Rank: "3", Suit: "Clubs"}) {
			boardA.PlayingIndex = i
		}
	}

	for i := 0; i < 4; i++ {
		fmt.Println(boardA.Players[i])
		if len(boardA.Players[i].CardList) != 13 {
			fmt.Println("Starting cards on player", i, "is invalid")
		}
	}

	// for boardA.IsOver == false {
	// 	for i := 0; i < len(boardA.Players); i++ {

	// 	}
	// }
}
