package truco

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	ErrGameFull              = errors.New("the game has reached the maximum amount of players")
	ErrNameTooLong           = errors.New("player name has more than 100 characters")
	ErrNameTooShort          = errors.New("player name has less than 2 characters")
	ErrPlayerAlreadyInGame   = errors.New("player is already in the game")
	ErrPlayerNotFound        = errors.New("player id not found")
	ErrNotEnoughPlayers      = errors.New("not enough players to start the game")
	ErrGameNotRunning        = errors.New("game is not running")
	ErrNotPlayerTurn         = errors.New("it's not the player's turn")
	ErrPlayerDoesNotHaveCard = errors.New("player does not have the card")
)

// Actions
type Action int

const (
	PlayerOnePoint Action = iota
	PlayerTwoPoint
	Draw
	PlayerOneWin
	PlayerTwoWin
)

type Game struct {
	id            string
	manilha       string
	deckWeights   map[Card]int
	currentPlayer *Player
	deck          []Card
	pile          [][]Card
	players       []*Player
	maxPlayers    int
	cardPointer   int
	round         int
	seed1         uint64
	seed2         uint64
	nextPlayer    int
	running       bool
	lastAction    Action
}

type Player struct {
	id     string
	name   string
	cards  []Card
	points int
}

func NewGame() (*Game, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	game := Game{
		id:            id,
		maxPlayers:    2,
		players:       make([]*Player, 0),
		deck:          nil,
		deckWeights:   DefaultDeckWeights(),
		cardPointer:   0,
		currentPlayer: nil,
		round:         0,
		seed1:         0,
		seed2:         0,
		nextPlayer:    0,
		pile:          make([][]Card, 1),
	}
	return &game, nil
}

func NewPlayer(name string) (*Player, error) {
	if len(name) > 100 {
		return nil, ErrNameTooLong
	}
	if len(name) < 2 {
		return nil, ErrNameTooShort
	}
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	player := Player{id: id, name: name, points: 0, cards: make([]Card, 3)}
	return &player, nil
}

func (g *Game) Seed(seed1, seed2 uint64) {
	g.seed1 = seed1
	g.seed2 = seed2
}

func (g *Game) AddPlayer(player *Player) error {
	if g.maxPlayers == len(g.players) {
		return ErrGameFull
	}
	for _, p := range g.players {
		if p.id == player.id {
			return ErrPlayerAlreadyInGame
		}
	}
	g.players = append(g.players, player)
	return nil
}

func (g *Game) RemovePlayer(player *Player) error {
	removePlayerPosition := -1
	for i, p := range g.players {
		if p == nil {
			continue
		}
		if p.id == player.id {
			removePlayerPosition = i
			break
		}
	}
	if removePlayerPosition == -1 {
		return ErrPlayerNotFound
	}
	g.players[removePlayerPosition] = nil
	return nil
}

func (g *Game) Start() error {
	if len(g.players) != g.maxPlayers {
		return ErrNotEnoughPlayers
	}

	g.deck = ShuffledDeck(g.seed1, g.seed2)
	g.currentPlayer = g.players[0]
	g.setManilha()
	g.drawCards()
	g.running = true

	return nil
}

func (g *Game) setManilha() {
	g.manilha = string(g.deck[0])
	g.cardPointer += 1

	cardID := string(g.manilha[1])

	spades := Card(Spades + cardID)
	hearts := Card(Hearts + cardID)
	diamonds := Card(Diamonds + cardID)
	clubs := Card(Clubs + cardID)

	// add weight to the cards based on their suit
	g.deckWeights[clubs] += 10
	g.deckWeights[diamonds] += 11
	g.deckWeights[hearts] += 12
	g.deckWeights[spades] += 13
}

func (g *Game) drawCards() {
	for _, player := range g.players {
		for i := 0; i < 3; i++ {
			player.cards[i] = g.deck[g.cardPointer]
			g.cardPointer += 1
		}
	}
}

func (p *Player) hasCard(card Card) bool {
	for _, c := range p.cards {
		if c == card {
			return true
		}
	}
	return false
}

func (g *Game) Play(player *Player, card Card) error {
	if !g.running {
		return ErrGameNotRunning
	}
	if player.id != g.currentPlayer.id {
		return ErrNotPlayerTurn
	}
	if !player.hasCard(card) {
		return ErrPlayerDoesNotHaveCard
	}
	// play the card
	g.playCard(player, card)

	// only run the win check if it is the last turn
	if len(g.pile[g.round]) == 2 {
		// if the player wins the round, they start the next round
		// compare the current card with the previous played card
		switch g.compareCards(card, g.pile[g.round][0]) {
		case 1:
			g.players[g.nextPlayer].points += 1
			g.lastAction = Action(g.nextPlayer)
		case -1:
			g.players[g.nextPlayer].points += 1
			g.lastAction = Action(g.nextPlayer)
		case 0:
			g.lastAction = Draw
		}
	}

	// if any of the players reached 12 points or more, declare him as the winner
	if g.players[0].points >= 12 {
		g.running = false
		g.lastAction = PlayerOneWin
	}
	if g.players[1].points >= 12 {
		g.running = false
		g.lastAction = PlayerTwoWin
	}

	// next player
	g.nextPlayer += 1
	if g.nextPlayer == len(g.players) {
		g.nextPlayer = 0
	}
	g.currentPlayer = g.players[g.nextPlayer]

	g.round += 1
	return nil
}

func (g *Game) compareCards(card1, card2 Card) int {
	if g.deckWeights[card1] > g.deckWeights[card2] {
		return 1
	}
	if g.deckWeights[card1] < g.deckWeights[card2] {
		return -1
	}
	return 0
}

func (g *Game) playCard(player *Player, card Card) {
	// remove card from player
	for i, c := range player.cards {
		if c == card {
			player.cards = append(player.cards[:i], player.cards[i+1:]...)
			break
		}
	}
	// add card to pile
	g.pile[g.round] = append(g.pile[g.round], card)
}

func (g *Game) CurrentPlayer() *Player {
	return g.currentPlayer
}

func (p *Player) Cards() []Card {
	return p.cards
}
