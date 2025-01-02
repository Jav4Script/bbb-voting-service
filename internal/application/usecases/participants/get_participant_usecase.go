package participants

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	"bbb-voting-service/internal/domain/repositories"
)

type GetParticipantUsecase struct {
	InMemoryParticipantRepository repositories.InMemoryParticipantRepository
	ParticipantRepository         repositories.ParticipantRepository
}

func NewGetParticipantUsecase(participantRepository repositories.ParticipantRepository, inMemoryParticipantRepository repositories.InMemoryParticipantRepository) *GetParticipantUsecase {
	return &GetParticipantUsecase{
		InMemoryParticipantRepository: inMemoryParticipantRepository,
		ParticipantRepository:         participantRepository,
	}
}

func (usecase *GetParticipantUsecase) Execute(context context.Context, id string) (entities.Participant, error) {
	// Try to get the participant from the in-memory cache
	participant, err := usecase.InMemoryParticipantRepository.FindByID(context, id)
	if err == nil {
		return participant, nil
	}

	// If not found in the cache, fetch the participant from the database
	participant, err = usecase.ParticipantRepository.FindByID(id)
	if err != nil {
		log.Printf("Error retrieving participant from database: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to retrieve participant from database")
	}

	// Save the participant to the in-memory cache
	err = usecase.InMemoryParticipantRepository.Save(context, participant)
	if err != nil {
		log.Printf("Warning: Failed to update participant in cache: %v", err)
		// Optional: Decide whether to return an error or continue
	}

	return participant, nil
}
