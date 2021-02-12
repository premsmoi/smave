package main

import (
	"fmt"

	"github.com/premsmoi/smave/pkg/types"
)

func main() {
	cardList := types.CardList{}
	cardList.GenerateAllCards()
	for _, card := range cardList.Cards {
		fmt.Println(card)
	}
}
