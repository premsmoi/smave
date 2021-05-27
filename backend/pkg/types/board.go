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
	Name             string
	Deck             CardList
	PlayingCards     CardList
	Pile             CardList
	Players          []Player
	NextRoundPlayers []Player
	PlayerIndex      int
	PlayingIndex     int
	GameCount        int
	RoundCount       int
	IsOver           bool
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
	if len(*deck) != 52 {
		fmt.Println("Cannot initiate game. Number of cards in deck is invalid.")
		return
	}
	board.NextRoundPlayers = []Player{}
	board.RoundCount = 1
	board.PlayerIndex = 0
	board.PlayingIndex = 0
	for _, player := range board.Players {
		player.Pass = false
		player.Complete = false
	}
	for len(*deck) > 0 {
		i := len(*deck) % 4
		card := (*deck).Draw()
		board.Players[i].AddCard(card)
		if (board.GameCount == 1 && card == Card{Rank: "3", Suit: "Clubs"}) {
			board.PlayerIndex = i
			board.PlayingIndex = i
		}
	}
}

// AutoPlayFirstGame is exported
func (board *Board) AutoPlayFirstGame() {
	board.RoundCount = 1
	fmt.Println("Round:", board.RoundCount)
	i, _ := board.Players[board.PlayerIndex].CardList.Find(Card{Rank: "3", Suit: "Clubs"})
	board.PlayingCards = board.Players[board.PlayerIndex].Play([]int{i})
	fmt.Println(board.Players[board.PlayerIndex].Name, board.PlayingCards)
	board.NextPlayer()
	// board.Players[board.PlayerIndex].FindPlayableCards(board.PlayingCards)
}

// NextPlayer is exported
func (board *Board) NextPlayer() *Player {
	board.PlayerIndex = (board.PlayerIndex + 1) % 4
	return &board.Players[board.PlayerIndex]
}

// NextRound is exported
func (board *Board) NextRound() int {
	board.RoundCount = board.RoundCount + 1
	fmt.Println("Round:", board.RoundCount)
	board.PlayingCards = CardList{}
	for index, _ := range board.Players {
		board.Players[index].Pass = false
	}
	return board.RoundCount
}

// HaveCompleted is exported
func (board *Board) HaveCompleted(player *Player) {
	if len(player.CardList) == 0 {
		board.NextRoundPlayers = append(board.NextRoundPlayers, *player)
		player.Complete = true
		fmt.Println(player.Name, "have completed in this game")
		if board.NextGame() {
			return
		}
		board.NextRound()
	}
}

// NextGame is exported
func (board *Board) NextGame() bool {
	completedCount := 0
	for _, player := range board.Players {
		if player.Complete {
			completedCount++
		}
	}
	if completedCount == 3 {
		for _, player := range board.Players {
			if !player.Complete {
				board.NextRoundPlayers = append(board.NextRoundPlayers, player)
			}
		}
		board.Players = board.NextRoundPlayers
		board.IsOver = true
		// board.InitiateDeck()
		// board.InitiateGame()
		// fmt.Println("Round:", board.RoundCount)
		return true
	}
	return false
}
