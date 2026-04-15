package dto

import "time"

type DocumentResponse struct {
	ID         string     `json:"id"`
	OwnerID    string     `json:"owner_id"`
	OwnerType  string     `json:"owner_type"`
	DocType    string     `json:"doc_type"`
	FileURL    string     `json:"file_url"`
	Status     string     `json:"status"`
	ReviewerID string     `json:"reviewer_id,omitempty"`
	ReviewNote string     `json:"review_note,omitempty"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type VerificationResponse struct {
	DocumentID string    `json:"document_id"`
	Status     string    `json:"status"`
	ReviewerID string    `json:"reviewer_id"`
	ReviewNote string    `json:"review_note,omitempty"`
	VerifiedAt time.Time `json:"verified_at"`
}
