package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "other":
		return 0
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var cardSum int = ParseCard(card1) + ParseCard(card2)
	var move = "S"
	if ParseCard(card1) == 11 && ParseCard(card2) == 11 {
		move = "P"
	} else if cardSum == 21 {
		if ParseCard(dealerCard) != 10 && ParseCard(dealerCard) != 11 {
			move = "W"
		} else {
			move = "S"
		}
	} else if cardSum >= 17 && cardSum <= 20 {
		move = "S"
	} else if cardSum >= 12 && cardSum <= 16 {
		if ParseCard(dealerCard) >= 7 {
			move = "H"
		} else {
			move = "S"
		}
	} else if cardSum <= 11 {
		move = "H"
	}

	return move
}
