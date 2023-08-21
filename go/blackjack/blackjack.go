package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cardMap = map[string]int{
		"two" : 2, "three" : 3, "four" : 4, "five" : 5,
		"six" : 6, "seven" : 7, "eight" : 8, "nine" : 9,
		"ten" : 10, "jack" : 10, "queen" : 10, "king" : 10,
		"ace" : 11}
	return cardMap[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	val1 := ParseCard(card1)
	val2 := ParseCard(card2)
	valD := ParseCard(dealerCard)
	sum := val1 + val2
	switch {
		case card1 == "ace" && card2 == "ace":
			return "P"
		case sum == 21:
			if (valD >= 10) { return "S" }
			return "W"
		case sum <= 20 && sum >= 17:
			return "S"
		case sum <= 16 && sum >= 12:
			if (valD >= 7) { return "H" }
			return "S"
		default:
			return "H"
	}
}
