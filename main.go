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

	for i := -0; i < 20; i++ {
		fmt.Printf("=================== HAND %d ===================\n", i)
		fmt.Printf("manilha: %s\n", g.Manilha().Unicode())
		for j := i; j <= 6; j++ {
			cp := g.CurrentPlayer()
			printCards(cp)
			fmt.Printf("%s: playing card ( %s )\n", cp.Name(), cp.Cards()[0].Unicode())
			if err := g.Play(cp, cp.Cards()[0]); err != nil {
				return err
			}
			if j%2 == 0 {
				lp := g.LastPoint()
				if lp == nil {
					fmt.Println("point: draw")
				} else {
					fmt.Printf("point: %s\n", lp.Name())
				}
				fmt.Printf("---------------------- ROUND %d ----------------------\n", j/2)
			}
		}
		winner := g.Winner()
		if winner == nil {
			fmt.Println("won: draw")
		} else {
			fmt.Printf("won: %s\n", winner.Name())
		}
		if !g.Running() {
			fmt.Println("game finished")
			break
		}
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
