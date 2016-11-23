package goplaycards

type Suit struct{ name, symbol string }

type Suits struct {
	Club    Suit
	Diamond Suit
	Heart   Suit
	Spade   Suit
}

type FaceValue struct {
	Name  string
	Value int
}

type Card interface {
	greater_than(*Card) (ans bool)
	less_than(*Card) (ans bool)
	equal(*Card) (ans bool)
	facecard(*Card) (ans bool)
}

type Card struct {
	Suit  Suit
	Value Value
}

type Deck interface {
	(d Deck) Shuffle(count int) (err error)
	(d Deck) Draw(count int) (cards []Card, err error)
	(d Deck) Initialize(*Deck ) (err error)
}

type Deck struct {
	Cards []Card
}

func Initialize(*Deck) error {
	suits := []Suit{
		Suit{"Clubs", "♣"},
		Suit{"Diamonds", "♦"},
		Suit{"Hearts", "♥"},
		Suit{"Spades", "♠"},
	}
	facevalues := []Facevalue{
		FaceValue{"Two", 2},
		FaceValue{"Three", 3},
		FaceValue{"Four", 4},
		FaceValue{"Five", 5},
		FaceValue{"Six", 6},
		FaceValue{"Seven", 7},
		FaceValue{"Eight", 8},
		FaceValue{"Nine", 9},
		FaceValue{"Ten", 10},
		FaceValue{"Jack", 11},
		FaceValue{"Queen", 12},
		FaceValue{"King", 13},
		FaceValue{"Ace", 14},
	}

	for _, suit := range suits { 
		for _, facevalue := range  facevalues{
			
		}
	}

}
