package entities

type PartialResult struct {
	ID     string `json:"id" msgpack:"id"`
	Name   string `json:"name" msgpack:"name"`
	Age    int    `json:"age" msgpack:"age"`
	Gender string `json:"gender" msgpack:"gender"`
	Votes  int    `json:"votes" msgpack:"votes"`
}
