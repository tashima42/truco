package truco

import "testing"

func TestNewGame(t *testing.T) {
	g, err := NewGame()
	if err != nil {
		t.Error(err)
	}
	if g.id == "" {
		t.Error("empty id")
	}
	if g.maxPlayers != 2 {
		t.Error("game max players is greater than 2")
	}
	if len(g.players) != 0 {
		t.Error("game has more than 0 players")
	}
}

func TestNewPlayer(t *testing.T) {
	_, err := NewPlayer("a")
	if err != nil {
		if err != ErrNameTooShort {
			t.Error("expected error ErrNameTooShort, instead got: " + err.Error())
		}
	}
	_, err = NewPlayer("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		if err != ErrNameTooLong {
			t.Error("expected error ErrNameTooLong, instead got: " + err.Error())
		}
	}
}

func TestAddPlayer(t *testing.T) {
	g, err := NewGame()
	if err != nil {
		t.Error("failed to create game: " + err.Error())
	}
	p1, err := NewPlayer("player 1")
	if err != nil {
		t.Error("failed to create player: " + err.Error())
	}
	p2, err := NewPlayer("player 2")
	if err != nil {
		t.Error("failed to create player: " + err.Error())
	}

	if err := g.AddPlayer(p1); err != nil {
		t.Error("failed to add player 1 to game: " + err.Error())
	}

	if err := g.AddPlayer(p1); err != nil {
		if err != ErrPlayerAlreadyInGame {
			t.Error("expected error ErrPlayerAlreadyInGame, instead got: " + err.Error())
		}
	}

	if err := g.AddPlayer(p2); err != nil {
		t.Error("failed to add player 2 to game: " + err.Error())
	}

	if err := g.AddPlayer(p2); err != nil {
		if err != ErrGameFull {
			t.Error("expected error ErrGameFull, instead got: " + err.Error())
		}
	}
}

func TestRemovePlayer(t *testing.T) {
	g, err := NewGame()
	if err != nil {
		t.Error("failed to create game: " + err.Error())
	}
	p1, err := NewPlayer("player 1")
	if err != nil {
		t.Error("failed to create player: " + err.Error())
	}
	if err := g.AddPlayer(p1); err != nil {
		t.Error("failed to add player 1 to game: " + err.Error())
	}
	if err := g.RemovePlayer(p1); err != nil {
		t.Error(err.Error())
	}
	p2, err := NewPlayer("player 2")
	if err != nil {
		t.Error("failed to create player: " + err.Error())
	}
	if err := g.RemovePlayer(p2); err != nil {
		if err != ErrPlayerNotFound {
			t.Error("expected error ErrPlayerNotFound, instead got: " + err.Error())
		}
	}
}

func TestSetManilha(t *testing.T) {
	g, err := defaultGame(true)
	if err != nil {
		t.Error("failed to create game: " + err.Error())
	}

	if err := g.Start(); err != nil {
		t.Error("failed to start game: " + err.Error())
	}

	if g.manilha != string(g.deck[0]) {
		t.Error("manilha should be the same as the first card of the deck")
	}

	if g.manilha != string(ThreeHearts) {
		t.Error("seed isn't working properly, expected manilha to be B3, instead got: " + g.manilha)
	}

	// A2 A4 A7 CB
	if g.deckWeights[ThreeClubs] != 20 {
		t.Errorf("expected three clubs weight to be 20, instead got: %d", g.deckWeights[ThreeClubs])
	}
	if g.deckWeights[ThreeDiamonds] != 21 {
		t.Errorf("expected three diamonds weight to be 21, instead got: %d", g.deckWeights[ThreeDiamonds])
	}
	if g.deckWeights[ThreeHearts] != 22 {
		t.Errorf("expected three hearts weight to be 22, instead got: %d", g.deckWeights[ThreeHearts])
	}
	if g.deckWeights[ThreeSpades] != 23 {
		t.Errorf("expected three spades weight to be 23, instead got: %d", g.deckWeights[ThreeSpades])
	}
}

func TestDrawCards(t *testing.T) {
	g, err := defaultGame(true)
	if err != nil {
		t.Error("failed to create game: " + err.Error())
	}
	if err := g.Start(); err != nil {
		t.Error("failed to start game: " + err.Error())
	}
	if g.cardPointer != 7 {
		t.Errorf("card pointer is at wrong location, expected 7, instead got: %d", g.cardPointer)
	}

	p1 := g.players[0]
	p2 := g.players[1]
	if p1.cards[0] != QueenSpades {
		t.Error("wrong card for player, expected AC, instead got: " + p1.cards[0])
	}
	if p1.cards[1] != QueenHearts {
		t.Error("wrong card for player, expected BC, instead got: " + p1.cards[1])
	}
	if p1.cards[2] != ThreeDiamonds {
		t.Error("wrong card for player, expected C3, instead got: " + p1.cards[2])
	}
	if p2.cards[0] != SevenClubs {
		t.Error("wrong card for player, expected D7, instead got: " + p2.cards[0])
	}
	if p2.cards[1] != AceSpades {
		t.Error("wrong card for player, expected A1, instead got: " + p2.cards[1])
	}
	if p2.cards[2] != ThreeClubs {
		t.Error("wrong card for player, expected D3, instead got: " + p2.cards[2])
	}
}

func TestHasCard(t *testing.T) {
	p1, err := NewPlayer("player 1")
	if err != nil {
		t.Error("failed to create player 1")
	}
	p1.cards = []Card{QueenSpades, QueenHearts, ThreeDiamonds}

	if !p1.hasCard(QueenSpades) {
		t.Error("player 1 should have queen spades")
	}
	if p1.hasCard(AceClubs) {
		t.Error("player 1 should not have ace clubs")
	}
}

func TestPlay(t *testing.T) {
	g, err := defaultGame(true)
	if err != nil {
		t.Error("failed to create game: " + err.Error())
	}
	if err := g.Start(); err != nil {
		t.Error("failed to start game: " + err.Error())
	}
	p1 := g.players[0]
	if err := g.Play(p1, QueenSpades); err != nil {
		t.Error("failed to play card: " + err.Error())
	}
	if err := g.Play(p1, QueenSpades); err != nil {
		if err != ErrNotPlayerTurn {
			t.Error("error should have been not player turn, instead got: " + err.Error())
		}
	}
}

func TestPlayCard(t *testing.T) {
	g, err := defaultGame(true)
	if err != nil {
		t.Error("failed to create game: " + err.Error())
	}
	if err := g.Start(); err != nil {
		t.Error("failed to start game: " + err.Error())
	}
	p1 := g.players[0]
	g.playCard(p1, ThreeDiamonds)
	if len(p1.cards) != 2 {
		t.Error("player 1 should have 2 cards")
	}
	if len(g.pile) != 1 {
		t.Error("pile should have 1 card")
	}
	if g.pile[0][0] != ThreeDiamonds {
		t.Error("wrong card in pile")
	}
}

func defaultGame(seed bool) (*Game, error) {
	g, err := NewGame()
	if err != nil {
		return nil, err
	}
	p1, err := NewPlayer("player 1")
	if err != nil {
		return nil, err
	}
	p2, err := NewPlayer("player 2")
	if err != nil {
		return nil, err
	}

	if err := g.AddPlayer(p1); err != nil {
		return nil, err
	}

	if err := g.AddPlayer(p2); err != nil {
		return nil, err
	}
	if seed {
		g.Seed(123, 456)
	}
	return g, nil
}
