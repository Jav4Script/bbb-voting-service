package results

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	"bbb-voting-service/internal/domain/repositories"
)

type GetFinalResultsUsecase struct {
	ResultRepository repositories.ResultRepository
}

func NewGetFinalResultsUsecase(resultRepository repositories.ResultRepository) *GetFinalResultsUsecase {
	return &GetFinalResultsUsecase{
		ResultRepository: resultRepository,
	}
}

func (usecase *GetFinalResultsUsecase) Execute(ctx context.Context) (entities.FinalResults, error) {
	finalResults, err := usecase.ResultRepository.GetFinalResults()
	if err != nil {
		log.Printf("Error retrieving final results: %v", err)
		return entities.FinalResults{}, errors.NewInfrastructureError("Failed to retrieve final results")
	}

	return finalResults, nil
}
