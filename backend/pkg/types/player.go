package types

import "sort"

// Role
const (
	King      = "King"
	Queen     = "Queen"
	ViceSlave = "ViceSlave"
	Slave     = "Slave"
	People    = "People"
)

// Player is exported
type Player struct {
	CardList CardList
	Role     string
	Name     string
	Pass     bool
	Complete bool
}

// AddCard is exported
func (p *Player) AddCard(card Card) {
	p.CardList = append(p.CardList, card)
}

func (p Player) RandomToBeat(cardList CardList) {

}

// Play is exported
func (p *Player) Play(cardIndexList []int) CardList {
	if len(p.CardList) == 0 {
		return CardList{}
	}
	playingCardList := CardList{}
	for _, cardIndex := range cardIndexList {
		playingCardList = append(playingCardList, p.CardList[cardIndex])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cardIndexList)))

	for _, cardIndex := range cardIndexList {
		p.CardList.Remove(cardIndex)
	}

	return playingCardList
}

// Belows are the logics for non-player

// FindPlayableCards is exported
func (p *Player) FindPlayableCards(playingCards CardList) []int {
	n := len(playingCards)
	m := make(map[string][]int)
	if n == 0 {
		return []int{0}
	}
	for i := 0; i < len(p.CardList); i++ {
		m[p.CardList[i].Rank] = append(m[p.CardList[i].Rank], i)
	}

	for _, cardIndexList := range m {
		if len(cardIndexList) >= n {
			// fmt.Println("Key:", key, "cardIndexList:", cardIndexList)
			cardList := CardList{}
			for _, cardIndex := range cardIndexList {
				cardList = append(cardList, p.CardList[cardIndex])
			}
			if cardList[:n].CanBeat(playingCards) {
				// fmt.Println(cardList[:n], "can beat", playingCards)
				return cardIndexList[:n]
			} else if len(cardIndexList) >= n+2 && cardList[:n+2].CanBeat(playingCards) {
				// fmt.Println(cardList[:n+2], "can beat", playingCards)
				return cardIndexList[:n+2]
			}
		}
	}
	return []int{}
}
