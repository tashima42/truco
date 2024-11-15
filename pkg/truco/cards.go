package truco

type CardValue int
type CardSuit int

// The type Card holds all information needed for the cards in a game of Truco.
// All cards have a value and a suit. Values and suits are unsigned integers,
// for both values, the higher, the stronger the card is.
type Card struct {
	ID    int
	Value CardValue
	Suit  CardSuit
}

const (
	CardSuitDiamonds CardSuit = 0
	CardSuitSpades   CardSuit = 1
	CardSuitHearts   CardSuit = 2
	CardSuitClubs    CardSuit = 3
)

const (
	CardValue4 CardValue = 0
	CardValue5 CardValue = 1
	CardValue6 CardValue = 2
	CardValue7 CardValue = 3
	CardValueQ CardValue = 4
	CardValueJ CardValue = 5
	CardValueK CardValue = 6
	CardValueA CardValue = 7
	CardValue2 CardValue = 8
	CardValue3 CardValue = 9
)

var Cards = map[int]Card{
	0: {ID: 0, Value: CardValueA, Suit: CardSuitClubs},
	1: {ID: 1, Value: CardValueA, Suit: CardSuitDiamonds},
	2: {ID: 2, Value: CardValueA, Suit: CardSuitHearts},
	3: {ID: 3, Value: CardValueA, Suit: CardSuitSpades},

	4: {ID: 4, Value: CardValue2, Suit: CardSuitClubs},
	5: {ID: 5, Value: CardValue2, Suit: CardSuitDiamonds},
	6: {ID: 6, Value: CardValue2, Suit: CardSuitHearts},
	7: {ID: 7, Value: CardValue2, Suit: CardSuitSpades},

	8:  {ID: 8, Value: CardValue3, Suit: CardSuitClubs},
	9:  {ID: 9, Value: CardValue3, Suit: CardSuitDiamonds},
	10: {ID: 10, Value: CardValue3, Suit: CardSuitHearts},
	11: {ID: 11, Value: CardValue3, Suit: CardSuitSpades},

	12: {ID: 12, Value: CardValue4, Suit: CardSuitClubs},
	13: {ID: 13, Value: CardValue4, Suit: CardSuitDiamonds},
	14: {ID: 14, Value: CardValue4, Suit: CardSuitHearts},
	15: {ID: 15, Value: CardValue4, Suit: CardSuitSpades},

	16: {ID: 16, Value: CardValue5, Suit: CardSuitClubs},
	17: {ID: 17, Value: CardValue5, Suit: CardSuitDiamonds},
	18: {ID: 18, Value: CardValue5, Suit: CardSuitHearts},
	19: {ID: 19, Value: CardValue5, Suit: CardSuitSpades},

	20: {ID: 20, Value: CardValue6, Suit: CardSuitClubs},
	21: {ID: 21, Value: CardValue6, Suit: CardSuitDiamonds},
	22: {ID: 22, Value: CardValue6, Suit: CardSuitHearts},
	23: {ID: 23, Value: CardValue6, Suit: CardSuitSpades},

	24: {ID: 24, Value: CardValue7, Suit: CardSuitClubs},
	25: {ID: 25, Value: CardValue7, Suit: CardSuitDiamonds},
	26: {ID: 26, Value: CardValue7, Suit: CardSuitHearts},
	27: {ID: 27, Value: CardValue7, Suit: CardSuitSpades},

	28: {ID: 28, Value: CardValueK, Suit: CardSuitClubs},
	29: {ID: 29, Value: CardValueK, Suit: CardSuitDiamonds},
	30: {ID: 30, Value: CardValueK, Suit: CardSuitHearts},
	31: {ID: 31, Value: CardValueK, Suit: CardSuitSpades},

	32: {ID: 32, Value: CardValueQ, Suit: CardSuitClubs},
	33: {ID: 33, Value: CardValueQ, Suit: CardSuitDiamonds},
	34: {ID: 34, Value: CardValueQ, Suit: CardSuitHearts},
	35: {ID: 35, Value: CardValueQ, Suit: CardSuitSpades},

	36: {ID: 36, Value: CardValueJ, Suit: CardSuitClubs},
	37: {ID: 37, Value: CardValueJ, Suit: CardSuitDiamonds},
	38: {ID: 38, Value: CardValueJ, Suit: CardSuitHearts},
	39: {ID: 39, Value: CardValueJ, Suit: CardSuitSpades},
}
