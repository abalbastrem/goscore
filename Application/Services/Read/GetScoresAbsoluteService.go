package services_read

import (
	requests_read "Goscore/Application/Requests/Read"
	utils "Goscore/Application/Utils"
	domain "Goscore/Domain"
	repo "Goscore/Infrastructure/Repos"
)

type GetScoresAbsoluteService struct {
	Request requests_read.GetScoresAbsoluteRequest
}

func ExecAbsolute(s GetScoresAbsoluteService) map[int]domain.UserScore {
	userScoreList := repo.GetScoreList()
	utils.Sort(userScoreList)
	userScoreMap := make(map[int]domain.UserScore)
	for index, userScore := range userScoreList {
		userScoreMap[index+1] = userScore
		if len(userScoreMap) == s.Request.Rank {
			break
		}
	}

	return userScoreMap
}
