package results

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	"bbb-voting-service/internal/domain/repositories"
)

type GetPartialResultsUsecase struct {
	InMemoryResultRepository repositories.InMemoryResultRepository
}

func NewGetPartialResultsUsecase(inMemoryRepository repositories.InMemoryResultRepository) *GetPartialResultsUsecase {
	return &GetPartialResultsUsecase{
		InMemoryResultRepository: inMemoryRepository,
	}
}

func (usecase *GetPartialResultsUsecase) Execute(context context.Context) ([]entities.PartialResult, error) {
	// Call the repository's GetPartialResults method, passing the context
	partialResults, err := usecase.InMemoryResultRepository.GetPartialResults(context)
	if err != nil {
		log.Printf("Error retrieving partial results: %v", err)
		return nil, errors.NewInfrastructureError("Failed to retrieve partial results")
	}

	return partialResults, nil
}
