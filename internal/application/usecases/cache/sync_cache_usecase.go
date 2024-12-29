package cache

import (
	"log"

	"bbb-voting-service/internal/domain/repositories"
)

type SyncCacheUsecase struct {
	VoteRepository     repositories.VoteRepository
	InMemoryRepository repositories.InMemoryRepository
}

func NewSyncCacheUsecase(voteRepo repositories.VoteRepository, redisRepo repositories.InMemoryRepository) *SyncCacheUsecase {
	return &SyncCacheUsecase{
		VoteRepository:     voteRepo,
		InMemoryRepository: redisRepo,
	}
}

func (usecase *SyncCacheUsecase) Execute() error {
	finalResults, err := usecase.VoteRepository.GetFinalResults()
	if err != nil {
		log.Printf("Error getting final results from database: %v", err)
		return err
	}

	err = usecase.InMemoryRepository.UpdateCacheWithFinalResults(finalResults)
	if err != nil {
		log.Printf("Error updating cache with final results: %v", err)
		return err
	}

	return nil
}
