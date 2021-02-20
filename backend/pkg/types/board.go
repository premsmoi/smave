package types

import (
	"fmt"
)

const (
	maxPlayer     = 4
	maxCardInDeck = 52
)

// Board is exported
type Board struct {
	Name         string
	Deck         CardList
	Playing      CardList
	Pile         CardList
	Players      []Player
	PlayingIndex int
	GameCount    int
	TurnCount    int
	IsOver       bool
}

// AddPlayer is exported
func (board *Board) AddPlayer(player Player) bool {
	if len(board.Players) == maxPlayer {
		fmt.Println("The player", player.Name, "cannot be added. Board", board.Name, "is full.")
		return false
	}
	board.Players = append(board.Players, player)
	return true
}

// InitiateDeck is exported
func (board *Board) InitiateDeck() {
	cardList := CardList{}
	cardList.GenerateAllCards()
	cardList.Shuffle()
	board.Deck = cardList
}

// InitiateGame is exported
func (board *Board) InitiateGame() {
	deck := &board.Deck
	if len(board.Deck) != 52 {
		fmt.Println("Cannot initiate game. Number of cards in deck is invalid.")
		return
	}
	for len(*deck) > 0 {
		i := len(*deck) % 4
		card := (*deck).Draw()
		board.Players[i].AddCard(card)
		if (card == Card{Rank: "3", Suit: "Clubs"}) {
			board.PlayingIndex = i
		}
	}
}
