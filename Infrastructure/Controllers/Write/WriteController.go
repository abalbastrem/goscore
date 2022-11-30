package controllers_write

import (
	requests_write "Goscore/Application/Requests/Write"
	services_write "Goscore/Application/Services/Write"
	"encoding/json"
	"net/http"
)

var HOST = "http://localhost:8080"
var URL_NEW_TOTAL = "/new_total"
var URL_NEW_DIFFERENTIAL = "/new_diff"

func SaveScoreTotal(w http.ResponseWriter, r *http.Request) {
	var request requests_write.SaveScoreTotalRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	service := services_write.SaveScoreTotalService{Request: request}
	services_write.ExecTotal(service)
}

func SaveScoreDifferential(w http.ResponseWriter, r *http.Request) {
	var request requests_write.SaveScoreDiffRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	service := services_write.SaveScoreDiffService{Request: request}
	services_write.ExecDiff(service)
}
