package participants

import (
	"log"

	"bbb-voting-service/internal/domain/errors"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type DeleteParticipantUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
}

func NewDeleteParticipantUsecase(participantRepository repositories.ParticipantRepository) *DeleteParticipantUsecase {
	return &DeleteParticipantUsecase{
		ParticipantRepository: participantRepository,
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

	return nil
}
