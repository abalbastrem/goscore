package services_write

import (
	requests_write "Goscore/Application/Requests/Write"
	domain "Goscore/Domain"
	repo "Goscore/Infrastructure/Repos"
)

type SaveScoreTotalService struct {
	Request requests_write.SaveScoreTotalRequest
}

func ExecTotal(s SaveScoreTotalService) {
	repo.InsertScore(domain.UserScore{User: s.Request.User, Score: s.Request.Total})
}
