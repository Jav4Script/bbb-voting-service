package dtos

// CreateParticipantDTO struct
type CreateParticipantDTO struct {
	Name   string `json:"name" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	Gender string `json:"gender" binding:"required"`
}
