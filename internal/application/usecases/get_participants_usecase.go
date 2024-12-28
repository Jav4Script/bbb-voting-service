package usecase

import (
	"log"

	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type GetParticipantsUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
}

func NewGetParticipantsUsecase(participantRepository repositories.ParticipantRepository) *GetParticipantsUsecase {
	return &GetParticipantsUsecase{
		ParticipantRepository: participantRepository,
	}
}

func (usecase *GetParticipantsUsecase) Execute() ([]entities.Participant, error) {
	participants, err := usecase.ParticipantRepository.FindAll()
	if err != nil {
		log.Printf("Error retrieving participants: %v", err)
		return nil, errors.NewInfrastructureError("Failed to retrieve participants")
	}

	return participants, nil
}
