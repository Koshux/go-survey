package fuzzylogic

type FuzzyScore struct {
	Low    float64
	Medium float64
	High   float64
}

// Returns the membership of the users quiz score to the Low, Medium, and High
// categories.
func FuzzyScoreMembership(score int) FuzzyScore {
	return FuzzyScore{
		Low:    LowScoreMembership(score),
		Medium: MediumScoreMembership(score),
		High:   HighScoreMembership(score),
	}
}

func LowScoreMembership(score int) float64 {
	if score <= 3 {
		return 1 // Fully belongs to Low score
	} else if score > 3 && score < 5 {
		return float64(5-score) / 2 // Partially belongs to Low score
	} else {
		return 0 // Does not belong to Low score
	}
}

func MediumScoreMembership(score int) float64 {
	if score <= 3 || score >= 7 {
		return 0 // Does not belong to Medium score
	} else if score > 3 && score < 5 {
		return float64(score-3) / 2 // Partially belongs to Medium score
	} else if score >= 5 && score <= 7 {
		return 1 // Fully belongs to Medium score
	} else {
		return 0 // Does not belong to Medium score
	}
}

func HighScoreMembership(score int) float64 {
	if score <= 5 {
		return 0 // Does not belong to High score
	} else if score > 5 && score < 7 {
		return float64(score-5) / 2 // Partially belongs to High score
	} else {
		return 1 // Fully belongs to High score
	}
}
