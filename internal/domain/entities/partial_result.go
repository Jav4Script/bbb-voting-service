package entities

type PartialResult struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Votes  int    `json:"votes"`
}
