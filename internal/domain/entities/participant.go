package domain

type Participant struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Votes int    `json:"votes"`
}
