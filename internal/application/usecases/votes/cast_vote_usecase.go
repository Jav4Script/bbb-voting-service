package votes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	producer "bbb-voting-service/internal/domain/producer"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type CastVoteUsecase struct {
	InMemoryParticipantRepository repositories.InMemoryParticipantRepository
	InMemoryResultRepository      repositories.InMemoryResultRepository
	MessageProducer               producer.MessageProducer
	VoteQueue                     string
}

// NewCastVoteUsecase creates a new instance of the CastVoteUsecase.
func NewCastVoteUsecase(
	inMemoryParticipantRepository repositories.InMemoryParticipantRepository,
	inMemoryResultRepository repositories.InMemoryResultRepository,
	messageProducer producer.MessageProducer,
	voteQueue string,
) *CastVoteUsecase {
	return &CastVoteUsecase{
		InMemoryParticipantRepository: inMemoryParticipantRepository,
		InMemoryResultRepository:      inMemoryResultRepository,
		MessageProducer:               messageProducer,
		VoteQueue:                     voteQueue,
	}
}

// Execute performs the logic for casting a vote.
func (usecase *CastVoteUsecase) Execute(context context.Context, vote entities.Vote) error {
	// Validate participant by ID
	participant, err := usecase.InMemoryParticipantRepository.FindByID(context, vote.ParticipantID)
	if err != nil {
		log.Printf("Error finding participant in cache: %v", err)
		return errors.NewBusinessError("Participant not found", http.StatusNotFound)
	}

	// Publish vote to message queue
	voteData, err := json.Marshal(vote)
	if err != nil {
		log.Printf("Error marshaling vote data: %v", err)
		return errors.NewInfrastructureError("Failed to serialize vote data")
	}

	err = usecase.MessageProducer.Publish(usecase.VoteQueue, voteData)
	if err != nil {
		log.Printf("Error publishing vote to message queue: %v", err)
		return errors.NewInfrastructureError("Failed to publish vote to message queue")
	}

	// Update the partial results in cache
	err = usecase.InMemoryResultRepository.UpdatePartialResults(context, vote, participant)
	if err != nil {
		log.Printf("Error updating partial results: %v", err)
		return errors.NewInfrastructureError("Failed to update partial results in cache")
	}

	return nil
}
