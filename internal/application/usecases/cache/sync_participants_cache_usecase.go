package cache

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/repositories"
)

type SyncParticipantsCacheUsecase struct {
	context                       context.Context
	ParticipantRepository         repositories.ParticipantRepository
	InMemoryParticipantRepository repositories.InMemoryParticipantRepository
}

func NewSyncParticipantsCacheUsecase(participantRepository repositories.ParticipantRepository, inMemoryParticipantRepository repositories.InMemoryParticipantRepository) *SyncParticipantsCacheUsecase {
	return &SyncParticipantsCacheUsecase{
		context:                       context.Background(),
		ParticipantRepository:         participantRepository,
		InMemoryParticipantRepository: inMemoryParticipantRepository,
	}
}

func (usecase *SyncParticipantsCacheUsecase) Execute(context context.Context) error {
	participants, err := usecase.ParticipantRepository.FindAll()
	if err != nil {
		log.Printf("Error retrieving participants: %v", err)
		return err
	}

	for _, participant := range participants {
		err = usecase.InMemoryParticipantRepository.Save(context, participant)
		if err != nil {
			log.Printf("Error saving participant to cache: %v", err)
		}
	}

	return nil
}
