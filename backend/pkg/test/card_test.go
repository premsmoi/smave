package test

import (
	"testing"

	"github.com/premsmoi/smave/pkg/types"
)

// TestGenerateAllCards is exported
func TestGenerateAllCards(t *testing.T) {
	var generatedCardList = types.CardList{}
	generatedCardList.GenerateAllCards()
	expecetedLength := len(ExpectedCardList)
	actualLength := len(generatedCardList.Cards)

	if expecetedLength != actualLength {
		t.Error("Length is invalid. Expected:", expecetedLength, "Actual:", actualLength)
	}

	if generatedCardList.Type != types.NORMAL {
		t.Error("CardList type is invalid. Expected:", types.NORMAL, "Actual:", generatedCardList.Type)
	}

	for _, card := range ExpectedCardList {
		i, _ := generatedCardList.Find(card)
		if i == -1 {
			t.Error(card, "is missing.")
		}
	}
}

// TestFind1 is exported
func TestFind1(t *testing.T) {
	cardList := types.CardList{
		Cards: []types.Card{
			{Rank: "A", Suit: "Clubs"},
			{Rank: "5", Suit: "Spades"},
			{Rank: "6", Suit: "Diamonds"}},
		Type: types.TRIPLE}
	card := types.Card{Rank: "A", Suit: "Clubs"}
	i, res := cardList.Find(card)
	if i != 0 || res == false {
		t.Error("Find method returned invalid value.")
	}
}

// TestFind2 is exported
func TestFind2(t *testing.T) {
	cardList := types.CardList{
		Cards: []types.Card{
			{Rank: "A", Suit: "Clubs"},
			{Rank: "5", Suit: "Spades"},
			{Rank: "6", Suit: "Diamonds"}},
		Type: types.TRIPLE}
	card := types.Card{Rank: "A", Suit: "Spades"}
	i, res := cardList.Find(card)
	if i != -1 || res == true {
		t.Error("Find method returned invalid value.")
	}
}

// TestCanBeat1 is exported
func TestCanBeat1(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "A", Suit: "Spades"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "A", Suit: "Hearts"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat2 is exported
func TestCanBeat2(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "2", Suit: "Spades"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "2", Suit: "Diamonds"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat3 is exported
func TestCanBeat3(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "3", Suit: "Spades"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "3", Suit: "Clubs"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat4 is exported
func TestCanBeat4(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "4", Suit: "Hearts"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "4", Suit: "Diamonds"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat5 is exported
func TestCanBeat5(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "5", Suit: "Hearts"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "5", Suit: "Clubs"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat6 is exported
func TestCanBeat6(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "6", Suit: "Diamonds"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "6", Suit: "Clubs"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat7 is exported
func TestCanBeat7(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "2", Suit: "Diamonds"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "3", Suit: "Diamonds"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat8 is exported
func TestCanBeat8(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "2", Suit: "Diamonds"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "3", Suit: "Spades"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat9 is exported
func TestCanBeat9(t *testing.T) {
	cardListA := types.CardList{Cards: []types.Card{{Rank: "4", Suit: "Diamonds"}}, Type: types.SINGLE}
	cardListB := types.CardList{Cards: []types.Card{{Rank: "3", Suit: "Spades"}}, Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat10 is exported
func TestCanBeat10(t *testing.T) {
	cardListA := types.CardList{
		Cards: []types.Card{
			{Rank: "3", Suit: "Clubs"},
			{Rank: "3", Suit: "Spades"}},
		Type: types.COUPLE}

	cardListB := types.CardList{
		Cards: []types.Card{
			{Rank: "3", Suit: "Diamonds"},
			{Rank: "3", Suit: "Hearts"}},
		Type: types.COUPLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat11 is exported
func TestCanBeat11(t *testing.T) {
	cardListA := types.CardList{
		Cards: []types.Card{
			{Rank: "2", Suit: "Clubs"},
			{Rank: "2", Suit: "Diamonds"}},
		Type: types.COUPLE}

	cardListB := types.CardList{
		Cards: []types.Card{
			{Rank: "3", Suit: "Spades"},
			{Rank: "3", Suit: "Hearts"}},
		Type: types.COUPLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat12 is exported
func TestCanBeat12(t *testing.T) {
	cardListA := types.CardList{
		Cards: []types.Card{
			{Rank: "3", Suit: "Clubs"},
			{Rank: "3", Suit: "Spades"},
			{Rank: "3", Suit: "Diamonds"}},
		Type: types.TRIPLE}

	cardListB := types.CardList{
		Cards: []types.Card{
			{Rank: "2", Suit: "Spades"}},
		Type: types.SINGLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat13 is exported
func TestCanBeat13(t *testing.T) {
	cardListA := types.CardList{
		Cards: []types.Card{
			{Rank: "J", Suit: "Clubs"},
			{Rank: "J", Suit: "Spades"},
			{Rank: "J", Suit: "Diamonds"}},
		Type: types.TRIPLE}

	cardListB := types.CardList{
		Cards: []types.Card{
			{Rank: "10", Suit: "Clubs"},
			{Rank: "10", Suit: "Spades"},
			{Rank: "10", Suit: "Diamonds"}},
		Type: types.TRIPLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat14 is exported
func TestCanBeat14(t *testing.T) {
	cardListA := types.CardList{
		Cards: []types.Card{
			{Rank: "3", Suit: "Clubs"},
			{Rank: "3", Suit: "Hearts"},
			{Rank: "3", Suit: "Spades"},
			{Rank: "3", Suit: "Diamonds"}},
		Type: types.QUADRUPLE}

	cardListB := types.CardList{
		Cards: []types.Card{
			{Rank: "2", Suit: "Clubs"},
			{Rank: "2", Suit: "Diamonds"}},
		Type: types.COUPLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat15 is exported
func TestCanBeat15(t *testing.T) {
	cardListA := types.CardList{
		Cards: []types.Card{
			{Rank: "A", Suit: "Clubs"},
			{Rank: "A", Suit: "Hearts"},
			{Rank: "A", Suit: "Spades"},
			{Rank: "A", Suit: "Diamonds"}},
		Type: types.QUADRUPLE}

	cardListB := types.CardList{
		Cards: []types.Card{
			{Rank: "K", Suit: "Clubs"},
			{Rank: "K", Suit: "Hearts"},
			{Rank: "K", Suit: "Spades"},
			{Rank: "K", Suit: "Diamonds"}},
		Type: types.QUADRUPLE}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestGetRankValue1 is exported
func TestGetRankValue1(t *testing.T) {
	card := types.Card{Rank: "3", Suit: "Clubs"}
	expectedValue := card.GetRankValue()
	actualValue := 0

	if expectedValue != actualValue {
		t.Error("Invalid rank value. Expected:", expectedValue, "Actual:", actualValue)
	}
}

// TestGetRankValue2 is exported
func TestGetRankValue2(t *testing.T) {
	card := types.Card{Rank: "2", Suit: "Diamonds"}
	expectedValue := card.GetRankValue()
	actualValue := 12

	if expectedValue != actualValue {
		t.Error("Invalid rank value. Expected:", expectedValue, "Actual:", actualValue)
	}
}

// TestGetRankValue3 is exported
func TestGetRankValue3(t *testing.T) {
	card := types.Card{Rank: "1", Suit: "Diamonds"}
	expectedValue := card.GetRankValue()
	actualValue := -1

	if expectedValue != actualValue {
		t.Error("Invalid rank value. Expected:", expectedValue, "Actual:", actualValue)
	}
}

// TestGetSuitValue1 is exported
func TestGetSuitValue1(t *testing.T) {
	card := types.Card{Rank: "3", Suit: "Clubs"}
	expectedValue := card.GetSuitValue()
	actualValue := 0

	if expectedValue != actualValue {
		t.Error("Invalid rank value. Expected:", expectedValue, "Actual:", actualValue)
	}
}

// TestGetSuitValue2 is exported
func TestGetSuitValue2(t *testing.T) {
	card := types.Card{Rank: "2", Suit: "Spades"}
	expectedValue := card.GetSuitValue()
	actualValue := 3

	if expectedValue != actualValue {
		t.Error("Invalid rank value. Expected:", expectedValue, "Actual:", actualValue)
	}
}

// TestGetSuitValue3 is exported
func TestGetSuitValue3(t *testing.T) {
	card := types.Card{Rank: "A", Suit: "Smoi"}
	expectedValue := card.GetSuitValue()
	actualValue := -1

	if expectedValue != actualValue {
		t.Error("Invalid rank value. Expected:", expectedValue, "Actual:", actualValue)
	}
}
