package controllers

import (
	"log"
	"net/http"

	"bbb-voting-service/internal/application/usecases/captcha"
	"bbb-voting-service/internal/domain/dtos"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	GenerateCaptchaUsecase      *captcha.GenerateCaptchaUsecase
	ValidateCaptchaUsecase      *captcha.ValidateCaptchaUsecase
	ValidateCaptchaTokenUsecase *captcha.ValidateCaptchaTokenUsecase
	ServeCaptchaUsecase         *captcha.ServeCaptchaUsecase
}

func NewCaptchaController(
	generateCaptchaUsecase *captcha.GenerateCaptchaUsecase,
	validateCaptchaUsecase *captcha.ValidateCaptchaUsecase,
	validateCaptchaTokenUsecase *captcha.ValidateCaptchaTokenUsecase,
	serveCaptchaUsecase *captcha.ServeCaptchaUsecase,
) *CaptchaController {
	return &CaptchaController{
		GenerateCaptchaUsecase:      generateCaptchaUsecase,
		ValidateCaptchaUsecase:      validateCaptchaUsecase,
		ValidateCaptchaTokenUsecase: validateCaptchaTokenUsecase,
		ServeCaptchaUsecase:         serveCaptchaUsecase,
	}
}

// GenerateCaptcha godoc
// @Summary Generate CAPTCHA
// @Description Generates a new CAPTCHA
// @Tags captcha
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "captcha_id and captcha_image"
// @Router /v1/generate-captcha [get]
func (captchaController *CaptchaController) GenerateCaptcha(context *gin.Context) {
	captcha := captchaController.GenerateCaptchaUsecase.Execute()
	context.JSON(http.StatusOK, captcha)
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
// @Router /v1/captcha/{captcha_id} [get]
func (captchaController *CaptchaController) ServeCaptcha(context *gin.Context) {
	captchaID := context.Param("captcha_id")
	captchaController.ServeCaptchaUsecase.Execute(context.Writer, context.Request, captchaID)
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
// @Router /v1/validate-captcha [post]
func (captchaController *CaptchaController) ValidateCaptcha(context *gin.Context) {
	var request dtos.ValidateCaptchaDTO

	if err := context.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, valid := captchaController.ValidateCaptchaUsecase.Execute(request.CaptchaID, request.CaptchaSolution)
	if !valid {
		log.Printf("Invalid CAPTCHA for ID: %s", request.CaptchaID)
		context.JSON(http.StatusForbidden, gin.H{"error": "Invalid CAPTCHA"})
		return
	}

	context.Header("X-Captcha-Token", token)
	context.Header("Access-Control-Expose-Headers", "X-Captcha-Token")
	context.JSON(http.StatusOK, gin.H{"message": "CAPTCHA validated successfully"})
}
