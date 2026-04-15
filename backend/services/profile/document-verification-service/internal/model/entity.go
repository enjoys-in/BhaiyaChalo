package model

import "time"

type OwnerType string

const (
	OwnerTypeDriver  OwnerType = "driver"
	OwnerTypeVehicle OwnerType = "vehicle"
)

type DocType string

const (
	DocTypeLicense   DocType = "license"
	DocTypeRC        DocType = "rc"
	DocTypeInsurance DocType = "insurance"
	DocTypePermits   DocType = "permits"
	DocTypeAadhar    DocType = "aadhar"
	DocTypePAN       DocType = "pan"
)

type DocumentStatus string

const (
	DocumentStatusPending  DocumentStatus = "pending"
	DocumentStatusApproved DocumentStatus = "approved"
	DocumentStatusRejected DocumentStatus = "rejected"
	DocumentStatusExpired  DocumentStatus = "expired"
)

type Document struct {
	ID         string         `json:"id" db:"id"`
	OwnerID    string         `json:"owner_id" db:"owner_id"`
	OwnerType  OwnerType      `json:"owner_type" db:"owner_type"`
	DocType    DocType        `json:"doc_type" db:"doc_type"`
	FileURL    string         `json:"file_url" db:"file_url"`
	Status     DocumentStatus `json:"status" db:"status"`
	ReviewerID string         `json:"reviewer_id,omitempty" db:"reviewer_id"`
	ReviewNote string         `json:"review_note,omitempty" db:"review_note"`
	ExpiresAt  *time.Time     `json:"expires_at,omitempty" db:"expires_at"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
}

type VerificationResult struct {
	DocumentID string         `json:"document_id"`
	Status     DocumentStatus `json:"status"`
	ReviewerID string         `json:"reviewer_id"`
	ReviewNote string         `json:"review_note"`
	VerifiedAt time.Time      `json:"verified_at"`
}
