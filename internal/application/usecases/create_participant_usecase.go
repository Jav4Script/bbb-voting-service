package usecase

import (
	"log"

	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type CreateParticipantUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
}

func NewCreateParticipantUsecase(participantRepository repositories.ParticipantRepository) *CreateParticipantUsecase {
	return &CreateParticipantUsecase{
		ParticipantRepository: participantRepository,
	}
}

func (usecase *CreateParticipantUsecase) Execute(participant entities.Participant) (entities.Participant, error) {
	createdParticipant, err := usecase.ParticipantRepository.Save(participant)
	if err != nil {
		log.Printf("Error creating participant: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to create participant")
	}

	return createdParticipant, nil
}
