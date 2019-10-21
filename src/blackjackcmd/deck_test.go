package main

import "testing"

func Test_NewDeck(t *testing.T) {
	if NewDeck() == nil {
		t.Fail()
	}
}

func Test_Deck_Draw(t *testing.T) {
	deck :=
		Deck([]Card{
			Card{
				Value: two,
				Suit:  suitHeart,
			},
			Card{
				Value: queen,
				Suit:  suitHeart,
			},
		})

	expected := Card{
		Value: two,
		Suit:  suitHeart,
	}

	if deck.Draw() != expected {
		t.Fail()
	}
}
