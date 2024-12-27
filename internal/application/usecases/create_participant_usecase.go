package usecase

import (
	domain "bbb-voting-service/internal/domain/entities"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type CreateParticipantUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
}

func NewCreateParticipantUsecase(participantRepository repositories.ParticipantRepository) *CreateParticipantUsecase {
	return &CreateParticipantUsecase{ParticipantRepository: participantRepository}
}

func (usecase *CreateParticipantUsecase) Execute(participant domain.Participant) (*domain.Participant, error) {
	existingParticipant, err := usecase.ParticipantRepository.FindByName(participant.Name)
	if err == nil {
		return &existingParticipant, nil
	}

	participant, err = usecase.ParticipantRepository.Save(participant)
	if err != nil {
		return nil, err
	}

	return &participant, nil
}
