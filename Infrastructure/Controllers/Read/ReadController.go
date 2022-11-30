package controllers_read

import (
	requests_read "Goscore/Application/Requests/Read"
	services_read "Goscore/Application/Services/Read"
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

	request := requests_read.GetScoresAbsoluteRequest{Rank: rank}
	service := services_read.GetScoresAbsoluteService{Request: request}

	userScoreMap := services_read.ExecAbsolute(service)
	jsonUserScores, _ := json.Marshal(userScoreMap)
	fmt.Fprintf(w, string(jsonUserScores))
}

func GetRelatives(w http.ResponseWriter, r *http.Request) {
	rankParam := r.URL.Query().Get("rank")
	nRelativesParam := r.URL.Query().Get("n_relatives")
	rank, _ := strconv.Atoi(rankParam)
	nRelatives, _ := strconv.Atoi(nRelativesParam)

	request := requests_read.GetScoresRelativeRequest{Rank: rank, NRelatives: nRelatives}
	service := services_read.GetScoresRelativeService{Request: request}

	userScoreMap := services_read.ExecRelative(service)
	jsonUserScores, _ := json.Marshal(userScoreMap)
	fmt.Fprintf(w, string(jsonUserScores))
}
