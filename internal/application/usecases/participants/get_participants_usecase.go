package participants

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	"bbb-voting-service/internal/domain/repositories"
)

type GetParticipantsUsecase struct {
	InMemoryParticipantRepository repositories.InMemoryParticipantRepository
	ParticipantRepository         repositories.ParticipantRepository
}

func NewGetParticipantsUsecase(participantRepository repositories.ParticipantRepository, inMemoryParticipantRepository repositories.InMemoryParticipantRepository) *GetParticipantsUsecase {
	return &GetParticipantsUsecase{
		InMemoryParticipantRepository: inMemoryParticipantRepository,
		ParticipantRepository:         participantRepository,
	}
}

func (usecase *GetParticipantsUsecase) Execute(context context.Context) ([]entities.Participant, error) {
	// Attempt to get participants from the in-memory cache
	participants, err := usecase.InMemoryParticipantRepository.FindAll(context)
	if err == nil {
		return participants, nil
	}

	// If not found in the cache, fetch them from the database
	participants, err = usecase.ParticipantRepository.FindAll()
	if err != nil {
		log.Printf("Error retrieving participants from database: %v", err)
		return nil, errors.NewInfrastructureError("Failed to retrieve participants from database")
	}

	// Save participants to the in-memory cache (only if they are not already present)
	for _, participant := range participants {
		_, cacheErr := usecase.InMemoryParticipantRepository.FindByID(context, participant.ID)
		if cacheErr == nil {
			// Participant is already in the cache, skip saving
			continue
		}

		err = usecase.InMemoryParticipantRepository.Save(context, participant)
		if err != nil {
			log.Printf("Warning: Failed to save participant %s in cache: %v", participant.ID, err)
		}
	}

	return participants, nil
}
