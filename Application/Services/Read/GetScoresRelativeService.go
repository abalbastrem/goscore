package services_read

import (
	requests_read "Goscore/Application/Requests/Read"
	utils "Goscore/Application/Utils"
	domain "Goscore/Domain"
	repo "Goscore/Infrastructure/Repos"
)

type GetScoresRelativeService struct {
	Request requests_read.GetScoresRelativeRequest
}

func ExecRelative(s GetScoresRelativeService) map[int]domain.UserScore {
	userScoreList := repo.GetScores()
	utils.Sort(userScoreList)
	userScoreMap := make(map[int]domain.UserScore)
	for index, userScore := range userScoreList {
		if index >= s.Request.Rank-s.Request.NRelatives &&
			index <= s.Request.Rank+s.Request.NRelatives {
			userScoreMap[index] = userScore
		}
	}

	return userScoreMap
}
