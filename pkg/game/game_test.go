package game

import "testing"

func TestNewGame(t *testing.T) {
	g, err := NewGame(0)
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
	g, err := NewGame(0)
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
	g, err := NewGame(0)
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
	g, err := NewGame(1352522007503861735)
	if err != nil {
		t.Error("failed to create game: " + err.Error())
	}
	g.setManilha()

	if g.manilha != string(g.deck[0]) {
		t.Error("manilha should be the same as the first card of the deck")
	}

	if g.manilha != string(JackHearts) {
		t.Error("seed isn't working properly, expected manilha to be jack hearts, instead got: " + g.manilha)
	}

	if g.deckWeights[JackClubs] != 16 {
		t.Errorf("expected jack clubs weight to be 16, instead got: %d", g.deckWeights[JackClubs])
	}
	if g.deckWeights[JackDiamonds] != 17 {
		t.Errorf("expected jack diamonds weight to be 16, instead got: %d", g.deckWeights[JackDiamonds])
	}
	if g.deckWeights[JackHearts] != 18 {
		t.Errorf("expected jack hearts weight to be 16, instead got: %d", g.deckWeights[JackHearts])
	}
	if g.deckWeights[JackSpades] != 19 {
		t.Errorf("expected jack spades weight to be 16, instead got: %d", g.deckWeights[JackSpades])
	}
}
