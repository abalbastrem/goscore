package tests

import (
	requests_write "Goscore/Application/Requests/Write"
	domain "Goscore/Domain"
	repo "Goscore/Infrastructure/Repos"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type topScores map[int]domain.UserScore

var client = http.Client{
	Timeout: 5 * time.Second,
}
var host = "http://localhost.com"
var urlRoot = "/"
var urlNew = "/new"
var urlFetchAbsolute = "/fetch_absolute"
var urlFetchRelative = "/fetch_relative"

func TestFirst(t *testing.T) {
	fmt.Println("TESTS A GO")
}

// MAINTENANCE
func TestPopulateDatabase(t *testing.T) {
	for i := 0; i <= 1000; i++ {
		score := rand.Int31()
		userScore := domain.UserScore{User: i, Score: int(score)}
		repo.SaveScore(userScore)
	}

	fmt.Println(repo.UserScores)
}

// READ
func TestHelloWorld(t *testing.T) {
	url := host + urlRoot
	expected := `{"hello": "world"}`

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	got := string(body)

	if expected != got {
		t.Errorf("Expected '%s', but got '%s'", expected, got)
	}
}

func TestGetAbsoluteTop100(t *testing.T) {
	rank := 100
	params := "?type=absolute&rank=" + strconv.Itoa(rank)
	url := host + urlFetchAbsolute + params

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var scores topScores
	json.Unmarshal(body, &scores)
	if len(scores) != rank {
		t.Errorf("ERROR there should be '%d' scores, but were fetched '%d'", 100, len(scores))
	}
	currTotal := 0
	for i := 1; i <= rank; i++ {
		if scores[i].Score < currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

func TestGetAbsoluteTop200(t *testing.T) {
	rank := 200
	params := "?type=absolute&rank=" + strconv.Itoa(rank)
	url := host + urlFetchAbsolute + params

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var scores topScores
	json.Unmarshal(body, &scores)
	if len(scores) != rank {
		t.Errorf("ERROR there should be '%d' scores, but were fetched '%d'", 100, len(scores))
	}
	currTotal := 0
	for i := 1; i <= rank; i++ {
		if scores[i].Score < currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

func TestGetAbsoluteTop500(t *testing.T) {
	rank := 500
	params := "?type=absolute&rank=" + strconv.Itoa(rank)
	url := host + urlFetchAbsolute + params

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var scores topScores
	json.Unmarshal(body, &scores)
	if len(scores) != rank {
		t.Errorf("ERROR there should be '%d' scores, but were fetched '%d'", 100, len(scores))
	}
	currTotal := 0
	for i := 1; i <= rank; i++ {
		if scores[i].Score < currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

func TestGetRelativeAndThereAreEnoughRelatives(t *testing.T) {
	rank := 120
	nRelatives := 3
	params := "?type=relative&rank=" + strconv.Itoa(rank) + "&n_relatives=" + strconv.Itoa(nRelatives)
	url := host + urlFetchRelative + params

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var scores topScores
	json.Unmarshal(body, &scores)

	currTotal := scores[rank].Score
	for i := rank; i <= rank+nRelatives; i++ {
		if scores[i].Score < currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
	currTotal = scores[rank].Score
	for i := rank; i >= rank-nRelatives; i-- {
		if scores[i].Score > currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

func TestGetRelativeAndThereAreNotEnoughRelatives(t *testing.T) {
	rank := 998
	nRelatives := 5
	params := "?type=relative&rank=" + strconv.Itoa(rank) + "&n_relatives=" + strconv.Itoa(nRelatives)
	url := host + urlFetchRelative + params

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var scores topScores
	json.Unmarshal(body, &scores)

	currTotal := scores[rank].Score
	for i := rank; i <= rank+nRelatives; i++ {
		if scores[i].Score < currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
	currTotal = scores[rank].Score
	for i := rank; i >= rank-nRelatives; i-- {
		if scores[i].Score > currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

// WRITE
func TestPostFirstSaveScoreAbsolute(t *testing.T) {
	url := host + urlNew
	saveScore := requests_write.SaveScoreTotalRequest{User: 2000, Total: 20000}
	saveScoreJson, _ := json.Marshal(saveScore)
	res, err := client.Post(url, "application/json", bytes.NewBuffer(saveScoreJson))
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Error()
	}
}

func TestPostUpdateScoreWithAbsolute(t *testing.T) {
	url := host + urlNew
	saveScore := requests_write.SaveScoreTotalRequest{User: 2000, Total: 20001}
	saveScoreJson, _ := json.Marshal(saveScore)
	res, err := client.Post(url, "application/json", bytes.NewBuffer(saveScoreJson))
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Error()
	}
}

func TestPostUpdateScoreWithDifferential(t *testing.T) {
	url := host + urlNew
	saveScore := requests_write.SaveScoreDiffRequest{User: 2000, Score: "+10"}
	saveScoreJson, _ := json.Marshal(saveScore)
	res, err := client.Post(url, "application/json", bytes.NewBuffer(saveScoreJson))
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Error()
	}
}

func TestPostFirstSaveScoreRelative(t *testing.T) {
	url := host + urlNew
	saveScore := requests_write.SaveScoreDiffRequest{User: 3000, Score: "+30000"}
	saveScoreJson, _ := json.Marshal(saveScore)
	res, err := client.Post(url, "application/json", bytes.NewBuffer(saveScoreJson))
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Error()
	}
}

// MAINTENANCE
func TestEmptyDatabase(t *testing.T) {

}
