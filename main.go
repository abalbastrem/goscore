package main

import (
	domain "Goscore/Domain"
	infra "Goscore/Infrastructure"
	"fmt"
)

func main() {
	fmt.Println("MAIN")
	infra.Hello()
	a := domain.UserScore{User: 1, Score: 45}
	fmt.Println(a)
}
