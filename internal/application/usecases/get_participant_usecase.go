package usecase

import (
	domain "bbb-voting-service/internal/domain/entities"
	repositories "bbb-voting-service/internal/domain/repositories"
)

type GetParticipantUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
}

func NewGetParticipantUsecase(participantRepository repositories.ParticipantRepository) *GetParticipantUsecase {
	return &GetParticipantUsecase{ParticipantRepository: participantRepository}
}

func (usecase *GetParticipantUsecase) Execute(id string) (domain.Participant, error) {
	participant, err := usecase.ParticipantRepository.FindByID(id)

	if err != nil {
		return domain.Participant{}, err
	}

	return participant, nil
}
