package requests_write

type SaveScoreTotalRequest struct {
	User  int `json:"user"`
	Total int `json:"total"`
}
