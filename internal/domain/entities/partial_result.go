package domain

type PartialResult struct {
	ParticipantID     string `json:"participant_id"`
	ParticipantName   string `json:"participant_name"`
	ParticipantAge    int    `json:"participant_age"`
	ParticipantGender string `json:"participant_gender"`
	Votes             int    `json:"votes"`
}
