package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/tashima42/truco/pkg/truco"
)

func main() {
	if err := runGame(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}

func runGame() error {
	g, err := truco.NewGame()
	if err != nil {
		return errors.New("failed to create game: " + err.Error())
	}
	g.Seed(123, 456)
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

	fmt.Printf("manilha: %s\n", g.Manilha().Unicode())
	printCards(p1)
	printCards(p2)
	fmt.Printf("%s: playing card ( %s )\n", p1.Name(), p1.Cards()[0].Unicode())
	if err := g.Play(p1, p1.Cards()[0]); err != nil {
		return err
	}

	fmt.Printf("%s: playing card ( %s )\n", p2.Name(), p2.Cards()[0].Unicode())
	if err := g.Play(p2, p2.Cards()[0]); err != nil {
		return err
	}

	lp := g.LastPoint()
	if lp == nil {
		fmt.Println("draw")
	} else {
		fmt.Printf("point: %s\n", lp.Name())
	}

	fmt.Printf("manilha: %s\n", g.Manilha().Unicode())
	printCards(p2)
	printCards(p1)
	fmt.Printf("%s: playing card ( %s )\n", p2.Name(), p2.Cards()[0].Unicode())
	if err := g.Play(p2, p2.Cards()[0]); err != nil {
		return err
	}

	fmt.Printf("%s: playing card ( %s )\n", p1.Name(), p1.Cards()[0].Unicode())
	if err := g.Play(p1, p1.Cards()[0]); err != nil {
		return err
	}

	lp = g.LastPoint()
	if lp == nil {
		fmt.Println("draw")
	} else {
		fmt.Printf("point: %s\n", lp.Name())
	}

	return nil
}

func printCards(player *truco.Player) {
	fmt.Print(player.Name() + ": cards ")
	for _, c := range player.Cards() {
		fmt.Print(" " + c.Unicode())
	}
	fmt.Print("\n")
}
