package results

import (
	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/repositories"
)

type GetFinalResultsUsecase struct {
	VoteRepository repositories.VoteRepository
}

func NewGetFinalResultsUsecase(voteRepository repositories.VoteRepository) *GetFinalResultsUsecase {
	return &GetFinalResultsUsecase{
		VoteRepository: voteRepository,
	}
}

func (usecase *GetFinalResultsUsecase) Execute() (entities.FinalResults, error) {
	finalResults, err := usecase.VoteRepository.GetFinalResults()
	if err != nil {
		return entities.FinalResults{}, err
	}

	return finalResults, nil
}
