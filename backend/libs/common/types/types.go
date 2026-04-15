package types

import "time"

// Pagination holds pagination request params.
type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

// DefaultPagination returns a safe default.
func DefaultPagination() Pagination {
	return Pagination{Page: 1, PerPage: 20}
}

// Coordinate represents a GPS point.
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// CityID is a typed city identifier used for partitioning.
type CityID string

// Timestamps embeds standard created/updated fields.
type Timestamps struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
