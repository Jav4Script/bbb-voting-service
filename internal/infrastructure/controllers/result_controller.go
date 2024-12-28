package controllers

import (
	"log"
	"net/http"

	usecase "bbb-voting-service/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

type ResultController struct {
	GetPartialResultsUsecase *usecase.GetPartialResultsUsecase
	GetFinalResultsUsecase   *usecase.GetFinalResultsUsecase
}

func NewResultController(partialResultsUseCase *usecase.GetPartialResultsUsecase, finalResultsUsecase *usecase.GetFinalResultsUsecase) *ResultController {
	return &ResultController{
		GetPartialResultsUsecase: partialResultsUseCase,
		GetFinalResultsUsecase:   finalResultsUsecase,
	}
}

// GetPartialResults godoc
// @Summary Get Partial Results
// @Description Retrieves partial voting results
// @Tags results
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]int
// @Failure 404 {object} map[string]string
// @Router /results/partial [get]
func (controller *ResultController) GetPartialResults(context *gin.Context) {
	results, err := controller.GetPartialResultsUsecase.Execute()
	if err != nil {
		log.Printf("Error retrieving partial results: %v", err)
		context.JSON(http.StatusNotFound, gin.H{"error": "Partial results not found"})
		return
	}

	context.JSON(http.StatusOK, results)
}

// GetFinalResults godoc
// @Summary Get Final Results
// @Description Retrieves final voting results
// @Tags results
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /results/final [get]
func (controller *ResultController) GetFinalResults(context *gin.Context) {
	results, err := controller.GetFinalResultsUsecase.Execute()
	if err != nil {
		log.Printf("Error retrieving final results: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get final results"})
		return
	}

	context.JSON(http.StatusOK, results)
}
