package votes

import (
	entities "bbb-voting-service/internal/domain/entities"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type ProcessVoteUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
	ResultRepository      repositories.ResultRepository
}

func NewProcessVoteUsecase(
	resultRepository repositories.ResultRepository,
	participantRepository repositories.ParticipantRepository,
) *ProcessVoteUsecase {
	return &ProcessVoteUsecase{
		ResultRepository:      resultRepository,
		ParticipantRepository: participantRepository,
	}
}

func (usecase *ProcessVoteUsecase) Execute(vote entities.Vote) error {
	// Validate participant
	_, err := usecase.ParticipantRepository.FindByID(vote.ParticipantID)
	if err != nil {
		return err
	}

	// Persist vote in the database
	err = usecase.ResultRepository.Save(vote)
	if err != nil {
		return err
	}

	return nil
}
