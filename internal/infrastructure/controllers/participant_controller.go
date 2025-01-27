package controllers

import (
	"log"
	"net/http"

	"bbb-voting-service/internal/application/usecases/participants"
	"bbb-voting-service/internal/domain/dtos"
	"bbb-voting-service/internal/infrastructure/mappers"

	"github.com/gin-gonic/gin"
)

type ParticipantController struct {
	GetParticipantsUsecase   *participants.GetParticipantsUsecase
	GetParticipantUsecase    *participants.GetParticipantUsecase
	CreateParticipantUsecase *participants.CreateParticipantUsecase
	DeleteParticipantUsecase *participants.DeleteParticipantUsecase
}

func NewParticipantController(
	getParticipantsUsecase *participants.GetParticipantsUsecase,
	getParticipantUsecase *participants.GetParticipantUsecase,
	createParticipantUsecase *participants.CreateParticipantUsecase,
	deleteParticipantUsecase *participants.DeleteParticipantUsecase,
) *ParticipantController {
	return &ParticipantController{
		GetParticipantsUsecase:   getParticipantsUsecase,
		GetParticipantUsecase:    getParticipantUsecase,
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
// @Router /v1/participants [get]
func (controller *ParticipantController) GetParticipants(context *gin.Context) {
	participants, err := controller.GetParticipantsUsecase.Execute(context)
	if err != nil {
		log.Printf("Error retrieving participants: %v", err)
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, participants)
}

// GetParticipant godoc
// @Summary Get Participant
// @Description Retrieves a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param id path string true "Participant ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/participants/{id} [get]
func (controller *ParticipantController) GetParticipant(context *gin.Context) {
	id := context.Param("id")
	participant, err := controller.GetParticipantUsecase.Execute(context, id)
	if err != nil {
		log.Printf("Error retrieving participant: %v", err)
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, participant)
}

// CreateParticipant godoc
// @Summary Create Participant
// @Description Creates a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param participant body dtos.CreateParticipantDTO true "Participant"
// @Success 201 {object} map[string]interface{}
// @Router /v1/participants [post]
func (controller *ParticipantController) CreateParticipant(context *gin.Context) {
	var dto dtos.CreateParticipantDTO
	if err := context.ShouldBindJSON(&dto); err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	participantEntity := mappers.FromCreateParticipantDTO(dto)

	participant, err := controller.CreateParticipantUsecase.Execute(context, participantEntity)
	if err != nil {
		log.Printf("Error creating participant: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create participant"})
		return
	}

	context.JSON(http.StatusCreated, participant)
}

// DeleteParticipant godoc
// @Summary Delete Participant
// @Description Deletes a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param id path string true "Participant ID"
// @Success 204
// @Router /v1/participants/{id} [delete]
func (controller *ParticipantController) DeleteParticipant(context *gin.Context) {
	id := context.Param("id")
	err := controller.DeleteParticipantUsecase.Execute(context, id)
	if err != nil {
		log.Printf("Error deleting participant: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete participant"})
		return
	}

	context.Status(http.StatusNoContent)
}
