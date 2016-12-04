package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Suit struct{ Name, Symbol string }

type Facevalue struct {
	Name  string
	Value int
}

type Card struct {
	Suit  Suit
	Value Facevalue
}

func (card_a *Card) greater_than(card_b *Card) bool {
	return card_a.Value.Value > card_b.Value.Value
}

func (card_a *Card) less_than(card_b *Card) bool {
	return card_a.Value.Value < card_b.Value.Value
}

func (card_a *Card) equal(card_b *Card) bool {
	return card_a.Value.Value == card_b.Value.Value
}

func (card *Card) facecard() (ans bool) {
	n := card.Value.Name
	return n == "Jack" || n == "Queen" || n == "King" || n == "Ace"
}

type Deck struct {
	Cards []Card
}

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

func remove(slice []Card, i int64) []Card {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func (d *Deck) Draw(count int) (cards []Card, err error) {
	// TODO: Implement this function
	return nil, nil
}

func main() {
	var d Deck
	d.Initialize()
	for i := 5; i > 0; i-- {
		fmt.Printf("First Card: %+v\n", d.Cards[0])
		fmt.Printf("length of deck: %d\n", len(d.Cards))
		fmt.Printf("Shuffing...\n")
		e := d.Shuffle()
		if e != nil {
			panic(e)
		}
	}
	return
}
