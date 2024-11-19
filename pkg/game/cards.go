package game

import "math/rand/v2"

type (
	Card        string
	DeckWeights map[Card]int
)

var (
	AceSpades     Card = "A1"
	AceHearts     Card = "B1"
	AceDiamonds   Card = "C1"
	AceClubs      Card = "D1"
	TwoSpades     Card = "A2"
	TwoHearts     Card = "B2"
	TwoDiamonds   Card = "C2"
	TwoClubs      Card = "D2"
	ThreeSpades   Card = "A3"
	ThreeHearts   Card = "B3"
	ThreeDiamonds Card = "C3"
	ThreeClubs    Card = "D3"
	FourSpades    Card = "A4"
	FourHearts    Card = "B4"
	FourDiamonds  Card = "C4"
	FourClubs     Card = "D4"
	FiveSpades    Card = "A5"
	FiveHearts    Card = "B5"
	FiveDiamonds  Card = "C5"
	FiveClubs     Card = "D5"
	SixSpades     Card = "A6"
	SixHearts     Card = "B6"
	SixDiamonds   Card = "C6"
	SixClubs      Card = "D6"
	SevenSpades   Card = "A7"
	SevenHearts   Card = "B7"
	SevenDiamonds Card = "C7"
	SevenClubs    Card = "D7"
	JackSpades    Card = "AB"
	JackHearts    Card = "BB"
	JackDiamonds  Card = "CB"
	JackClubs     Card = "DB"
	QueenSpades   Card = "AC"
	QueenHearts   Card = "BC"
	QueenDiamonds Card = "CC"
	QueenClubs    Card = "DC"
	KingSpades    Card = "AE"
	KingHearts    Card = "BE"
	KingDiamonds  Card = "CE"
	KingClubs     Card = "DE"
)

func DefaultDeck() []Card {
	return []Card{
		AceSpades,
		AceHearts,
		AceDiamonds,
		AceClubs,
		TwoSpades,
		TwoHearts,
		TwoDiamonds,
		TwoClubs,
		ThreeSpades,
		ThreeHearts,
		ThreeDiamonds,
		ThreeClubs,
		FourSpades,
		FourHearts,
		FourDiamonds,
		FourClubs,
		FiveSpades,
		FiveHearts,
		FiveDiamonds,
		FiveClubs,
		SixSpades,
		SixHearts,
		SixDiamonds,
		SixClubs,
		SevenSpades,
		SevenHearts,
		SevenDiamonds,
		SevenClubs,
		JackSpades,
		JackHearts,
		JackDiamonds,
		JackClubs,
		QueenSpades,
		QueenHearts,
		QueenDiamonds,
		QueenClubs,
		KingSpades,
		KingHearts,
		KingDiamonds,
		KingClubs,
	}
}

func DefaultDeckWeights() map[Card]int {
	return map[Card]int{
		FourSpades:    1,
		FourHearts:    1,
		FourDiamonds:  1,
		FourClubs:     1,
		FiveSpades:    2,
		FiveHearts:    2,
		FiveDiamonds:  2,
		FiveClubs:     2,
		SixSpades:     3,
		SixHearts:     3,
		SixDiamonds:   3,
		SixClubs:      3,
		SevenSpades:   4,
		SevenHearts:   4,
		SevenDiamonds: 4,
		SevenClubs:    5,
		QueenSpades:   5,
		QueenHearts:   5,
		QueenDiamonds: 5,
		QueenClubs:    6,
		JackSpades:    6,
		JackHearts:    6,
		JackDiamonds:  6,
		JackClubs:     6,
		KingSpades:    7,
		KingHearts:    7,
		KingDiamonds:  7,
		KingClubs:     7,
		AceSpades:     8,
		AceHearts:     8,
		AceDiamonds:   8,
		AceClubs:      8,
		TwoSpades:     9,
		TwoHearts:     9,
		TwoDiamonds:   9,
		TwoClubs:      9,
		ThreeSpades:   10,
		ThreeHearts:   10,
		ThreeDiamonds: 10,
		ThreeClubs:    10,
	}
}

func ShuffledDeck(seed1, seed2 uint64) []Card {
	if seed1 == 0 || seed2 == 0 {
		seed1 = rand.Uint64()
		seed2 = rand.Uint64()
	}
	r := rand.New(rand.NewPCG(seed1, seed2))

	deck := DefaultDeck()
	r.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}
