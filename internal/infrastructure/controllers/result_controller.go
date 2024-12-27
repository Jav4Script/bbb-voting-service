package controllers

import (
	"net/http"

	usecase "bbb-voting-service/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

type ResultController struct {
	GetPartialResultsUsecase *usecase.GetPartialResultsUsecase
	GetFinalResultsUsecase   *usecase.GetFinalResultsUsecase
}

func NewResultController(partialResultsUseCase *usecase.GetPartialResultsUsecase, finalResultsUseCase *usecase.GetFinalResultsUsecase) *ResultController {
	return &ResultController{
		GetPartialResultsUsecase: partialResultsUseCase,
		GetFinalResultsUsecase:   finalResultsUseCase,
	}
}

// GetPartialResults godoc
// @Summary Get Partial Results
// @Description Retrieves partial voting results
// @Tags results
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]int
// @Router /results/partial [get]
func (controller *ResultController) GetPartialResults(context *gin.Context) {
	results, err := controller.GetPartialResultsUsecase.Execute()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, results)
}
