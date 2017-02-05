package deck

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

// Suit will be Clubs, Diamonds, Hearts or Spades
type Suit struct{ Name, Symbol string }

// Facevalue will be 2 to 10, Jack Queen King ro Ace
type Facevalue struct {
	Name  string
	Value int
}

// Card is a Suit and Facevalue
type Card struct {
	Suit  Suit
	Value Facevalue
}

// GreaterThan tells if one card's Value is GreaterThan the other
func (card *Card) GreaterThan(b *Card) bool {
	return card.Value.Value > b.Value.Value
}

// LessThan tells if one card's Value is LessThan the other
func (card *Card) LessThan(b *Card) bool {
	return card.Value.Value < b.Value.Value
}

// Equal returns true or false based on the Facevalue
func (card *Card) Equal(b *Card) bool {
	return card.Value.Value == b.Value.Value
}

// Facecard returns true for J,Q,K,A, false for all others
func (card *Card) Facecard() (ans bool) {
	n := card.Value.Name
	return n == "Jack" || n == "Queen" || n == "King" || n == "Ace"
}

// ToStr returns a pretty string for the Cards
func (card *Card) ToStr() string {
	if card.Facecard() {
		return fmt.Sprintf(" %c%s", card.Value.Name[0], card.Suit.Symbol)
	}
	return fmt.Sprintf("%2d%s", card.Value.Value, card.Suit.Symbol)
}

// Deck is a collection of up to 52 cards
type Deck struct {
	Cards []Card
}

// Initialize Initializes a deck of 52 Cards
func (d *Deck) Initialize() error {
	// TODO: Empty Deck before initialization
	// Ensure cards past 52 are deleted.

	suits := []Suit{
		Suit{"Clubs", "♣"},
		Suit{"Diamonds", "♦"},
		Suit{"Hearts", "♥"},
		Suit{"Spades", "♠"},
	}
	facevalues := []Facevalue{
		Facevalue{"Two", 2},
		Facevalue{"Three", 3},
		Facevalue{"Four", 4},
		Facevalue{"Five", 5},
		Facevalue{"Six", 6},
		Facevalue{"Seven", 7},
		Facevalue{"Eight", 8},
		Facevalue{"Nine", 9},
		Facevalue{"Ten", 10},
		Facevalue{"Jack", 11},
		Facevalue{"Queen", 12},
		Facevalue{"King", 13},
		Facevalue{"Ace", 14},
	}

	for _, suit := range suits {
		for _, facevalue := range facevalues {
			d.Cards = append(d.Cards, Card{Suit: suit, Value: facevalue})
		}
	}
	d.Shuffle()
	return nil
}

// Shuffle randomizes a Deck using crypto/rand
func (d *Deck) Shuffle() (err error) {
	var old []Card
	old = d.Cards
	var shuffled []Card
	for i := len(old); i > 0; i-- {
		nBig, e := rand.Int(rand.Reader, big.NewInt(int64(i)))
		if e != nil {
			panic(e)
		}
		j := nBig.Int64()
		shuffled = append(shuffled, old[j])
		old = remove(old, j)
	}
	d.Cards = shuffled
	return nil
}

//remove removes a card at index i from a slice of Cards
func remove(slice []Card, i int64) []Card {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// Draw will remove cards from the beginning of the deck (index 0)
// and return a slice of []Card type.
func (d *Deck) Draw(count int) (cards []Card, err error) {
	if count > len(d.Cards) {
		return nil, errors.New("Not enough cards left in the deck")
	}

	hand := d.Cards[0:count]
	d.Cards = d.Cards[count:]
	return hand, nil
}

// returns the number of cards left in the deck
func (d *Deck) CardsLeft() int {
	return len(d.Cards)
}
