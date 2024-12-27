package repositories

import (
	"time"

	domain "bbb-voting-service/internal/domain/entities"
)

type VoteRepository interface {
	Save(vote domain.Vote) error
	CountTotalVotes() (int, error)
	CountVotesByParticipant(participantID string) (int, error)
	CountVotesByHour(sessionID string) (map[time.Time]int, error)
	GetFinalResults() (map[string]int, error)
}
