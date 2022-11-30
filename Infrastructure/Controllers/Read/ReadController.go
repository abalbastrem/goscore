package controllers_read

import (
	utils "Goscore/Application/Utils"
	domain "Goscore/Domain"
	repo "Goscore/Infrastructure/Repos"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var HOST = "http://localhost:8080"
var URL_ROOT = "/"
var URL_NEW = "/new"
var URL_FETCH_ABSOLUTE = "/fetch_absolute"
var URL_FETCH_RELATIVE = "/fetch_relative"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func GetAbsoluteTop(w http.ResponseWriter, r *http.Request) {
	rankParam := r.URL.Query().Get("rank")
	rank, _ := strconv.Atoi(rankParam)
	var userScoreList []domain.UserScore
	for user, score := range repo.UserScores {
		userScoreList = append(userScoreList, domain.UserScore{User: user, Score: score})
	}
	utils.Sort(userScoreList)
	userScoreMap := make(map[int]domain.UserScore)
	for index, userScore := range userScoreList {
		userScoreMap[index+1] = userScore
		if len(userScoreMap) == rank {
			break
		}
	}
	jsonUserScores, _ := json.Marshal(userScoreMap)
	fmt.Fprintf(w, string(jsonUserScores))
}

func GetRelatives(w http.ResponseWriter, r *http.Request) {

}
