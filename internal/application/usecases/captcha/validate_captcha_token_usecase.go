package captcha

import (
	"bbb-voting-service/internal/domain/services"
)

type ValidateCaptchaTokenUsecase struct {
	CaptchaService services.CaptchaService
}

func NewValidateCaptchaTokenUsecase(captchaService services.CaptchaService) *ValidateCaptchaTokenUsecase {
	return &ValidateCaptchaTokenUsecase{CaptchaService: captchaService}
}

func (usecase *ValidateCaptchaTokenUsecase) Execute(token string) bool {
	return usecase.CaptchaService.ValidateCaptchaToken(token)
}
