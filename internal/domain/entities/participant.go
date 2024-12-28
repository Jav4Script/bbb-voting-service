package domain

type Participant struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
