package truco

import (
	"math/rand/v2"
	"strconv"
)

type (
	// Card values are based on unicode playing cards U+1F0<card-value>
	Card        string
	DeckWeights map[Card]int
)

// Card
const (
	Ace   = "1"
	Two   = "2"
	Three = "3"
	Four  = "4"
	Five  = "5"
	Six   = "6"
	Seven = "7"
	Jack  = "B"
	Queen = "D"
	King  = "E"
)

// Suit
const (
	Spades   = "A"
	Hearts   = "B"
	Diamonds = "C"
	Clubs    = "D"
)

const (
	AceSpades     Card = Spades + Ace     // "A1"
	AceHearts     Card = Hearts + Ace     // "B1"
	AceDiamonds   Card = Diamonds + Ace   // "C1"
	AceClubs      Card = Clubs + Ace      // "D1"
	TwoSpades     Card = Spades + Two     // "A2"
	TwoHearts     Card = Hearts + Two     // "B2"
	TwoDiamonds   Card = Diamonds + Two   // "C2"
	TwoClubs      Card = Clubs + Two      // "D2"
	ThreeSpades   Card = Spades + Three   // "A3"
	ThreeHearts   Card = Hearts + Three   // "B3"
	ThreeDiamonds Card = Diamonds + Three // "C3"
	ThreeClubs    Card = Clubs + Three    // "D3"
	FourSpades    Card = Spades + Four    // "A4"
	FourHearts    Card = Hearts + Four    // "B4"
	FourDiamonds  Card = Diamonds + Four  // "C4"
	FourClubs     Card = Clubs + Four     // "D4"
	FiveSpades    Card = Spades + Five    // "A5"
	FiveHearts    Card = Hearts + Five    // "B5"
	FiveDiamonds  Card = Diamonds + Five  // "C5"
	FiveClubs     Card = Clubs + Five     // "D5"
	SixSpades     Card = Spades + Six     // "A6"
	SixHearts     Card = Hearts + Six     // "B6"
	SixDiamonds   Card = Diamonds + Six   // "C6"
	SixClubs      Card = Clubs + Six      // "D6"
	SevenSpades   Card = Spades + Seven   // "A7"
	SevenHearts   Card = Hearts + Seven   // "B7"
	SevenDiamonds Card = Diamonds + Seven // "C7"
	SevenClubs    Card = Clubs + Seven    // "D7"
	JackSpades    Card = Spades + Jack    // "AB"
	JackHearts    Card = Hearts + Jack    // "BB"
	JackDiamonds  Card = Diamonds + Jack  // "CB"
	JackClubs     Card = Clubs + Jack     // "DB"
	QueenSpades   Card = Spades + Queen   // "AD"
	QueenHearts   Card = Hearts + Queen   // "BD"
	QueenDiamonds Card = Diamonds + Queen // "CD"
	QueenClubs    Card = Clubs + Queen    // "DD"
	KingSpades    Card = Spades + King    // "AE"
	KingHearts    Card = Hearts + King    // "BE"
	KingDiamonds  Card = Diamonds + King  // "CE"
	KingClubs     Card = Clubs + King     // "DE"
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
		SevenClubs:    4,
		QueenSpades:   5,
		QueenHearts:   5,
		QueenDiamonds: 5,
		QueenClubs:    5,
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

func (c Card) Unicode() string {
	num, err := strconv.ParseInt("1F0"+string(c), 16, 32)
	if err != nil {
		panic(err)
	}
	return string(rune(num))
}
