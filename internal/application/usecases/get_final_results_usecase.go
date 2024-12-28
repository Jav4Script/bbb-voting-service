package usecase

import (
	repositories "bbb-voting-service/internal/domain/repositories"
	mappers "bbb-voting-service/internal/infrastructure/mappers"
)

type GetFinalResultsUsecase struct {
	VoteRepository        repositories.VoteRepository
	ParticipantRepository repositories.ParticipantRepository
}

func NewGetFinalResultsUseCase(voteRepository repositories.VoteRepository, participantRepository repositories.ParticipantRepository) *GetFinalResultsUsecase {
	return &GetFinalResultsUsecase{
		VoteRepository:        voteRepository,
		ParticipantRepository: participantRepository,
	}
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

	// Get votes by hour
	votesByHour, err := usecase.VoteRepository.CountVotesByHour()
	if err != nil {
		return nil, err
	}

	// Prepare the final results with participant details
	finalResultsWithDetails := make([]map[string]interface{}, 0, len(finalResults))
	for participantID, votes := range finalResults {
		participant, err := usecase.ParticipantRepository.FindByID(participantID)
		if err != nil {
			return nil, err
		}

		result := mappers.ToParticipantMap(participant)
		result["votes"] = votes
		finalResultsWithDetails = append(finalResultsWithDetails, result)
	}

	results := map[string]interface{}{
		"total_votes":   totalVotes,
		"final_results": finalResultsWithDetails,
		"votes_by_hour": votesByHour,
	}

	return results, nil
}
