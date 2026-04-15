package dto

type CreateDriverRequest struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	LicenseNumber string `json:"license_number"`
	CityID        string `json:"city_id"`
}

type UpdateDriverRequest struct {
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Email         string `json:"email,omitempty"`
	AvatarURL     string `json:"avatar_url,omitempty"`
	LicenseNumber string `json:"license_number,omitempty"`
	CityID        string `json:"city_id,omitempty"`
	Status        string `json:"status,omitempty"`
}

type UpdateLocationRequest struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Heading float64 `json:"heading"`
	Speed   float64 `json:"speed"`
}

type UpdatePreferenceRequest struct {
	AutoAccept     *bool    `json:"auto_accept,omitempty"`
	MaxDistance    *float64 `json:"max_distance,omitempty"`
	PreferredZones []string `json:"preferred_zones,omitempty"`
}
