package results

import (
	domain "bbb-voting-service/internal/domain/entities"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type GetPartialResultsUsecase struct {
	InMemoryResultRepository repositories.InMemoryResultRepository
}

func NewGetPartialResultsUsecase(inMemoryRepository repositories.InMemoryResultRepository) *GetPartialResultsUsecase {
	return &GetPartialResultsUsecase{
		InMemoryResultRepository: inMemoryRepository,
	}
}

func (usecase *GetPartialResultsUsecase) Execute() ([]domain.PartialResult, error) {
	partialResults, err := usecase.InMemoryResultRepository.GetPartialResults()
	if err != nil {
		return nil, err
	}

	return partialResults, nil
}
