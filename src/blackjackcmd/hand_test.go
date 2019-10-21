package main

import (
	"testing"
)

func Test_Hand_FormatHand(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			Card{
				Value: queen,
				Suit:  suitHeart,
			}},
		Total: 10,
	}
	expected := "[ Q♡ ] "

	if hand.FormatHand() != expected {
		t.Fail()
	}
}

func Test_Hand_FormatDealerHand(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			Card{
				Value: queen,
				Suit:  suitHeart,
			},
			Card{
				Value: king,
				Suit:  suitDiamond,
			}},
		Total: 10,
	}
	expected := "[  ] [ K♢ ] "

	if hand.FormatDealerHand() != expected {
		t.Fail()
	}
}

func Test_Hand_CalculateHand(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			Card{
				Value: queen,
				Suit:  suitHeart,
			},
		},
		Total: 0,
	}
	expected := 10

	if hand.CalculateHand() != expected {
		t.Fail()
	}
}

func Test_Hand_CalculateHand_Aces(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			Card{
				Value: ace,
				Suit:  suitHeart,
			},
			Card{
				Value: ace,
				Suit:  suitDiamond,
			},
		},
		Total: 1,
	}
	expected := 2

	if hand.CalculateHand() != expected {
		t.Fail()
	}
}

func Test_Hand_IsBusted_False(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			Card{
				Value: ten,
				Suit:  suitHeart,
			},
		},
		Total: 21,
	}
	expected := false

	if hand.IsBusted() != expected {
		t.Fail()
	}
}

func Test_Hand_IsBusted_True(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			Card{
				Value: ten,
				Suit:  suitHeart,
			},
		},
		Total: 22,
	}
	expected := true

	if hand.IsBusted() != expected {
		t.Fail()
	}
}
