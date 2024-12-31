package participants

import (
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

func (usecase *GetParticipantUsecase) Execute(id string) (entities.Participant, error) {
	// Try to get participant from cache
	participant, err := usecase.InMemoryParticipantRepository.FindByID(id)
	if err == nil {
		return participant, nil
	}

	// If not found in cache, get from database
	participant, err = usecase.ParticipantRepository.FindByID(id)
	if err != nil {
		log.Printf("Error retrieving participant: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to retrieve participant")
	}

	// Save participant to cache
	err = usecase.InMemoryParticipantRepository.Save(participant)
	if err != nil {
		log.Printf("Error saving participant to cache: %v", err)
	}

	return participant, nil
}
