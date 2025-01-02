package participants

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	"bbb-voting-service/internal/domain/repositories"
)

type CreateParticipantUsecase struct {
	InMemoryParticipantRepository repositories.InMemoryParticipantRepository
	ParticipantRepository         repositories.ParticipantRepository
}

func NewCreateParticipantUsecase(participantRepository repositories.ParticipantRepository, inMemoryParticipantRepository repositories.InMemoryParticipantRepository) *CreateParticipantUsecase {
	return &CreateParticipantUsecase{
		InMemoryParticipantRepository: inMemoryParticipantRepository,
		ParticipantRepository:         participantRepository,
	}
}

func (usecase *CreateParticipantUsecase) Execute(context context.Context, participant entities.Participant) (entities.Participant, error) {
	// Save the participant in the database
	createdParticipant, err := usecase.ParticipantRepository.Save(participant)
	if err != nil {
		log.Printf("Error creating participant in database: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to create participant in database")
	}

	// Save the participant in the in-memory cache
	err = usecase.InMemoryParticipantRepository.Save(context, createdParticipant)
	if err != nil {
		log.Printf("Error saving participant in memory: %v", err)

		// Revert the database operation to maintain consistency
		rollbackErr := usecase.ParticipantRepository.Delete(createdParticipant)
		if rollbackErr != nil {
			log.Printf("Error rolling back participant creation: %v", rollbackErr)
		}

		return entities.Participant{}, errors.NewInfrastructureError("Failed to save participant in memory")
	}

	return createdParticipant, nil
}
