package dto

type UploadDocumentRequest struct {
	OwnerID   string `json:"owner_id" validate:"required"`
	OwnerType string `json:"owner_type" validate:"required,oneof=driver vehicle"`
	DocType   string `json:"doc_type" validate:"required,oneof=license rc insurance permits aadhar pan"`
	FileURL   string `json:"file_url" validate:"required"`
	ExpiresAt string `json:"expires_at,omitempty"`
}

type ReviewDocumentRequest struct {
	DocumentID string `json:"document_id" validate:"required"`
	Status     string `json:"status" validate:"required,oneof=approved rejected"`
	ReviewerID string `json:"reviewer_id" validate:"required"`
	ReviewNote string `json:"review_note,omitempty"`
}

type ListDocumentsRequest struct {
	OwnerID   string `json:"owner_id"`
	OwnerType string `json:"owner_type"`
	Status    string `json:"status"`
	Page      int    `json:"page"`
	PerPage   int    `json:"per_page"`
}
