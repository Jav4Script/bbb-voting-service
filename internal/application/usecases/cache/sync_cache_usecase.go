package cache

import (
	"log"

	"bbb-voting-service/internal/application/usecases/results"
	"bbb-voting-service/internal/domain/repositories"
)

type SyncCacheUsecase struct {
	GetFinalResultsUsecase *results.GetFinalResultsUsecase
	InMemoryRepository     repositories.InMemoryRepository
}

func NewSyncCacheUsecase(getFinalResultsUsecase *results.GetFinalResultsUsecase, redisRepo repositories.InMemoryRepository) *SyncCacheUsecase {
	return &SyncCacheUsecase{
		GetFinalResultsUsecase: getFinalResultsUsecase,
		InMemoryRepository:     redisRepo,
	}
}

func (usecase *SyncCacheUsecase) Execute() error {
	finalResults, err := usecase.GetFinalResultsUsecase.Execute()
	if err != nil {
		log.Printf("Error getting final results: %v", err)
		return err
	}

	err = usecase.InMemoryRepository.UpdateCacheWithFinalResults(*finalResults)
	if err != nil {
		log.Printf("Error updating cache with final results: %v", err)
		return err
	}

	return nil
}
