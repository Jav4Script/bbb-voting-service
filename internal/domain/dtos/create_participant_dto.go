package dtos

// CreateParticipantDTO struct
type CreateParticipantDTO struct {
	Name string `json:"name" binding:"required"`
}
