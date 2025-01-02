package repositories

import (
	domain "bbb-voting-service/internal/domain/entities"
	"context"
)

type InMemoryResultRepository interface {
	GetPartialResults(ctx context.Context) ([]domain.PartialResult, error)
	UpdatePartialResults(ctx context.Context, vote domain.Vote, participant domain.Participant) error
	UpdateCacheWithFinalResults(ctx context.Context, finalResults domain.FinalResults) error
}
