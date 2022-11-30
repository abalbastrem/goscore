package main

import (
	domain "Goscore/Domain"
	cr "Goscore/Infrastructure/Controllers/Read"
	cw "Goscore/Infrastructure/Controllers/Write"
	repo "Goscore/Infrastructure/Repos"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	fmt.Println("START")
	populate()

	http.HandleFunc(cr.URL_ROOT, cr.HelloWorld)
	http.HandleFunc(cr.URL_FETCH_ABSOLUTE, cr.GetAbsoluteTop)
	http.HandleFunc(cr.URL_FETCH_RELATIVE, cr.GetRelatives)

	http.HandleFunc(cw.URL_NEW_TOTAL, cw.SaveScoreTotal)
	http.HandleFunc(cw.URL_NEW_DIFFERENTIAL, cw.SaveScoreDifferential)

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func populate() {
	fmt.Print("populating database...")
	for i := 1; i <= 1000; i++ {
		score := rand.Int31()
		userScore := domain.UserScore{User: i, Score: int(score)}
		repo.InsertScore(userScore)
	}
	fmt.Println(" DONE")
}
