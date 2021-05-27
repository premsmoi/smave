package test

import (
	"testing"

	"github.com/premsmoi/smave/pkg/types"
)

func TestPlay(t *testing.T) {
	player := types.Player{}
	player.CardList = []types.Card{
		{Rank: "A", Suit: "Clubs"},
		{Rank: "4", Suit: "Spades"},
		{Rank: "10", Suit: "Diamonds"},
	}
	returnedCard := player.Play([]int{0})[0]
	expectedCard := types.Card{Rank: "A", Suit: "Clubs"}
	actualLenth := len(player.CardList)
	expectedLength := 2

	if returnedCard != expectedCard {
		t.Error("Returned card is invalid. Expected:", expectedCard, "Actual:", returnedCard)
	}

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	returnedCard = player.Play([]int{1})[0]
	expectedCard = types.Card{Rank: "10", Suit: "Diamonds"}
	actualLenth = len(player.CardList)
	expectedLength = 1

	if returnedCard != expectedCard {
		t.Error("Returned card is invalid. Expected:", expectedCard, "Actual:", returnedCard)
	}

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	returnedCard = player.Play([]int{0})[0]
	expectedCard = types.Card{Rank: "4", Suit: "Spades"}
	actualLenth = len(player.CardList)
	expectedLength = 0

	if returnedCard != expectedCard {
		t.Error("Returned card is invalid. Expected:", expectedCard, "Actual:", returnedCard)
	}

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	player.Play([]int{0})
}

func TestAddCard(t *testing.T) {
	player := types.Player{}
	actualLenth := len(player.CardList)
	expectedLength := 0

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	player.AddCard(types.Card{Rank: "3", Suit: "Clubs"})
	actualLenth = len(player.CardList)
	expectedLength = 1

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	player.AddCard(types.Card{Rank: "Q", Suit: "Spades"})
	actualLenth = len(player.CardList)
	expectedLength = 2

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}
}

func TestFindPlayableCards(t *testing.T) {
	// player := types.Player{}
	// player.FindPlayableCards(2)
}
