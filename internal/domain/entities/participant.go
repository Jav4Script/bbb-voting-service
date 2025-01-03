package entities

type Participant struct {
	ID        string `json:"id" msgpack:"id"`
	Name      string `json:"name" msgpack:"name"`
	Age       int    `json:"age" msgpack:"age"`
	Gender    string `json:"gender" msgpack:"gender"`
	CreatedAt string `json:"created_at" msgpack:"created_at"`
	UpdatedAt string `json:"updated_at" msgpack:"updated_at"`
}
