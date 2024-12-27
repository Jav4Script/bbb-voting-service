package controllers

import (
	"net/http"

	"bbb-voting-service/internal/domain/dtos"
	"bbb-voting-service/internal/infrastructure/services"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	CaptchaService *services.CaptchaService
}

func NewCaptchaController(captchaService *services.CaptchaService) *CaptchaController {
	return &CaptchaController{CaptchaService: captchaService}
}

// GenerateCaptcha godoc
// @Summary Generate CAPTCHA
// @Description Generates a new CAPTCHA
// @Tags captcha
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "captcha_id and captcha_image"
// @Router /generate-captcha [get]
func (captcha_controller *CaptchaController) GenerateCaptcha(context *gin.Context) {
	captcha := captcha_controller.CaptchaService.GenerateCaptcha()
	context.JSON(http.StatusOK, structs.Map(captcha))
}

// ServeCaptcha godoc
// @Summary Serve CAPTCHA
// @Description Serves the CAPTCHA image
// @Tags captcha
// @Accept  json
// @Produce  json
// @Param captcha_id path string true "CAPTCHA ID"
// @Success 200 {object} map[string]interface{} "captcha_id and captcha_image"
// @Failure 404 {object} map[string]string "CAPTCHA not found"
// @Router /captcha/{captcha_id} [get]
func (captcha_controller *CaptchaController) ServeCaptcha(context *gin.Context) {
	captchaID := context.Param("captcha_id")
	captcha_controller.CaptchaService.ServeCaptcha(context.Writer, context.Request, captchaID)
}

// ValidateCaptcha godoc
// @Summary Validate CAPTCHA
// @Description Validates the CAPTCHA solution
// @Tags captcha
// @Accept  json
// @Produce  json
// @Param captcha body dtos.ValidateCaptchaDTO true "CAPTCHA ID and solution"
// @Success 200 {object} map[string]string "CAPTCHA validated successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 403 {object} map[string]string "Invalid CAPTCHA"
// @Router /validate-captcha [post]
func (captcha_controller *CaptchaController) ValidateCaptcha(context *gin.Context) {
	var request dtos.ValidateCaptchaDTO

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !captcha_controller.CaptchaService.ValidateCaptcha(request.CaptchaID, request.CaptchaSolution) {
		context.JSON(http.StatusForbidden, gin.H{"error": "Invalid CAPTCHA"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "CAPTCHA validated successfully"})
}
