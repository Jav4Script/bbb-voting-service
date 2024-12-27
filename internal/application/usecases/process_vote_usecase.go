package usecase

import (
	entities "bbb-voting-service/internal/domain/entities"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type ProcessVoteUsecase struct {
	VoteRepository        repositories.VoteRepository
	ParticipantRepository repositories.ParticipantRepository
}

func NewProcessVoteUsecase(
	voteRepository repositories.VoteRepository,
	participantRepository repositories.ParticipantRepository,
) *ProcessVoteUsecase {
	return &ProcessVoteUsecase{
		VoteRepository:        voteRepository,
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
	err = usecase.VoteRepository.Save(vote)
	if err != nil {
		return err
	}

	return nil
}
