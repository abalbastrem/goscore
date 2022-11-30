package requests_write

type SaveScoreDiffRequest struct {
	User  int    `json:"user"`
	Score string `json:"score"`
}
