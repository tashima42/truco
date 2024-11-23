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
	ErrNotPlayerTurn         = errors.New("not the player's turn")
	ErrPlayerDoesNotHaveCard = errors.New("player does not have the card")
)

// Actions
type Action int

type Game struct {
	// ID of the game
	id string
	// registered players
	players []*Player
	// max number of players per game
	maxPlayers int
	// first seed for the random number generator
	seed1 uint64
	// second seed for the random number generator
	seed2 uint64
	// true if the game has already started and not ended yet
	running bool
	// state of hands of rounds
	hands []*Hand
}

type Hand struct {
	// deck of cards
	deck []Card
	// manilha card
	manilha Card
	// played cards in order
	pile []Card
	// weight of each card
	deckWeights map[Card]int
	// who won the round 0 = draw, 1 = player 1, 2 = player 2
	points []int
	// current round
	round uint
	// -1 = draw, 0 = player 1, 1 = player 2
	wonPosition int
	// next card to pull from the deck
	deckPosition uint
	// player who will play the next card
	currentPlayer uint
}

type Player struct {
	id    string
	name  string
	cards []Card
}

func NewGame() (*Game, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	game := Game{
		id:         id,
		maxPlayers: 2,
		players:    make([]*Player, 0),
		seed1:      0,
		seed2:      0,
		hands:      []*Hand{newHand()},
	}
	return &game, nil
}

func newHand() *Hand {
	return &Hand{
		deck:          nil,
		pile:          make([]Card, 0),
		deckWeights:   DefaultDeckWeights(),
		currentPlayer: 0,
		deckPosition:  0,
		round:         0,
		points:        make([]int, 3),
	}
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
	player := Player{
		id:    id,
		name:  name,
		cards: make([]Card, 0),
	}
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

	g.startHand()
	g.running = true

	return nil
}

func (g *Game) hand() *Hand {
	return g.hands[len(g.hands)-1]
}

func (g *Game) startHand() error {
	currentHand := len(g.hands) - 1
	if g.seed2 != 0 {
		g.seed2 += uint64(currentHand)
	}
	g.hand().deck = ShuffledDeck(g.seed1, g.seed2)
	if err := g.hand().setManilha(); err != nil {
		return err
	}
	g.drawCards()
	return nil
}

func (h *Hand) setManilha() error {
	h.manilha = Card(h.deck[0])
	h.deckPosition += 1

	cardID := string(h.manilha[1])
	manilhaID, err := nextCardID(cardID)
	if err != nil {
		return err
	}

	spades := Card(Spades + manilhaID)
	hearts := Card(Hearts + manilhaID)
	diamonds := Card(Diamonds + manilhaID)
	clubs := Card(Clubs + manilhaID)

	// add weight to the cards based on their suit
	h.deckWeights[clubs] += 10
	h.deckWeights[diamonds] += 11
	h.deckWeights[hearts] += 12
	h.deckWeights[spades] += 13

	return nil
}

func (g *Game) drawCards() {
	for _, player := range g.players {
		for i := 0; i < 3; i++ {
			player.cards = append(player.cards, g.hand().deck[g.hand().deckPosition])
			g.hand().deckPosition += 1
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
	if player.id != g.CurrentPlayer().id {
		return ErrNotPlayerTurn
	}
	if !player.hasCard(card) {
		return ErrPlayerDoesNotHaveCard
	}
	// play the card
	g.hand().playCard(player, card)

	// only check who won the round on even number of cards
	if len(g.hand().pile) != 0 && len(g.hand().pile)%2 == 0 {
		// if the player wins the round, they start the next round
		// compare the current card with the previous played card
		compare := g.hand().compareCards(card, g.hand().pile[len(g.hand().pile)-2])
		switch compare {
		// case 1 means the current card is greater than the previous card
		case 1:
			g.hand().points[g.hand().round] = int(g.hand().currentPlayer)
			// case 2 means the current card is less than the previous card
		case 2:
			g.hand().currentPlayer = g.hand().currentPlayer ^ 1
			g.hand().points[g.hand().round] = int(g.hand().currentPlayer)
			// case 0 means the current card is equal to the previous card
		case 0:
			g.hand().points[g.hand().round] = -1
			g.hand().currentPlayer = uint(g.hand().points[0])
		}

		g.hand().round += 1
	} else {
		// toggle player between zero and one
		g.hand().currentPlayer = g.hand().currentPlayer ^ 1
	}

	// check if the hand is over
	if g.hand().round == 3 {
		// cehck who won the hand
		playerOnePoints := 0
		playerTwoPoints := 0
		for _, point := range g.hand().points {
			if point == 0 {
				playerOnePoints += 1
			} else if point == 1 {
				playerTwoPoints += 1
			}
		}
		if playerOnePoints > playerTwoPoints {
			g.hand().wonPosition = 0
		} else if playerTwoPoints > playerOnePoints {
			g.hand().wonPosition = 1
		} else {
			g.hand().wonPosition = g.hand().points[0]
		}
		g.hands = append(g.hands, newHand())
		if err := g.startHand(); err != nil {
			return err
		}
	}

	playerOneHands := 0
	playerTwoHands := 0
	for _, hand := range g.hands {
		if hand.wonPosition == 0 {
			playerOneHands += 1
		} else if hand.wonPosition == 1 {
			playerTwoHands += 1
		}
	}

	if playerOneHands == 12 || playerTwoHands == 12 {
		g.running = false
	}

	return nil
}

// compareCards compares two cards and returns:
// 1 if card1 weight is greater than card2
// 2 if card2 weight is greater than card1
// 0 if they are equal
func (h *Hand) compareCards(card1, card2 Card) int {
	deckWeightOne := h.deckWeights[card1]
	deckWeightTwo := h.deckWeights[card2]
	if deckWeightOne > deckWeightTwo {
		return 1
	}
	if deckWeightOne < deckWeightTwo {
		return 2
	}
	return 0
}

func (h *Hand) playCard(player *Player, card Card) {
	// remove card from player
	for i, c := range player.cards {
		if c == card {
			player.cards = append(player.cards[:i], player.cards[i+1:]...)
			break
		}
	}
	h.pile = append(h.pile, card)
}

func (g *Game) CurrentPlayer() *Player {
	return g.players[g.hand().currentPlayer]
}

func (g *Game) Finished() bool {
	return !g.running
}

func (p *Player) Cards() []Card {
	return p.cards
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) ID() string {
	return p.id
}

func (g *Game) LastPoint() *Player {
	if g.hand().round == 0 {
		if len(g.hands) == 1 {
			return nil
		}
		previousHand := g.hands[len(g.hands)-2]
		if previousHand.points[previousHand.round-1] == -1 {
			return nil
		}
		return g.players[previousHand.points[previousHand.round-1]]
	}
	if g.hand().points[g.hand().round-1] == -1 {
		return nil
	}
	return g.players[g.hand().points[g.hand().round-1]]
}

func (g *Game) Winner() *Player {
	if g.hand().wonPosition == -1 {
		return nil
	}
	return g.players[g.hand().wonPosition]
}

func (g *Game) Running() bool {
	return g.running
}

func (g *Game) Manilha() Card {
	return Card(g.hand().manilha)
}
