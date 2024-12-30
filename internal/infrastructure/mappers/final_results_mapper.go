package mappers

import (
	entities "bbb-voting-service/internal/domain/entities"
	"time"
)

func ToFinalResults(participantResults map[string]entities.ParticipantResult, totalVotes int, votesByHour map[time.Time]int) entities.FinalResults {
	votesByHourString := make(map[string]int)
	for key, value := range votesByHour {
		votesByHourString[key.Format(time.RFC3339)] = value
	}

	finalResults := make([]entities.ParticipantResult, 0, len(participantResults))
	for _, result := range participantResults {
		finalResults = append(finalResults, result)
	}

	return entities.FinalResults{
		FinalResults: finalResults,
		TotalVotes:   totalVotes,
		VotesByHour:  votesByHourString,
	}
}
