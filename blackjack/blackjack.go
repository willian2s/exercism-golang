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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case (value1+value2 == 21) && (dealerValue != 10 && dealerValue != 11):
		return "W"
	case value1+value2 >= 17 && value1+value2 <= 20:
		return "S"
	case value1+value2 >= 12 && value1+value2 <= 16:
		switch {
		case dealerValue >= 7:
			return "H"
		default:
			return "S"
		}
	case value1+value2 == 11:
		return "H"
	case value1+value2 <= 10:
		return "H"
	default:
		return "S"
	}
}
