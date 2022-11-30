package repo

import (
	domain "Goscore/Domain"
)

var UserScores = make(map[int]int)

func SaveScore(score domain.UserScore) {
	UserScores[score.User] = score.Score
}

func GetScores() []domain.UserScore {
	var userScoreList []domain.UserScore
	for user, score := range UserScores {
		userScoreList = append(userScoreList, domain.UserScore{User: user, Score: score})
	}

	return userScoreList
}
