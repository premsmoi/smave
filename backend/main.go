package main

import (
	"fmt"
	"time"

	"github.com/premsmoi/smave/pkg/types"
)

func main() {
	boardA := types.Board{}
	boardA.GameCount = 1
	boardA.AddPlayer(types.Player{Name: "Smoi", Role: types.People})
	boardA.AddPlayer(types.Player{Name: "Ake", Role: types.People})
	boardA.AddPlayer(types.Player{Name: "John", Role: types.People})
	boardA.AddPlayer(types.Player{Name: "Park", Role: types.People})
	boardA.InitiateDeck()
	boardA.InitiateGame()
	boardA.AutoPlayFirstGame()

	player := &boardA.Players[boardA.PlayerIndex]

	for boardA.IsOver == false {
		if player.Pass || player.Complete {
			player = boardA.NextPlayer()
			continue
		}
		if boardA.PlayingIndex == boardA.PlayerIndex {
			fmt.Println("Next Round!!")
			boardA.NextRound()
		}

		playableCardIndexList := player.FindPlayableCards(boardA.PlayingCards)

		if len(playableCardIndexList) > 0 {
			playableCardList := player.Play(playableCardIndexList)
			boardA.PlayingCards = playableCardList
			boardA.PlayingIndex = boardA.PlayerIndex
			fmt.Println(player.Name, boardA.PlayingCards)
			boardA.HaveCompleted(player)
		} else {
			player.Pass = true
			fmt.Println(player.Name, "Pass!")
		}

		player = boardA.NextPlayer()
		 time.Sleep(10000000)
	}
}
