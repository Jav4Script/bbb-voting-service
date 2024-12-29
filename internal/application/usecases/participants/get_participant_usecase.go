package participants

import (
	"log"

	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type GetParticipantUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
}

func NewGetParticipantUsecase(participantRepository repositories.ParticipantRepository) *GetParticipantUsecase {
	return &GetParticipantUsecase{
		ParticipantRepository: participantRepository,
	}
}

func (usecase *GetParticipantUsecase) Execute(id string) (entities.Participant, error) {
	participant, err := usecase.ParticipantRepository.FindByID(id)
	if err != nil {
		log.Printf("Error retrieving participant: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to retrieve participant")
	}

	return participant, nil
}
