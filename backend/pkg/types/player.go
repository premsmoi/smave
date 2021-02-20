package types

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
	for i := 0; i < len(cardIndexList); i++ {
		playingCardList = append(playingCardList, p.CardList[cardIndexList[i]])
		p.CardList.Remove(cardIndexList[i])
	}
	return playingCardList
}
