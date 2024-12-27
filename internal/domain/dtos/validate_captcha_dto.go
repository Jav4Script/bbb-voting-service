package dtos

// ValidateCaptchaDTO struct
type ValidateCaptchaDTO struct {
	CaptchaID       string `json:"captcha_id" binding:"required"`
	CaptchaSolution string `json:"captcha_solution" binding:"required"`
}
