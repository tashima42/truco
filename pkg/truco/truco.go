package truco

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	ErrGameFull            = errors.New("the game has reached the maximum amount of players")
	ErrNameTooLong         = errors.New("player name has more than 100 characters")
	ErrNameTooShort        = errors.New("player name has less than 2 characters")
	ErrPlayerAlreadyInGame = errors.New("player is already in the game")
	ErrPlayerNotFound      = errors.New("player id not found")
	ErrNotEnoughPlayers    = errors.New("not enough players to start the game")
)

type Game struct {
	id            string
	manilha       string
	deckWeights   map[Card]int
	currentPlayer *Player
	deck          []Card
	players       []*Player
	maxPlayers    int
	cardPointer   int
	round         int
}

type Player struct {
	id     string
	name   string
	cards  []Card
	points int
}

func NewGame(seed1, seed2 uint64) (*Game, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	game := Game{
		id:            id,
		maxPlayers:    2,
		players:       make([]*Player, 0),
		deck:          ShuffledDeck(seed1, seed2),
		deckWeights:   DefaultDeckWeights(),
		cardPointer:   0,
		currentPlayer: nil,
		round:         0,
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

	g.currentPlayer = g.players[0]
	g.setManilha()
	g.drawCards()
	g.round += 1

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