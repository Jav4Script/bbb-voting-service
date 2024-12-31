package participants

import (
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

func (usecase *GetParticipantsUsecase) Execute() ([]entities.Participant, error) {
	// Try to get participants from cache
	participants, err := usecase.InMemoryParticipantRepository.FindAll()
	if err == nil {
		return participants, nil
	}

	// If not found in cache, get from database
	participants, err = usecase.ParticipantRepository.FindAll()
	if err != nil {
		log.Printf("Error retrieving participants: %v", err)
		return nil, errors.NewInfrastructureError("Failed to retrieve participants")
	}

	// Save participants to cache
	for _, participant := range participants {
		err = usecase.InMemoryParticipantRepository.Save(participant)
		if err != nil {
			log.Printf("Error saving participant to cache: %v", err)
		}
	}

	return participants, nil
}
