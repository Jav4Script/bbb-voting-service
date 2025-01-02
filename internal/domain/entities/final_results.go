package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type FinalResults struct {
	ParticipantResults []ParticipantResult `json:"participant_results"`
	TotalVotes         int                 `json:"total_votes"`
	VotesByHour        map[string]int      `json:"votes_by_hour"`
}

// Scan implements the sql.Scanner interface for ParticipantResults
func (fr *FinalResults) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &fr.ParticipantResults)
}

// Value implements the driver.Valuer interface for ParticipantResults
func (fr FinalResults) Value() (driver.Value, error) {
	return json.Marshal(fr.ParticipantResults)
}
