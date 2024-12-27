package controllers

import (
	"net/http"

	usecase "bbb-voting-service/internal/application/usecases"
	"bbb-voting-service/internal/infrastructure/models"

	"github.com/gin-gonic/gin"
)

type VoteController struct {
	CastVoteUsecase *usecase.CastVoteUsecase
}

func NewVoteController(castVoteUseCase *usecase.CastVoteUsecase) *VoteController {
	return &VoteController{CastVoteUsecase: castVoteUseCase}
}

// CastVote godoc
// @Summary Cast Vote
// @Description Casts a vote for a participant
// @Tags vote
// @Accept  json
// @Produce  json
// @Param vote body models.VoteModel true "Vote details"
// @Success 200 {object} map[string]string
// @Router /votes [post]
func (controller *VoteController) CastVote(context *gin.Context) {
	var vote models.VoteModel
	if err := context.ShouldBindJSON(&vote); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainVote := models.ToDomainVote(vote)
	if err := controller.CastVoteUsecase.Execute(domainVote); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Vote cast successfully"})
}
