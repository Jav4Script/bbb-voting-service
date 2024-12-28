package domain

type Vote struct {
	ID            string `json:"id"`
	ParticipantID string `json:"participant_id"`
	VoterID       string `json:"voter_id"`
	IPAddress     string `json:"ip_address"`
	UserAgent     string `json:"user_agent"`
	Region        string `json:"region"`
	Device        string `json:"device"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
