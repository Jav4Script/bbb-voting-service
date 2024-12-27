package usecase

import (
	"encoding/json"

	entities "bbb-voting-service/internal/domain/entities"
	producer "bbb-voting-service/internal/domain/producer"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type CastVoteUsecase struct {
	InMemoryRepository    repositories.InMemoryRepository
	MessageProducer       producer.MessageProducer
	ParticipantRepository repositories.ParticipantRepository
}

func NewCastVoteUsecase(
	inMemoryRepository repositories.InMemoryRepository,
	messageProducer producer.MessageProducer,
	participantRepository repositories.ParticipantRepository,
) *CastVoteUsecase {
	return &CastVoteUsecase{
		InMemoryRepository:    inMemoryRepository,
		MessageProducer:       messageProducer,
		ParticipantRepository: participantRepository,
	}
}

func (uscase *CastVoteUsecase) Execute(vote entities.Vote) error {
	// Validate participant
	_, err := uscase.ParticipantRepository.FindByID(vote.ParticipantID)
	if err != nil {
		return err
	}

	// Store vote in Redis for partial results
	err = uscase.InMemoryRepository.SavePartialVote(vote)
	if err != nil {
		return err
	}

	// Publish vote to message queue
	voteData, err := json.Marshal(vote)
	if err != nil {
		return err
	}

	err = uscase.MessageProducer.Publish("vote_queue", voteData)
	if err != nil {
		return err
	}

	return nil
}
