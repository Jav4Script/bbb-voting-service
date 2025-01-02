package cache

import (
	"context"
	"log"

	"bbb-voting-service/internal/application/usecases/results"
	"bbb-voting-service/internal/domain/repositories"
)

type SyncResultsCacheUsecase struct {
	GetFinalResultsUsecase   *results.GetFinalResultsUsecase
	InMemoryResultRepository repositories.InMemoryResultRepository
}

func NewSyncResultsCacheUsecase(getFinalResultsUsecase *results.GetFinalResultsUsecase, redisRepo repositories.InMemoryResultRepository) *SyncResultsCacheUsecase {
	return &SyncResultsCacheUsecase{
		GetFinalResultsUsecase:   getFinalResultsUsecase,
		InMemoryResultRepository: redisRepo,
	}
}

func (usecase *SyncResultsCacheUsecase) Execute(context context.Context) error {
	// Fetch final results using the provided use case
	finalResults, err := usecase.GetFinalResultsUsecase.Execute(context)
	if err != nil {
		log.Printf("Error getting final results: %v", err)
		return err
	}

	// Update the cache with the final results
	err = usecase.InMemoryResultRepository.UpdateCacheWithFinalResults(context, finalResults)
	if err != nil {
		log.Printf("Error updating cache with final results: %v", err)
		return err
	}

	return nil
}
