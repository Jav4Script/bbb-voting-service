package usecase

import (
	repositories "bbb-voting-service/internal/domain/repositories"
)

type GetPartialResultsUsecase struct {
	InMemoryRepository repositories.InMemoryRepository
}

func NewGetPartialResultsUsecase(inMemoryRepository repositories.InMemoryRepository) *GetPartialResultsUsecase {
	return &GetPartialResultsUsecase{
		InMemoryRepository: inMemoryRepository,
	}
}

func (usecase *GetPartialResultsUsecase) Execute() (map[string]int, error) {
	partialResults, err := usecase.InMemoryRepository.GetPartialResults()
	if err != nil {
		return nil, err
	}

	return partialResults, nil
}
