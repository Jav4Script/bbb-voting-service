package participants

import (
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

func (usecase *CreateParticipantUsecase) Execute(participant entities.Participant) (entities.Participant, error) {
	createdParticipant, err := usecase.ParticipantRepository.Save(participant)
	if err != nil {
		log.Printf("Error creating participant: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to create participant")
	}

	err = usecase.InMemoryParticipantRepository.Save(createdParticipant)
	if err != nil {
		log.Printf("Error saving participant in memory: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to save participant in memory")
	}

	return createdParticipant, nil
}
