package usecase

import (
	domain "bbb-voting-service/internal/domain/entities"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type GetParticipantsUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
}

func NewGetParticipantsUsecase(participantRepository repositories.ParticipantRepository) *GetParticipantsUsecase {
	return &GetParticipantsUsecase{ParticipantRepository: participantRepository}
}

func (usecase *GetParticipantsUsecase) Execute() ([]domain.Participant, error) {
	participants, err := usecase.ParticipantRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return participants, nil
}
