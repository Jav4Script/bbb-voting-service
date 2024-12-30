package entities

type FinalResults struct {
	ParticipantResults []ParticipantResult `json:"participant_results"`
	TotalVotes         int                 `json:"total_votes"`
	VotesByHour        map[string]int      `json:"votes_by_hour"`
}
