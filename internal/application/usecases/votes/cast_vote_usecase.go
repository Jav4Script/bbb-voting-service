package votes

import (
	"encoding/json"
	"log"
	"net/http"

	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	producer "bbb-voting-service/internal/domain/producer"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type CastVoteUsecase struct {
	InMemoryRepository    repositories.InMemoryRepository
	MessageProducer       producer.MessageProducer
	ParticipantRepository repositories.ParticipantRepository
	VoteQueue             string
}

func NewCastVoteUsecase(
	inMemoryRepository repositories.InMemoryRepository,
	messageProducer producer.MessageProducer,
	participantRepository repositories.ParticipantRepository,
	voteQueue string,
) *CastVoteUsecase {
	return &CastVoteUsecase{
		InMemoryRepository:    inMemoryRepository,
		MessageProducer:       messageProducer,
		ParticipantRepository: participantRepository,
		VoteQueue:             voteQueue,
	}
}

func (uscase *CastVoteUsecase) Execute(vote entities.Vote) error {
	// Validate participant
	participant, err := uscase.ParticipantRepository.FindByID(vote.ParticipantID)
	if err != nil {
		log.Printf("Error finding participant: %v", err)
		return errors.NewBusinessError("Participant not found", http.StatusNotFound)
	}

	// Update the partial results in cache
	err = uscase.InMemoryRepository.UpdatePartialResults(vote, participant)
	if err != nil {
		log.Printf("Error updating partial results: %v", err)
		return errors.NewInfrastructureError("Failed to update partial results in cache")
	}

	// Publish vote to message queue
	voteData, err := json.Marshal(vote)
	if err != nil {
		log.Printf("Error marshaling vote data: %v", err)
		return errors.NewInfrastructureError("Failed to serialize vote data")
	}

	err = uscase.MessageProducer.Publish(uscase.VoteQueue, voteData)
	if err != nil {
		log.Printf("Error publishing vote to message queue: %v", err)
		return errors.NewInfrastructureError("Failed to publish vote to message queue")
	}

	return nil
}
