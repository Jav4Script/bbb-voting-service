package participants

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/errors"
	"bbb-voting-service/internal/domain/repositories"
)

type DeleteParticipantUsecase struct {
	InMemoryParticipantRepository repositories.InMemoryParticipantRepository
	ParticipantRepository         repositories.ParticipantRepository
}

func NewDeleteParticipantUsecase(participantRepository repositories.ParticipantRepository, inMemoryParticipantRepository repositories.InMemoryParticipantRepository) *DeleteParticipantUsecase {
	return &DeleteParticipantUsecase{
		InMemoryParticipantRepository: inMemoryParticipantRepository,
		ParticipantRepository:         participantRepository,
	}
}

func (usecase *DeleteParticipantUsecase) Execute(context context.Context, id string) error {
	// Find the participant in the database to ensure it exists
	participant, err := usecase.ParticipantRepository.FindByID(id)
	if err != nil {
		log.Printf("Error finding participant: %v", err)
		return errors.NewInfrastructureError("Failed to find participant")
	}

	// Delete the participant from the database
	err = usecase.ParticipantRepository.Delete(participant)
	if err != nil {
		log.Printf("Error deleting participant from database: %v", err)
		return errors.NewInfrastructureError("Failed to delete participant from database")
	}

	// Attempt to delete the participant from the in-memory cache (non-critical)
	err = usecase.InMemoryParticipantRepository.Delete(context, id)
	if err != nil {
		log.Printf("Warning: Failed to delete participant %s from cache: %v", id, err)
	}

	return nil
}
