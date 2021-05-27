package test

import (
	"testing"

	"github.com/premsmoi/smave/pkg/types"
)

func TestAddPlayer(t *testing.T) {
	boardA := types.Board{Name: "TestBoard"}
	actualLenth := len(boardA.Players)
	expectedLength := 0

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	boardA.AddPlayer(types.Player{Name: "TestPlayer", Role: types.People})
	actualLenth = len(boardA.Players)
	expectedLength = 1

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	boardA.AddPlayer(types.Player{Name: "TestPlayer", Role: types.People})
	actualLenth = len(boardA.Players)
	expectedLength = 2

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	boardA.AddPlayer(types.Player{Name: "TestPlayer", Role: types.People})
	actualLenth = len(boardA.Players)
	expectedLength = 3

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	boardA.AddPlayer(types.Player{Name: "TestPlayer", Role: types.People})
	actualLenth = len(boardA.Players)
	expectedLength = 4

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	boardA.AddPlayer(types.Player{Name: "TestPlayer", Role: types.People})
	actualLenth = len(boardA.Players)
	expectedLength = 4

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}
}

func TestInitiateDeck(t *testing.T) {
	board := types.Board{}
	actualLenth := len(board.Deck)
	expectedLength := 0

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	board.InitiateDeck()
	actualLenth = len(board.Deck)
	expectedLength = 52

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}
}

func TestInitiateGame(t *testing.T) {
	board := types.Board{}
	for i := 0; i < 4; i++ {
		board.AddPlayer(types.Player{Name: "TestPlyer", Role: types.People})
	}

	board.InitiateDeck()

	actualLenth := len(board.Deck)
	expectedLength := 52

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

	board.InitiateGame()

	expectedLength = 13

	for i := 0; i < 4; i++ {
		actualLenth := len(board.Players[i].CardList)
		if actualLenth != expectedLength {
			t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
		}
	}

	actualLenth = len(board.Deck)
	expectedLength = 0

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}

}

func TestNextPlayer(t *testing.T) {
	board := types.Board{}

	for i := 0; i < 4; i++ {
		board.AddPlayer(types.Player{Name: "TestPlyer", Role: types.People})
	}

	actualPlayerIndex := board.PlayerIndex
	expectedPlayerIndex := 0

	if actualPlayerIndex != expectedPlayerIndex {
		t.Error("Next player index is invalid. Expected:", expectedPlayerIndex, "Actual:", actualPlayerIndex)
	}

	board.NextPlayer()
	actualPlayerIndex = board.PlayerIndex
	expectedPlayerIndex = 1

	if actualPlayerIndex != expectedPlayerIndex {
		t.Error("Next player index is invalid. Expected:", expectedPlayerIndex, "Actual:", actualPlayerIndex)
	}

	board.NextPlayer()
	actualPlayerIndex = board.PlayerIndex
	expectedPlayerIndex = 2

	if actualPlayerIndex != expectedPlayerIndex {
		t.Error("Next player index is invalid. Expected:", expectedPlayerIndex, "Actual:", actualPlayerIndex)
	}

	board.NextPlayer()
	actualPlayerIndex = board.PlayerIndex
	expectedPlayerIndex = 3

	if actualPlayerIndex != expectedPlayerIndex {
		t.Error("Next player index is invalid. Expected:", expectedPlayerIndex, "Actual:", actualPlayerIndex)
	}

	board.NextPlayer()
	actualPlayerIndex = board.PlayerIndex
	expectedPlayerIndex = 0

	if actualPlayerIndex != expectedPlayerIndex {
		t.Error("Next player index is invalid. Expected:", expectedPlayerIndex, "Actual:", actualPlayerIndex)
	}
}
