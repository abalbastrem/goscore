package repo

import (
	domain "Goscore/Domain"
)

var UserScores = make(map[int]int)

func SaveScore(score domain.UserScore) {
	UserScores[score.User] = score.Score
}
