package entities

type FinalResults struct {
	FinalResults []ParticipantResult `json:"final_results"`
	TotalVotes   int                 `json:"total_votes"`
	VotesByHour  map[string]int      `json:"votes_by_hour"`
}
