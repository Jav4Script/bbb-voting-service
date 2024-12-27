package repositories

import domain "bbb-voting-service/internal/domain/entities"

type InMemoryRepository interface {
	GetPartialResults() (map[string]int, error)
	SavePartialVote(vote domain.Vote) error
}
