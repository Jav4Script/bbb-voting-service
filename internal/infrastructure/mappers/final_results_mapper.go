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

	participantResultsSlice := make([]entities.ParticipantResult, 0, len(participantResults))
	for _, result := range participantResults {
		participantResultsSlice = append(participantResultsSlice, result)
	}

	return entities.FinalResults{
		ParticipantResults: participantResultsSlice,
		TotalVotes:         totalVotes,
		VotesByHour:        votesByHourString,
	}
}
