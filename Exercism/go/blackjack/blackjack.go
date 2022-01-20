package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValues := map[string]int{"ace": 11, "two": 2, "three": 3,
		"four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10,
		"jack": 10, "queen": 10, "king": 10, "other": 0}
	return cardValues[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	score1 := ParseCard(card1)
	score2 := ParseCard(card2)
	if score1+score2 == 21 {
		return true
	} else {
		return false
	}
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if !isBlackjack {
		return "P"
	} else {
		if isBlackjack && dealerScore >= 10 {
			return "S"
		} else {
			return "W"
		}
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore >= 17 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else {
		if dealerScore >= 7 {
			return "H"
		} else {
			return "S"
		}
	}
}
