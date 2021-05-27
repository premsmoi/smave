package types

import (
	"math/rand"
	"time"
)

// RankList is exported
var RankList = [...]string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2"}

// SuitList is exprted
var SuitList = [...]string{"Clubs", "Diamonds", "Hearts", "Spades"}

// Card is exported
type Card struct {
	Rank string
	Suit string
}

// CardList is exported
type CardList []Card

// Remove is exported
func (cardList *CardList) Remove(index int) Card {
	cards := *cardList
	card := cards[index]
	// (*cardList)[len(cards)-1], (*cardList)[index] = cards[index], cards[len(cards)-1]
	// (*cardList) = (*cardList)[:len((*cardList))-1]
	(*cardList) = append((*cardList)[:index], (*cardList)[index+1:]...)
	return card
}

// Find is exported
func (cardList CardList) Find(card Card) (int, bool) {
	for i, item := range cardList {
		if item == card {
			return i, true
		}
	}
	return -1, false
}

// FindDuplicateCardList is exported
// func (cardList CardList) FindDuplicateCardList(targetCount int) (rCardList CardList) {
// 	cards := cardList.Cards
// 	count := 1
// 	for i := 0; i < len(cards); i++ {
// 		for j := i + 1; j < len(cards); j++ {
// 			if cards
// 		}
// 	}
// }

// GenerateAllCards is exported
func (cardList *CardList) GenerateAllCards() {
	for _, rank := range RankList {
		for _, suit := range SuitList {
			card := Card{Rank: rank, Suit: suit}
			*cardList = append(*cardList, card)
		}
	}
}

// Shuffle is exported
func (cardList *CardList) Shuffle() {
	cardListValue := *cardList
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*cardList), func(i, j int) {
		cardListValue[i], cardListValue[j] = cardListValue[j], cardListValue[i]
	})
}

// Draw is exported
func (cardList *CardList) Draw() Card {
	cards := *cardList
	card := cards[len(cards)-1]
	if len(cards) > 0 {
		*cardList = (*cardList)[:len(cards)-1]
	}
	return card
}

// CanBeat of CardList is exported
func (cardList CardList) CanBeat(cardListB CardList) bool {
	cardsA := cardList
	cardsB := cardListB
	lengthA := len(cardsA)
	lengthB := len(cardsB)

	if lengthA == 4 {
		if lengthB == 2 {
			return true
		} else if lengthB == 4 && cardsA[0].GetRankValue() > cardsB[0].GetRankValue() {
			return true
		}
	} else if lengthA == 3 {
		if lengthB == 1 {
			return true
		} else if lengthB == 3 && cardsA[0].GetRankValue() > cardsB[0].GetRankValue() {
			return true
		}
	} else if lengthA == 2 {
		if lengthB == 2 && cardsA[0].GetRankValue() > cardsB[0].GetRankValue() {
			return true
		}
		if lengthB == 2 && cardsA[0].GetRankValue() == cardsB[0].GetRankValue() &&
			(cardsA[0].Suit == "Spades" || cardsA[1].Suit == "Spades") {
			return true
		}
	} else if lengthA == 1 && lengthB == 1 {
		if cardsA[0].GetRankValue() > cardsB[0].GetRankValue() {
			return true
		} else if cardsA[0].GetRankValue() == cardsB[0].GetRankValue() && cardsA[0].GetSuitValue() > cardsB[0].GetSuitValue() {
			return true
		}
	}
	return false
}

// GetRankValue is exported
func (card Card) GetRankValue() int {
	for i, rank := range RankList {
		if rank == card.Rank {
			return i
		}
	}
	return -1
}

// GetSuitValue is exported
func (card Card) GetSuitValue() int {
	for i, suit := range SuitList {
		if suit == card.Suit {
			return i
		}
	}
	return -1
}
