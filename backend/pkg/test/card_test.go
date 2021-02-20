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
	actualLength := len(generatedCardList)

	if expecetedLength != actualLength {
		t.Error("Length is invalid. Expected:", expecetedLength, "Actual:", actualLength)
	}

	for _, card := range ExpectedCardList {
		i, _ := generatedCardList.Find(card)
		if i == -1 {
			t.Error(card, "is missing.")
		}
	}
}

// TestShuffle is exported
func TestShuffle(t *testing.T) {
	isChanged := false
	var initialCardList = types.CardList{}
	initialCardList.GenerateAllCards()
	var shuffledCardList = types.CardList{}
	shuffledCardList.GenerateAllCards()
	shuffledCardList.Shuffle()

	for i := 0; i < len(initialCardList); i++ {
		if initialCardList[i] != shuffledCardList[i] {
			isChanged = true
		}
	}

	if isChanged == false {
		t.Error("Shuffle method is not working. generatedCardList not changed.")
	}
}

// TestRemove is exported
func TestRemove(t *testing.T) {
	cards := types.CardList{
		{Rank: "A", Suit: "Clubs"},
		{Rank: "4", Suit: "Spades"},
		{Rank: "10", Suit: "Diamonds"}}
	removedCard := cards.Remove(0)
	expectedCard := types.Card{Rank: "A", Suit: "Clubs"}

	if removedCard != expectedCard {
		t.Error("Remove method is not working. Expected:", expectedCard, "Actual:", removedCard)
	}

	actualLenth := len(cards)
	expectedLength := 2

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}
}

// TestDraw is exported
func TestDraw(t *testing.T) {
	deck := types.CardList{
		{Rank: "A", Suit: "Clubs"},
		{Rank: "4", Suit: "Spades"},
		{Rank: "10", Suit: "Diamonds"}}
	drawnCard := deck.Draw()
	expextedCard := types.Card{Rank: "10", Suit: "Diamonds"}

	if drawnCard != expextedCard {
		t.Error("Draw method is not working. Expected:", expextedCard, "Actual:", drawnCard)
	}

	actualLenth := len(deck)
	expectedLength := 2

	if actualLenth != expectedLength {
		t.Error("New length is invalid. Expected:", expectedLength, "Actual:", actualLenth)
	}
}

// TestFind1 is exported
func TestFind1(t *testing.T) {
	cardList := types.CardList{
		{Rank: "A", Suit: "Clubs"},
		{Rank: "5", Suit: "Spades"},
		{Rank: "6", Suit: "Diamonds"}}
	card := types.Card{Rank: "A", Suit: "Clubs"}
	i, res := cardList.Find(card)
	if i != 0 || res == false {
		t.Error("Find method returned invalid value.")
	}
}

// TestFind2 is exported
func TestFind2(t *testing.T) {
	cardList := types.CardList{
		{Rank: "A", Suit: "Clubs"},
		{Rank: "5", Suit: "Spades"},
		{Rank: "6", Suit: "Diamonds"}}
	card := types.Card{Rank: "A", Suit: "Spades"}
	i, res := cardList.Find(card)
	if i != -1 || res == true {
		t.Error("Find method returned invalid value.")
	}
}

// TestCanBeat1 is exported
func TestCanBeat1(t *testing.T) {
	cardListA := types.CardList{{Rank: "A", Suit: "Spades"}}
	cardListB := types.CardList{{Rank: "A", Suit: "Hearts"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat2 is exported
func TestCanBeat2(t *testing.T) {
	cardListA := types.CardList{{Rank: "2", Suit: "Spades"}}
	cardListB := types.CardList{{Rank: "2", Suit: "Diamonds"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat3 is exported
func TestCanBeat3(t *testing.T) {
	cardListA := types.CardList{{Rank: "3", Suit: "Spades"}}
	cardListB := types.CardList{{Rank: "3", Suit: "Clubs"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat4 is exported
func TestCanBeat4(t *testing.T) {
	cardListA := types.CardList{{Rank: "4", Suit: "Hearts"}}
	cardListB := types.CardList{{Rank: "4", Suit: "Diamonds"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat5 is exported
func TestCanBeat5(t *testing.T) {
	cardListA := types.CardList{{Rank: "5", Suit: "Hearts"}}
	cardListB := types.CardList{{Rank: "5", Suit: "Clubs"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat6 is exported
func TestCanBeat6(t *testing.T) {
	cardListA := types.CardList{{Rank: "6", Suit: "Diamonds"}}
	cardListB := types.CardList{{Rank: "6", Suit: "Clubs"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat7 is exported
func TestCanBeat7(t *testing.T) {
	cardListA := types.CardList{{Rank: "2", Suit: "Diamonds"}}
	cardListB := types.CardList{{Rank: "3", Suit: "Diamonds"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat8 is exported
func TestCanBeat8(t *testing.T) {
	cardListA := types.CardList{{Rank: "2", Suit: "Diamonds"}}
	cardListB := types.CardList{{Rank: "3", Suit: "Spades"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat9 is exported
func TestCanBeat9(t *testing.T) {
	cardListA := types.CardList{{Rank: "4", Suit: "Diamonds"}}
	cardListB := types.CardList{{Rank: "3", Suit: "Spades"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat10 is exported
func TestCanBeat10(t *testing.T) {
	cardListA := types.CardList{
		{Rank: "3", Suit: "Clubs"},
		{Rank: "3", Suit: "Spades"}}

	cardListB := types.CardList{
		{Rank: "3", Suit: "Diamonds"},
		{Rank: "3", Suit: "Hearts"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat11 is exported
func TestCanBeat11(t *testing.T) {
	cardListA := types.CardList{
		{Rank: "2", Suit: "Clubs"},
		{Rank: "2", Suit: "Diamonds"}}

	cardListB := types.CardList{
		{Rank: "3", Suit: "Spades"},
		{Rank: "3", Suit: "Hearts"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat12 is exported
func TestCanBeat12(t *testing.T) {
	cardListA := types.CardList{
		{Rank: "3", Suit: "Clubs"},
		{Rank: "3", Suit: "Spades"},
		{Rank: "3", Suit: "Diamonds"}}

	cardListB := types.CardList{
		{Rank: "2", Suit: "Spades"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat13 is exported
func TestCanBeat13(t *testing.T) {
	cardListA := types.CardList{
		{Rank: "J", Suit: "Clubs"},
		{Rank: "J", Suit: "Spades"},
		{Rank: "J", Suit: "Diamonds"}}

	cardListB := types.CardList{
		{Rank: "10", Suit: "Clubs"},
		{Rank: "10", Suit: "Spades"},
		{Rank: "10", Suit: "Diamonds"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat14 is exported
func TestCanBeat14(t *testing.T) {
	cardListA := types.CardList{
		{Rank: "3", Suit: "Clubs"},
		{Rank: "3", Suit: "Hearts"},
		{Rank: "3", Suit: "Spades"},
		{Rank: "3", Suit: "Diamonds"}}

	cardListB := types.CardList{
		{Rank: "2", Suit: "Clubs"},
		{Rank: "2", Suit: "Diamonds"}}

	if cardListA.CanBeat(cardListB) == false {
		t.Error("CanBeat method returned invalid value.")
	}
}

// TestCanBeat15 is exported
func TestCanBeat15(t *testing.T) {
	cardListA := types.CardList{
		{Rank: "A", Suit: "Clubs"},
		{Rank: "A", Suit: "Hearts"},
		{Rank: "A", Suit: "Spades"},
		{Rank: "A", Suit: "Diamonds"}}

	cardListB := types.CardList{
		{Rank: "K", Suit: "Clubs"},
		{Rank: "K", Suit: "Hearts"},
		{Rank: "K", Suit: "Spades"},
		{Rank: "K", Suit: "Diamonds"}}

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
