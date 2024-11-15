package truco

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/tashima42/truco/pkg/rand"
)

type RoundResult int

const (
	RoundResultWon  RoundResult = 0
	RoundResultLost RoundResult = 1
	RoundResultDraw RoundResult = 2
)

type Game struct {
	ID                   int
	teams                []*Team
	rounds               []*Round
	currentRoundPosition int
	numberOfPlayers      int
}

type Round struct {
	ID                    int
	number                int
	vira                  Card
	manilha               Card
	playerOrder           []teamPlayer
	currentPlayerPosition int
	discardedCards        []playerCard
	currentCards          []playerCard
}

type teamPlayer struct {
	teamID   int
	playerID int
}

type playerCard struct {
	card     Card
	playerID int
}

type Team struct {
	ID          int
	players     []*Player
	gamePoints  int
	roundPoints int
}

type Player struct {
	ID        int
	name      string
	cards     []Card
	usedCards map[int]bool
}

func NewPlayer(ID int, name string) *Player {
	return &Player{
		ID:    ID,
		name:  name,
		cards: make([]Card, 3),
	}
}

func NewTeam(ID int, players []*Player) *Team {
	return &Team{
		ID:          ID,
		players:     players,
		gamePoints:  0,
		roundPoints: 0,
	}
}

func NewGame(ID int) *Game {
	return &Game{
		ID:                   ID,
		teams:                []*Team{},
		rounds:               []*Round{},
		currentRoundPosition: -1,
		numberOfPlayers:      0,
	}
}

func (g *Game) AddTeams(teams [2]*Team) error {
	if len(teams) != 2 {
		return errors.New("teams length must be 2, got: " + strconv.Itoa(len(teams)))
	}
	for i := 0; i < 2; i++ {
		g.teams = append(g.teams, teams[i])
		g.numberOfPlayers += len(teams[i].players)
	}
	return nil
}

func (g *Game) NewRound() (*Round, error) {
	log.Println("generating vira")
	vira := Cards[int(rand.Int(39))]
	log.Println("generating manilha")
	manilha := Cards[int(rand.Int(39))]
	for vira.ID == manilha.ID {
		log.Println("manilha is the same as vira, generating a new one")
		manilha = Cards[int(rand.Int(39))]
	}

	playerOrder := []teamPlayer{}
	drawedCards := map[int]bool{}

	for _, t := range g.teams {
		for _, p := range t.players {
			log.Println("inserting in player order")
			playerOrder = append(playerOrder, teamPlayer{teamID: t.ID, playerID: p.ID})
			log.Println("generating random cards")
			nc := getRandomCards(drawedCards)
			for i := 0; i < 3; i++ {
				log.Println("setting card: " + strconv.Itoa(int(nc[i].ID)))
				drawedCards[nc[i].ID] = true
			}
			p.cards = nc
		}
	}

	for _, t := range g.teams {
		t.roundPoints = 0
	}

	round := &Round{
		vira:           vira,
		manilha:        manilha,
		number:         g.currentRoundPosition,
		playerOrder:    playerOrder,
		currentCards:   []playerCard{},
		discardedCards: []playerCard{},
	}

	g.currentRoundPosition += +1
	g.rounds = append(g.rounds, round)
	return round, nil
}

func (g *Game) Play(cardPosition int) error {
	log.Println("playing")
	r := g.currentRound()
	tp := r.playerOrder[r.currentPlayerPosition]
	r.currentPlayerPosition += 1
	t, err := g.getTeam(tp.teamID)
	if err != nil {
		return err
	}
	p, err := t.getPlayer(tp.playerID)
	if err != nil {
		return err
	}
	if _, ok := p.usedCards[cardPosition]; ok {
		return fmt.Errorf("card already used: %d", cardPosition)
	}
	r.discardedCards = append(r.discardedCards, playerCard{card: p.cards[cardPosition], playerID: p.ID})
	log.Println("discarded length: ", len(r.discardedCards))
	if len(r.discardedCards) <= 1 {
		return nil
	}
	currentCard := r.discardedCards[len(r.discardedCards)-1].card
	lastCard := r.discardedCards[len(r.discardedCards)-2].card

	log.Printf("Current Player played: %+v - Last Player played: %+v", currentCard, lastCard)

	currentCardManilha := r.manilha.Value == currentCard.Value
	lastCardManilha := r.manilha.Value == lastCard.Value

	result := RoundResultDraw

	if currentCard.Value > lastCard.Value {
		log.Println("value better won")
		result = RoundResultWon
	}
	if currentCard.Value < lastCard.Value {
		log.Println("value worst lost")
		result = RoundResultLost
	}
	if currentCardManilha && !lastCardManilha {
		log.Println("manilha won")
		result = RoundResultWon
	}

	otp := r.playerOrder[r.currentPlayerPosition-2]
	ot, err := g.getTeam(otp.teamID)
	if err != nil {
		return err
	}

	if lastCardManilha && !currentCardManilha {
		log.Println("manilha lost")
		result = RoundResultLost
	}

	if lastCardManilha && currentCardManilha {
		log.Println("both manilha")
		if currentCard.Suit > lastCard.Suit {
			log.Println("both manilha - won suit")
			result = RoundResultWon
		} else {
			log.Println("both manilha - lost suit")
			result = RoundResultLost
		}
	}

	switch result {
	case RoundResultWon:
		log.Println("running points - increate current")
		t.increaseRoundPoints(1)
		return nil
	case RoundResultLost:
		log.Println("running points - increate last")
		ot.increaseRoundPoints(1)
		return nil
	default:
		log.Println("running points - do nothing")
		return nil
	}
}

func (g *Game) currentRound() *Round {
	return g.rounds[g.currentRoundPosition]
}

func (g *Game) getTeam(ID int) (*Team, error) {
	for _, t := range g.teams {
		if t.ID == ID {
			return t, nil
		}
	}
	return nil, fmt.Errorf("team with ID '%d' not found", ID)
}

func (t *Team) getPlayer(ID int) (*Player, error) {
	for _, p := range t.players {
		if p.ID == ID {
			return p, nil
		}
	}
	return nil, fmt.Errorf("player with ID '%d' not found", ID)
}

func getRandomCards(usedCards map[int]bool) []Card {
	cards := make([]Card, 3)

	for i := range cards {
		log.Println("generating random card")
		c := int(rand.Int(39))
		_, ok := usedCards[c]
		for ok {
			log.Println("card already set, generating a new one")
			c = int(rand.Int(39))
			_, ok = usedCards[c]
		}
		usedCards[c] = true
		cards[i] = Cards[c]
	}

	return cards
}

func (g *Game) PrintState() {
	log.Printf("Game ID: %d", g.ID)
	log.Printf("Current Round position: %d", g.currentRoundPosition)
	log.Printf("Number of players: %d", g.numberOfPlayers)

	for _, t := range g.teams {
		log.Printf("Team [%d] - game points: %d", t.ID, t.gamePoints)
		log.Printf("Team [%d] - round points: %d", t.ID, t.roundPoints)
		log.Printf("Team [%d] - players: %+v", t.ID, t.players)
	}

	r := g.rounds[g.currentRoundPosition]
	log.Printf("Round [%d] - %+v", r.ID, r.discardedCards)
}

func (t *Team) increaseRoundPoints(points int) {
	t.roundPoints = t.roundPoints + points
}

func (p *Player) markCardUsed(cardPosition int) {
	p.usedCards[cardPosition] = true
}
