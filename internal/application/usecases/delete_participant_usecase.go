package usecase

import (
	repositories "bbb-voting-service/internal/domain/repositories"
)

type DeleteParticipantUsecase struct {
	ParticipantRepository repositories.ParticipantRepository
}

func NewDeleteParticipantUsecase(participantRepository repositories.ParticipantRepository) *DeleteParticipantUsecase {
	return &DeleteParticipantUsecase{ParticipantRepository: participantRepository}
}

func (usecase *DeleteParticipantUsecase) Execute(id string) error {
	participant, err := usecase.ParticipantRepository.FindByID(id)

	if err != nil {
		return err
	}

	err = usecase.ParticipantRepository.Delete(participant)

	if err != nil {
		return err
	}

	return nil
}
