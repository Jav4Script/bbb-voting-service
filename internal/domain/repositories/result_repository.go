package repositories

import (
	"time"

	"bbb-voting-service/internal/domain/entities"
)

type ResultRepository interface {
	Save(vote entities.Vote) error
	CountTotalVotes() (int, error)
	CountVotesByParticipant(participantID string) (int, error)
	CountVotesByHour() (map[time.Time]int, error)
	GetParticipantResults() (map[string]entities.ParticipantResult, error)
	GetFinalResults() (entities.FinalResults, error)
}
