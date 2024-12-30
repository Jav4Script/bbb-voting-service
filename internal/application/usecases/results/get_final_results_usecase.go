package results

import (
	"bbb-voting-service/internal/domain/entities"
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

func (usecase *GetFinalResultsUsecase) Execute() (*entities.FinalResults, error) {
	// Get total votes
	totalVotes, err := usecase.VoteRepository.CountTotalVotes()
	if err != nil {
		return nil, err
	}

	// Get votes by participant
	participantResults, err := usecase.VoteRepository.GetParticipantResults()
	if err != nil {
		return nil, err
	}

	// Get votes by hour
	votesByHour, err := usecase.VoteRepository.CountVotesByHour()
	if err != nil {
		return nil, err
	}

	finalResultsEntity := mappers.ToFinalResults(participantResults, totalVotes, votesByHour)

	return &finalResultsEntity, nil
}
