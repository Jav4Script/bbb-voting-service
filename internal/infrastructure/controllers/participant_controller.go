package controllers

import (
	"net/http"

	usecase "bbb-voting-service/internal/application/usecases"
	dtos "bbb-voting-service/internal/domain/dtos"
	entities "bbb-voting-service/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

type ParticipantController struct {
	GetParticipantsUsecase   *usecase.GetParticipantsUsecase
	GetParticipantUsecase    *usecase.GetParticipantUsecase
	CreateParticipantUsecase *usecase.CreateParticipantUsecase
	DeleteParticipantUsecase *usecase.DeleteParticipantUsecase
}

func NewParticipantController(participantsUsecase *usecase.GetParticipantsUsecase, participantUsecase *usecase.GetParticipantUsecase, createParticipantUsecase *usecase.CreateParticipantUsecase, deleteParticipantUsecase *usecase.DeleteParticipantUsecase) *ParticipantController {
	return &ParticipantController{
		GetParticipantsUsecase:   participantsUsecase,
		GetParticipantUsecase:    participantUsecase,
		CreateParticipantUsecase: createParticipantUsecase,
		DeleteParticipantUsecase: deleteParticipantUsecase,
	}
}

// GetParticipants godoc
// @Summary Get Participants
// @Description Retrieves all participants
// @Tags participants
// @Accept  json
// @Produce  json
// @Success 200 {object} []map[string]interface{}
// @Router /participants [get]
func (controller *ParticipantController) GetParticipants(context *gin.Context) {
	participants, err := controller.GetParticipantsUsecase.Execute()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	participantMaps := make([]map[string]interface{}, len(participants))
	for i, participant := range participants {
		participantMaps[i] = map[string]interface{}{
			"id":    participant.ID,
			"name":  participant.Name,
			"votes": participant.Votes,
		}
	}

	context.JSON(http.StatusOK, participantMaps)
}

// GetParticipant godoc
// @Summary Get Participant
// @Description Retrieves a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param id path string true "Participant ID"
// @Success 200 {object} map[string]interface{}
// @Router /participants/{id} [get]
func (controller *ParticipantController) GetParticipant(context *gin.Context) {
	id := context.Param("id")
	participant, err := controller.GetParticipantUsecase.Execute(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	participantMap := map[string]interface{}{
		"id":    participant.ID,
		"name":  participant.Name,
		"votes": participant.Votes,
	}

	context.JSON(http.StatusOK, participantMap)
}

// CreateParticipant godoc
// @Summary Create Participant
// @Description Creates a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param participant body dtos.CreateParticipantDTO true "Participant"
// @Success 201 {object} map[string]interface{}
// @Router /participants [post]
func (controller *ParticipantController) CreateParticipant(context *gin.Context) {
	var participantDTO dtos.CreateParticipantDTO
	if err := context.ShouldBindJSON(&participantDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	participant := entities.Participant{
		Name: participantDTO.Name,
	}

	createdParticipant, err := controller.CreateParticipantUsecase.Execute(participant)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	participantMap := map[string]interface{}{
		"id":    createdParticipant.ID,
		"name":  createdParticipant.Name,
		"votes": createdParticipant.Votes,
	}

	context.JSON(http.StatusCreated, participantMap)
}

// DeleteParticipant godoc
// @Summary Delete Participant
// @Description Deletes a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param id path string true "Participant ID"
// @Success 204
// @Router /participants/{id} [delete]
func (controller *ParticipantController) DeleteParticipant(context *gin.Context) {
	id := context.Param("id")
	err := controller.DeleteParticipantUsecase.Execute(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}
