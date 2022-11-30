package tests

import (
	requests_write "Goscore/Application/Requests/Write"
	domain "Goscore/Domain"
	controller_read "Goscore/Infrastructure/Controllers/Read"
	controller_write "Goscore/Infrastructure/Controllers/Write"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type topScores map[int]domain.UserScore

var client = http.Client{
	Timeout: 5 * time.Second,
}

func TestFirst(t *testing.T) {
	fmt.Println("TESTS A GO")
}

// READ
func TestHelloWorld(t *testing.T) {
	url := controller_read.HOST + controller_read.URL_ROOT
	expected := "Hello world"

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
	params := "?rank=" + strconv.Itoa(rank)
	url := controller_read.HOST + controller_read.URL_FETCH_ABSOLUTE + params

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
	currTotal := scores[1].Score
	for i := 1; i <= rank; i++ {
		if scores[i].Score > currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

func TestGetAbsoluteTop200(t *testing.T) {
	rank := 200
	params := "?rank=" + strconv.Itoa(rank)
	url := controller_read.HOST + controller_read.URL_FETCH_ABSOLUTE + params

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
	currTotal := scores[1].Score
	for i := 1; i <= rank; i++ {
		if scores[i].Score > currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

func TestGetAbsoluteTop500(t *testing.T) {
	rank := 500
	params := "?rank=" + strconv.Itoa(rank)
	url := controller_read.HOST + controller_read.URL_FETCH_ABSOLUTE + params

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
	currTotal := scores[1].Score
	for i := 1; i <= rank; i++ {
		if scores[i].Score > currTotal {
			t.Error("ERROR scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

func TestGetRelativeAndThereAreEnoughRelatives(t *testing.T) {
	rank := 120
	nRelatives := 3
	params := "?rank=" + strconv.Itoa(rank) + "&n_relatives=" + strconv.Itoa(nRelatives)
	url := controller_read.HOST + controller_read.URL_FETCH_RELATIVE + params

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var scores topScores
	json.Unmarshal(body, &scores)

	if (nRelatives*2)+1 != len(scores) {
		t.Errorf("ERROR wrong number of scores fetched. Expected '%d', got '%d'", (nRelatives*2)+1, len(scores))
	}

	currTotal := scores[rank].Score
	for i := rank; i <= rank+nRelatives; i++ {
		if scores[i].Score > currTotal {
			t.Error("ERROR_1 scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
	currTotal = scores[rank].Score
	for i := rank; i >= rank-nRelatives; i-- {
		if scores[i].Score < currTotal {
			t.Error("ERROR_2 scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

func TestGetRelativeAndThereAreNotEnoughRelatives(t *testing.T) {
	rank := 998
	nRelatives := 5
	params := "?rank=" + strconv.Itoa(rank) + "&n_relatives=" + strconv.Itoa(nRelatives)
	url := controller_read.HOST + controller_read.URL_FETCH_RELATIVE + params

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var scores topScores
	json.Unmarshal(body, &scores)

	if (nRelatives*2)+1 <= len(scores) {
		t.Errorf("ERROR wrong number of scores fetched. Expected '%d', should be larger than got '%d'", (nRelatives*2)+1, len(scores))
	}

	currTotal := scores[rank].Score
	for i := rank; i <= rank+nRelatives; i++ {
		if scores[i].Score > currTotal {
			t.Error("ERROR_1 scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
	currTotal = scores[rank].Score
	for i := rank; i >= rank-nRelatives; i-- {
		if scores[i].Score < currTotal {
			t.Error("ERROR_2 scores were not fetched in the correct order")
		}
		currTotal = scores[i].Score
	}
}

// WRITE
func TestPostFirstSaveScoreTotal(t *testing.T) {
	url := controller_write.HOST + controller_write.URL_NEW_TOTAL
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

func TestPostUpdateScoreWithTotal(t *testing.T) {
	url := controller_write.HOST + controller_write.URL_NEW_TOTAL
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
	url := controller_write.HOST + controller_write.URL_NEW_DIFFERENTIAL
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

func TestPostFirstSaveScoreDifferential(t *testing.T) {
	url := controller_write.HOST + controller_write.URL_NEW_DIFFERENTIAL
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
