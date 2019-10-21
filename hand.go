package main

import (
	"fmt"
	"strconv"
)

type Hand struct {
	Cards Deck
	Total int
}

func (h Hand) FormatHand() string {
	hand := ""

	for _, card := range h.Cards {
		hand += fmt.Sprintf("[ %s%s ] ", card.Value, card.Suit)
	}

	return hand
}

func (h Hand) FormatDealerHand() string {
	hand := ""

	for i, card := range h.Cards {
		if i == 0 {
			hand += "[  ] "
		} else {
			hand += fmt.Sprintf("[ %s%s ] ", card.Value, card.Suit)
		}
	}

	return hand
}

func (h Hand) DisplayHand(player string, showTotal bool) {
	text := fmt.Sprintf("%s Hand: %s", player, h.FormatDealerHand())

	if showTotal == true {
		text = fmt.Sprintf("%s Hand: %s = %s", player, h.FormatHand(), strconv.Itoa(h.Total))
	}

	fmt.Println(text)
}

func (h Hand) CalculateHand() int {
	total := 0

	for i := range h.Cards {
		total += h.Cards[i].Value.BlackJackValue()
	}

	if total > 21 {
		total = 0
		for i := range h.Cards {
			if h.Cards[i].Value.BlackJackValue() == 11 {
				total++
			} else {
				total += h.Cards[i].Value.BlackJackValue()
			}
		}
	}

	return total
}

func (h Hand) IsBusted() bool {
	if h.Total > 21 {
		return true
	}
	return false
}
