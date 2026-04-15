package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/ports"
)

type driverService struct {
	repo      ports.DriverRepository
	publisher ports.EventPublisher
}

func NewDriverService(repo ports.DriverRepository, publisher ports.EventPublisher) ports.DriverService {
	return &driverService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *driverService) Create(ctx context.Context, req *dto.CreateDriverRequest) (*dto.DriverResponse, error) {
	now := time.Now().UTC()
	driver := &model.Driver{
		ID:             uuid.NewString(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Phone:          req.Phone,
		Email:          req.Email,
		LicenseNumber:  req.LicenseNumber,
		CityID:         req.CityID,
		Rating:         0,
		TotalTrips:     0,
		Status:         constants.StatusPending,
		OnboardingStep: 1,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repo.Create(ctx, driver); err != nil {
		return nil, err
	}

	_ = s.publisher.PublishDriverCreated(ctx, driver)

	return toDriverResponse(driver), nil
}

func (s *driverService) GetByID(ctx context.Context, id string) (*dto.DriverResponse, error) {
	driver, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toDriverResponse(driver), nil
}

func (s *driverService) GetByPhone(ctx context.Context, phone string) (*dto.DriverResponse, error) {
	driver, err := s.repo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	return toDriverResponse(driver), nil
}

func (s *driverService) Update(ctx context.Context, id string, req *dto.UpdateDriverRequest) (*dto.DriverResponse, error) {
	driver, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	oldStatus := driver.Status
	applyUpdate(driver, req)
	driver.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(ctx, driver); err != nil {
		return nil, err
	}

	_ = s.publisher.PublishDriverUpdated(ctx, driver)
	if req.Status != "" && req.Status != oldStatus {
		_ = s.publisher.PublishDriverStatusChanged(ctx, driver, oldStatus)
	}

	return toDriverResponse(driver), nil
}

func (s *driverService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *driverService) UpdatePreference(ctx context.Context, driverID string, req *dto.UpdatePreferenceRequest) (*dto.DriverPreferenceResponse, error) {
	pref, _ := s.repo.GetPreference(ctx, driverID)
	if pref == nil {
		pref = &model.DriverPreference{DriverID: driverID}
	}

	if req.AutoAccept != nil {
		pref.AutoAccept = *req.AutoAccept
	}
	if req.MaxDistance != nil {
		pref.MaxDistance = *req.MaxDistance
	}
	if req.PreferredZones != nil {
		pref.PreferredZones = req.PreferredZones
	}

	if err := s.repo.UpdatePreference(ctx, pref); err != nil {
		return nil, err
	}

	return toPreferenceResponse(pref), nil
}

func (s *driverService) GetPreference(ctx context.Context, driverID string) (*dto.DriverPreferenceResponse, error) {
	pref, err := s.repo.GetPreference(ctx, driverID)
	if err != nil {
		return nil, err
	}
	return toPreferenceResponse(pref), nil
}

func (s *driverService) ListByCityID(ctx context.Context, cityID string) ([]*dto.DriverResponse, error) {
	drivers, err := s.repo.ListByCityID(ctx, cityID)
	if err != nil {
		return nil, err
	}

	resp := make([]*dto.DriverResponse, len(drivers))
	for i, d := range drivers {
		resp[i] = toDriverResponse(d)
	}
	return resp, nil
}

func applyUpdate(driver *model.Driver, req *dto.UpdateDriverRequest) {
	if req.FirstName != "" {
		driver.FirstName = req.FirstName
	}
	if req.LastName != "" {
		driver.LastName = req.LastName
	}
	if req.Email != "" {
		driver.Email = req.Email
	}
	if req.AvatarURL != "" {
		driver.AvatarURL = req.AvatarURL
	}
	if req.LicenseNumber != "" {
		driver.LicenseNumber = req.LicenseNumber
	}
	if req.CityID != "" {
		driver.CityID = req.CityID
	}
	if req.Status != "" {
		driver.Status = req.Status
	}
}

func toDriverResponse(d *model.Driver) *dto.DriverResponse {
	return &dto.DriverResponse{
		ID:             d.ID,
		FirstName:      d.FirstName,
		LastName:       d.LastName,
		Phone:          d.Phone,
		Email:          d.Email,
		AvatarURL:      d.AvatarURL,
		LicenseNumber:  d.LicenseNumber,
		CityID:         d.CityID,
		Rating:         d.Rating,
		TotalTrips:     d.TotalTrips,
		Status:         d.Status,
		OnboardingStep: d.OnboardingStep,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
	}
}

func toPreferenceResponse(p *model.DriverPreference) *dto.DriverPreferenceResponse {
	zones := p.PreferredZones
	if zones == nil {
		zones = []string{}
	}
	return &dto.DriverPreferenceResponse{
		DriverID:       p.DriverID,
		AutoAccept:     p.AutoAccept,
		MaxDistance:    p.MaxDistance,
		PreferredZones: zones,
	}
}
