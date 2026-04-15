package dto

type MatchResponse struct {
	DriverID string  `json:"driver_id"`
	Distance float64 `json:"distance"`
	ETA      int     `json:"eta"`
}

type CandidatesResponse struct {
	BookingID  string          `json:"booking_id"`
	Candidates []MatchResponse `json:"candidates"`
}
