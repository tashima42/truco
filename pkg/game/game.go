package game

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
)

type Game struct {
	id         string
	players    []*Player
	maxPlayers int
}

type Player struct {
	id   string
	name string
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
	player := Player{id: id, name: name}
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
