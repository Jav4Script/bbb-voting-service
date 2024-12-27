package usecase

import (
	repositories "bbb-voting-service/internal/domain/repositories"
)

type GetFinalResultsUsecase struct {
	VoteRepository repositories.VoteRepository
}

func NewGetFinalResultsUseCase(voteRepository repositories.VoteRepository) *GetFinalResultsUsecase {
	return &GetFinalResultsUsecase{VoteRepository: voteRepository}
}

func (usecase *GetFinalResultsUsecase) Execute() (map[string]interface{}, error) {
	// Get total votes
	totalVotes, err := usecase.VoteRepository.CountTotalVotes()
	if err != nil {
		return nil, err
	}

	// Get votes by participant
	finalResults, err := usecase.VoteRepository.GetFinalResults()
	if err != nil {
		return nil, err
	}

	// Get votes by hour (assuming sessionID is provided)
	sessionID := "some-session-id" // This should be passed as a parameter
	votesByHour, err := usecase.VoteRepository.CountVotesByHour(sessionID)
	if err != nil {
		return nil, err
	}

	results := map[string]interface{}{
		"total_votes":   totalVotes,
		"final_results": finalResults,
		"votes_by_hour": votesByHour,
	}

	return results, nil
}
