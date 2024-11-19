package main

import (
	"errors"
	"log"
	"os"

	"github.com/tashima42/truco/pkg/truco"
)

func main() {
	if err := runGame(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func runGame() error {
	g, err := truco.NewGame()
	if err != nil {
		return errors.New("failed to create game: " + err.Error())
	}
	p1, err := truco.NewPlayer("player 1")
	if err != nil {
		return errors.New("failed to create player: " + err.Error())
	}
	p2, err := truco.NewPlayer("player 2")
	if err != nil {
		return errors.New("failed to create player: " + err.Error())
	}

	if err := g.AddPlayer(p1); err != nil {
		return errors.New("failed to add player 1 to game: " + err.Error())
	}
	if err := g.AddPlayer(p2); err != nil {
		return errors.New("failed to add player 2 to game: " + err.Error())
	}

	if err := g.Start(); err != nil {
		return errors.New("failed to start game: " + err.Error())
	}

	if err := g.Play(p1, p1.Cards()[0]); err != nil {
		return err
	}

	return nil
}
