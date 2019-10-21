package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("\n\n")
	fmt.Println("Welcome to Blackjack.  Would you like to play a hand? (y)es or (n)o")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		playNew(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func setup() (Hand, Hand, Deck) {
	yourHand := Hand{}
	dealerHand := Hand{}
	deck := NewDeck()
	deck.Shuffle()

	yourHand.Cards = append(yourHand.Cards, deck.Draw())
	yourHand.Cards = append(yourHand.Cards, deck.Draw())
	dealerHand.Cards = append(dealerHand.Cards, deck.Draw())
	dealerHand.Cards = append(dealerHand.Cards, deck.Draw())

	fmt.Printf("\n\n")
	yourHand.Total = yourHand.CalculateHand()
	dealerHand.Total = dealerHand.CalculateHand()
	yourHand.DisplayHand("Player", true)
	dealerHand.DisplayHand("Dealer", false)
	
	if isBlackjack(yourHand.Total) {
		hitDealer(yourHand, dealerHand, deck)
	}
	
	if isBlackjack(dealerHand.Total) {
		fmt.Printf("\n\n")
		yourHand.DisplayHand("Player", true)
		dealerHand.DisplayHand("Dealer", true)
		compareHand(yourHand, dealerHand)
	}

	return yourHand, dealerHand, deck
}

func playNew(bs string) {
	if bs == "y" {
		yourHand, dealerHand, deck := setup()
		fmt.Println("Would you like (h)it or (s)tand?")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			hitOrStand(scanner.Text(), yourHand, dealerHand, deck)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Ahhh... that's too bad. I would have beaten you anyway.\n\n\n")
		os.Exit(1)
	}

}

func hitOrStand(bs string, yourHand Hand, dealerHand Hand, deck Deck) {
	if bs == "h" {
		fmt.Printf("\n\n")
		yourHand.Cards = append(yourHand.Cards, deck.Draw())
		yourHand.Total = yourHand.CalculateHand()
		yourHand.DisplayHand("Player", true)
		if yourHand.IsBusted() {
			fmt.Printf("You Busted! You Lose!\n\n\n")
			PlayAgain()
		}

		dealerHand.DisplayHand("Dealer", false)
		if dealerHand.IsBusted() {
			fmt.Printf("Dealer Busted! You Win!\n\n\n")
			PlayAgain()
		}

		if isBlackjack(yourHand.Total) {
			hitDealer(yourHand, dealerHand, deck)
		}

		fmt.Println("Would you like (h)it or (s)tand?")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			hitOrStand(scanner.Text(), yourHand, dealerHand, deck)
		}
	} else if bs == "s" {
		hitDealer(yourHand, dealerHand, deck)
	} else {
		fmt.Println("Would you like (h)it or (s)tand?")
	}
}

func hitDealer(yourHand Hand, dealerHand Hand, deck Deck) {
	fmt.Printf("\n\n")

	if dealerHand.Total < 17 || yourHand.Total > dealerHand.Total {
		dealerHand.Cards = append(dealerHand.Cards, deck.Draw())
		dealerHand.Total = dealerHand.CalculateHand()
		yourHand.DisplayHand("Player", true)
		dealerHand.DisplayHand("Dealer", true)
		if dealerHand.IsBusted() {
			fmt.Printf("Dealer Busted! You Win!\n\n\n")
			PlayAgain()
		}
		hitDealer(yourHand, dealerHand, deck)
	} else {
		yourHand.DisplayHand("Player", true)
		dealerHand.DisplayHand("Dealer", true)
		compareHand(yourHand, dealerHand)
	}
}

func compareHand(yourHand Hand, dealerHand Hand) {
	if yourHand.Total > dealerHand.Total {
		if isBlackjack(yourHand.Total) {
			fmt.Printf("Blackjack! You Win!\n\n")
		} else {
			fmt.Printf("You Win!\n\n")
		}
	} else {
		if isBlackjack(dealerHand.Total) {
			fmt.Printf("Blackjack! You Lose!\n\n")
		} else {
			fmt.Printf("You Lose!\n\n")
		}
	}
	PlayAgain()
}

func isBlackjack(total int) bool {
	if total == 21 {
		return true
	}

	return false
}

func PlayAgain() {
	fmt.Println("Play Again (y)es or (n)o?")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		playNew(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
