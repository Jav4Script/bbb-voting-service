package participants

import (
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

func (usecase *DeleteParticipantUsecase) Execute(id string) error {
	participant, err := usecase.ParticipantRepository.FindByID(id)
	if err != nil {
		log.Printf("Error finding participant: %v", err)
		return errors.NewInfrastructureError("Failed to find participant")
	}

	err = usecase.ParticipantRepository.Delete(participant)
	if err != nil {
		log.Printf("Error deleting participant: %v", err)
		return errors.NewInfrastructureError("Failed to delete participant")
	}

	// Delete participant from cache
	err = usecase.InMemoryParticipantRepository.Delete(id)
	if err != nil {
		log.Printf("Error deleting participant from cache: %v", err)
	}

	return nil
}
