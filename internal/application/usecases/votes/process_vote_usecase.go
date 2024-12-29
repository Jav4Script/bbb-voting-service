package votes

import (
	entities "bbb-voting-service/internal/domain/entities"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type ProcessVoteUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
	VoteRepository        repositories.VoteRepository
	InMemoryRepository    repositories.InMemoryRepository
}

func NewProcessVoteUsecase(
	voteRepository repositories.VoteRepository,
	participantRepository repositories.ParticipantRepository,
	inMemoryRepository repositories.InMemoryRepository,
) *ProcessVoteUsecase {
	return &ProcessVoteUsecase{
		VoteRepository:        voteRepository,
		ParticipantRepository: participantRepository,
		InMemoryRepository:    inMemoryRepository,
	}
}

func (usecase *ProcessVoteUsecase) Execute(vote entities.Vote) error {
	// Validate participant
	participant, err := usecase.ParticipantRepository.FindByID(vote.ParticipantID)
	if err != nil {
		return err
	}

	// Persist vote in the database
	err = usecase.VoteRepository.Save(vote)
	if err != nil {
		return err
	}

	// Update the partial results in Redis
	err = usecase.InMemoryRepository.UpdatePartialResults(vote, participant)
	if err != nil {
		return err
	}

	return nil
}
