package cache

import (
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

func (usecase *SyncResultsCacheUsecase) Execute() error {
	finalResults, err := usecase.GetFinalResultsUsecase.Execute()
	if err != nil {
		log.Printf("Error getting final results: %v", err)
		return err
	}

	err = usecase.InMemoryResultRepository.UpdateCacheWithFinalResults(finalResults)
	if err != nil {
		log.Printf("Error updating cache with final results: %v", err)
		return err
	}

	return nil
}
