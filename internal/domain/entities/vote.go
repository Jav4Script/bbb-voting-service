package domain

import "time"

type Vote struct {
	ID                 string    `json:"id"`
	ParticipantID      string    `json:"participant_id"`
	VoterID            string    `json:"voter_id"`
	Timestamp          time.Time `json:"timestamp"`
	IPAddress          string    `json:"ip_address"`
	UserAgent          string    `json:"user_agent"`
	Region             string    `json:"region"`
	Device             string    `json:"device"`
	IsCaptchaCompleted bool      `json:"is_captcha_completed"`
}