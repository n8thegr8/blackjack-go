package main

import "testing"

func Test_Suit_String(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    Suit
		expected string
	}{
		{"heart", suitHeart, "♡"},
		{"diamond", suitDiamond, "♢"},
		{"spade", suitSpade, "♤"},
		{"club", suitClub, "♧"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input.String() != tc.expected {
				t.Fail()
			}
		})
	}
}

func Test_FaceValue_BlackJackValue(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    FaceValue
		expected int
	}{
		{"Ace", ace, 11},
		{"Two", two, 2},
		{"Three", three, 3},
		{"Four", four, 4},
		{"Five", five, 5},
		{"Six", six, 6},
		{"Seven", seven, 7},
		{"Eight", eight, 8},
		{"Nine", nine, 9},
		{"Ten", ten, 10},
		{"Jack", jack, 10},
		{"Queen", queen, 10},
		{"King", king, 10},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input.BlackJackValue() != tc.expected {
				t.Fail()
			}
		})
	}
}

func Test_FaceValue_String(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    FaceValue
		expected string
	}{
		{"Ace", ace, "A"},
		{"Two", two, "2"},
		{"Three", three, "3"},
		{"Four", four, "4"},
		{"Five", five, "5"},
		{"Six", six, "6"},
		{"Seven", seven, "7"},
		{"Eight", eight, "8"},
		{"Nine", nine, "9"},
		{"Ten", ten, "10"},
		{"Jack", jack, "J"},
		{"Queen", queen, "Q"},
		{"King", king, "K"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input.String() != tc.expected {
				t.Fail()
			}
		})
	}
}
