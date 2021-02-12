package types

// RankList is exported
var RankList = [...]string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2"}

// SuitList is exprted
var SuitList = [...]string{"Clubs", "Diamonds", "Hearts", "Spades"}

// Card set type are exprted
const (
	NORMAL    = "NORMAL"
	SINGLE    = "SINGLE"
	COUPLE    = "COUPLE"
	TRIPLE    = "TRIPLE"
	QUADRUPLE = "QUADRUPLE"
)

// Card is exported
type Card struct {
	Rank string
	Suit string
}

// CardList is exported
type CardList struct {
	Cards []Card
	Type  string
}

// PlayableCardList is exported
type PlayableCardList struct {
	Cards []Card
	Type  string
}

// Find is exported
func (cardList CardList) Find(card Card) (int, bool) {
	for i, item := range cardList.Cards {
		if item == card {
			return i, true
		}
	}
	return -1, false
}

// GenerateAllCards is exported
func (cardList *CardList) GenerateAllCards() {
	cardList.Type = NORMAL
	for _, rank := range RankList {
		for _, suit := range SuitList {
			card := Card{Rank: rank, Suit: suit}
			cardList.Cards = append(cardList.Cards, card)
		}
	}
}

// CanBeat of CardList is exported
func (cardList CardList) CanBeat(cardListB CardList) bool {
	cardsA := cardList.Cards
	cardsB := cardListB.Cards
	typeA := cardList.Type
	typeB := cardListB.Type

	if typeA == QUADRUPLE {
		if typeB == COUPLE {
			return true
		} else if typeB == QUADRUPLE && cardsA[0].GetRankValue() > cardsB[0].GetRankValue() {
			return true

		}
	} else if typeA == TRIPLE {
		if typeB == SINGLE {
			return true
		} else if typeB == TRIPLE && cardsA[0].GetRankValue() > cardsB[0].GetRankValue() {
			return true
		}
	} else if typeA == COUPLE {
		if typeB == COUPLE && cardsA[0].GetRankValue() > cardsB[0].GetRankValue() {
			return true
		}
		if typeB == COUPLE && cardsA[0].GetRankValue() == cardsB[0].GetRankValue() &&
			(cardsA[0].Suit == "Spades" || cardsA[1].Suit == "Spades") {
			return true
		}
	} else if typeA == SINGLE && typeB == SINGLE {
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
