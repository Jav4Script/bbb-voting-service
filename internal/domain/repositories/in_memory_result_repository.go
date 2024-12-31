package repositories

import domain "bbb-voting-service/internal/domain/entities"

type InMemoryResultRepository interface {
	GetPartialResults() ([]domain.PartialResult, error)
	UpdatePartialResults(vote domain.Vote, participant domain.Participant) error
	UpdateCacheWithFinalResults(finalResults domain.FinalResults) error
}
