package main

import (
	"log"

	"github.com/tashima42/truco/pkg/truco"
)

func main() {
	log.Println("starting game")
	g := truco.NewGame(0)

	log.Println("creating players")
	p1 := truco.NewPlayer(0, "Player 1")
	p2 := truco.NewPlayer(1, "Player 2")
	p3 := truco.NewPlayer(2, "Player 3")
	p4 := truco.NewPlayer(3, "Player 4")
	log.Println("creating teams and adding players")
	t1 := truco.NewTeam(0, []*truco.Player{p1, p2})
	t2 := truco.NewTeam(1, []*truco.Player{p3, p4})

	log.Println("adding teams to games")
	if err := g.AddTeams([2]*truco.Team{t1, t2}); err != nil {
		panic(err)
	}

	log.Println("starting new round")
	_, err := g.NewRound()
	if err != nil {
		panic(err)
	}

	g.PrintState()
	if err := g.Play(0); err != nil {
		panic(err)
	}
	g.PrintState()
	if err := g.Play(0); err != nil {
		panic(err)
	}
	g.PrintState()
	if err := g.Play(0); err != nil {
		panic(err)
	}
	g.PrintState()
	if err := g.Play(0); err != nil {
		panic(err)
	}
	g.PrintState()
}
