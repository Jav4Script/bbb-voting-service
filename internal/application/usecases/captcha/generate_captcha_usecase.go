package captcha

import (
	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/services"
)

type GenerateCaptchaUsecase struct {
	CaptchaService services.CaptchaService
}

func NewGenerateCaptchaUsecase(captchaService services.CaptchaService) *GenerateCaptchaUsecase {
	return &GenerateCaptchaUsecase{CaptchaService: captchaService}
}

func (usecase *GenerateCaptchaUsecase) Execute() entities.Captcha {
	return usecase.CaptchaService.GenerateCaptcha()
}
