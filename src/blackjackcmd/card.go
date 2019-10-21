package main

import (
	"fmt"
	"strconv"
)

type Suit int

const (
	suitHeart Suit = iota
	suitDiamond
	suitSpade
	suitClub
)

func (s Suit) String() string {
	switch s {
	case suitHeart:
		return "♡"

	case suitDiamond:
		return "♢"

	case suitSpade:
		return "♤"

	case suitClub:
		return "♧"
	}
	panic(fmt.Sprintf("invalid suit value: %d", s))
}

type FaceValue int

const (
	ace FaceValue = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
)

func (fv FaceValue) BlackJackValue() int {
	switch fv {
	case ace:
		return 11

	case jack:
		return 10

	case queen:
		return 10

	case king:
		return 10
	}
	return int(fv)
}

func (fv FaceValue) String() string {
	switch fv {
	case ace:
		return "A"

	case jack:
		return "J"

	case queen:
		return "Q"

	case king:
		return "K"
	}
	return strconv.Itoa(int(fv))
}

type Card struct {
	Value FaceValue
	Suit  Suit
}