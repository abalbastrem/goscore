package services_write

import (
	requests_write "Goscore/Application/Requests/Write"
	domain "Goscore/Domain"
	repo "Goscore/Infrastructure/Repos"
	"strconv"
)

type SaveScoreDiffService struct {
	Request requests_write.SaveScoreDiffRequest
}

func ExecDiff(s SaveScoreDiffService) {
	scores := repo.GetRawScores()
	currScore, found := scores[s.Request.User]
	if !found {
		currScore = 0
	}
	opStr := s.Request.Score[:1]
	noStr := s.Request.Score[1:]
	no, _ := strconv.Atoi(noStr)
	if opStr == "-" {
		no = no * (-1)
	}
	newScore := currScore + no
	repo.InsertScore(domain.UserScore{User: s.Request.User, Score: newScore})
}
