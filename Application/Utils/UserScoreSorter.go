package utils

import (
	domain "Goscore/Domain"
	"sort"
)

type userScoreList []domain.UserScore

func (us userScoreList) Len() int {
	return len(us)
}

func (us userScoreList) Less(i, j int) bool {
	return us[i].Score > us[j].Score
}

func (us userScoreList) Swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}

func Sort(us userScoreList) {
	sort.Sort(us)
}
