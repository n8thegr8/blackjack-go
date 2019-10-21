package main

import (
	"math/rand"
	"time"
)

type Deck []Card

func NewDeck() Deck {
	deck := Deck{}
	faceValues := []FaceValue{ace, two, three, four, five, six, seven, eight, nine, ten, jack, queen, king}
	suits := []Suit{suitHeart, suitDiamond, suitSpade, suitClub}

	for i := 0; i < len(faceValues); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Value: faceValues[i],
				Suit:  suits[n],
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d *Deck) Draw() Card {
	drawnCard := (*d)[0]
	*d = (*d)[1:]
	return drawnCard
}
