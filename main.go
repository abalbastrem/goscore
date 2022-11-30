package main

import (
	utils "Goscore/Application/Utils"
	domain "Goscore/Domain"
	cr "Goscore/Infrastructure/Controllers/Read"
	repo "Goscore/Infrastructure/Repos"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	fmt.Println("START")
	populate()
	rough()

	http.HandleFunc(cr.URL_ROOT, cr.HelloWorld)
	http.HandleFunc(cr.URL_FETCH_ABSOLUTE, cr.GetAbsoluteTop)
	http.HandleFunc(cr.URL_FETCH_RELATIVE, cr.GetRelatives)

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func populate() {
	fmt.Print("populating database...")
	for i := 1; i <= 1000; i++ {
		score := rand.Int31()
		userScore := domain.UserScore{User: i, Score: int(score)}
		repo.SaveScore(userScore)
	}
	fmt.Println(" DONE")
}

func rough() {
	userScoreList := repo.GetScores()
	utils.Sort(userScoreList)
	userScoreMap := make(map[int]domain.UserScore)
	for index, userScore := range userScoreList {
		if index >= 120-3 &&
			index <= 120+3 {
			userScoreMap[index] = userScore
		}
	}
}
