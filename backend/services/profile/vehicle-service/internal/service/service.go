package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/ports"
)

type vehicleService struct {
	repo      ports.VehicleRepository
	publisher ports.EventPublisher
}

func NewVehicleService(repo ports.VehicleRepository, publisher ports.EventPublisher) ports.VehicleService {
	return &vehicleService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *vehicleService) Create(ctx context.Context, req dto.CreateVehicleRequest) (*dto.VehicleResponse, error) {
	now := time.Now().UTC()
	vehicle := &model.Vehicle{
		ID:              uuid.NewString(),
		DriverID:        req.DriverID,
		Make:            req.Make,
		Model:           req.Model,
		Year:            req.Year,
		Color:           req.Color,
		PlateNumber:     req.PlateNumber,
		VehicleType:     model.VehicleType(req.VehicleType),
		InsuranceExpiry: req.InsuranceExpiry,
		FitnessExpiry:   req.FitnessExpiry,
		Status:          model.VehicleStatusPending,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	if err := s.repo.Create(ctx, vehicle); err != nil {
		return nil, fmt.Errorf("create vehicle: %w", err)
	}

	_ = s.publisher.PublishVehicleRegistered(ctx, vehicle)

	return toVehicleResponse(vehicle), nil
}

func (s *vehicleService) GetByID(ctx context.Context, id string) (*dto.VehicleResponse, error) {
	vehicle, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find vehicle: %w", err)
	}
	if vehicle == nil {
		return nil, fmt.Errorf("vehicle not found")
	}
	return toVehicleResponse(vehicle), nil
}

func (s *vehicleService) GetByDriverID(ctx context.Context, driverID string) ([]dto.VehicleResponse, error) {
	vehicles, err := s.repo.FindByDriverID(ctx, driverID)
	if err != nil {
		return nil, fmt.Errorf("find vehicles by driver: %w", err)
	}
	return toVehicleResponseList(vehicles), nil
}

func (s *vehicleService) Update(ctx context.Context, id string, req dto.UpdateVehicleRequest) (*dto.VehicleResponse, error) {
	vehicle, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find vehicle: %w", err)
	}
	if vehicle == nil {
		return nil, fmt.Errorf("vehicle not found")
	}

	applyUpdates(vehicle, req)
	vehicle.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(ctx, vehicle); err != nil {
		return nil, fmt.Errorf("update vehicle: %w", err)
	}

	if model.VehicleStatus(req.Status) == model.VehicleStatusApproved {
		_ = s.publisher.PublishVehicleApproved(ctx, vehicle)
	}

	return toVehicleResponse(vehicle), nil
}

func (s *vehicleService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete vehicle: %w", err)
	}
	return nil
}

func (s *vehicleService) ListByType(ctx context.Context, vehicleType string) ([]dto.VehicleResponse, error) {
	vehicles, err := s.repo.ListByType(ctx, model.VehicleType(vehicleType))
	if err != nil {
		return nil, fmt.Errorf("list vehicles by type: %w", err)
	}
	return toVehicleResponseList(vehicles), nil
}

func applyUpdates(vehicle *model.Vehicle, req dto.UpdateVehicleRequest) {
	if req.Make != "" {
		vehicle.Make = req.Make
	}
	if req.Model != "" {
		vehicle.Model = req.Model
	}
	if req.Year != 0 {
		vehicle.Year = req.Year
	}
	if req.Color != "" {
		vehicle.Color = req.Color
	}
	if req.PlateNumber != "" {
		vehicle.PlateNumber = req.PlateNumber
	}
	if req.VehicleType != "" {
		vehicle.VehicleType = model.VehicleType(req.VehicleType)
	}
	if req.InsuranceExpiry != nil {
		vehicle.InsuranceExpiry = *req.InsuranceExpiry
	}
	if req.FitnessExpiry != nil {
		vehicle.FitnessExpiry = *req.FitnessExpiry
	}
	if req.Status != "" {
		vehicle.Status = model.VehicleStatus(req.Status)
	}
}

func toVehicleResponse(v *model.Vehicle) *dto.VehicleResponse {
	return &dto.VehicleResponse{
		ID:              v.ID,
		DriverID:        v.DriverID,
		Make:            v.Make,
		Model:           v.Model,
		Year:            v.Year,
		Color:           v.Color,
		PlateNumber:     v.PlateNumber,
		VehicleType:     string(v.VehicleType),
		InsuranceExpiry: v.InsuranceExpiry,
		FitnessExpiry:   v.FitnessExpiry,
		Status:          string(v.Status),
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}
}

func toVehicleResponseList(vehicles []*model.Vehicle) []dto.VehicleResponse {
	resp := make([]dto.VehicleResponse, len(vehicles))
	for i, v := range vehicles {
		resp[i] = *toVehicleResponse(v)
	}
	return resp
}
