package dtos

type CastVoteDTO struct {
	ParticipantID string `json:"participant_id" binding:"required"`
	VoterID       string `json:"voter_id" binding:"required"`
	IPAddress     string `json:"ip_address" binding:"required"`
	UserAgent     string `json:"user_agent" binding:"required"`
	Region        string `json:"region" binding:"required"`
	Device        string `json:"device" binding:"required"`
}
