package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/ports"
)

type geofenceService struct {
	repo      ports.GeofenceRepository
	publisher ports.EventPublisher
}

func NewGeofenceService(repo ports.GeofenceRepository, publisher ports.EventPublisher) ports.GeofenceService {
	return &geofenceService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *geofenceService) Create(ctx context.Context, req dto.CreateGeofenceRequest) (*dto.GeofenceResponse, error) {
	now := time.Now().UTC()
	fence := &model.Geofence{
		ID:        uuid.NewString(),
		CityID:    req.CityID,
		Name:      req.Name,
		Type:      model.FenceType(req.Type),
		Polygon:   toModelCoordinates(req.Polygon),
		CenterLat: req.CenterLat,
		CenterLng: req.CenterLng,
		RadiusKM:  req.RadiusKM,
		Active:    true,
		Metadata:  req.Metadata,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if fence.Metadata == nil {
		fence.Metadata = make(map[string]string)
	}

	if err := s.repo.Create(ctx, fence); err != nil {
		return nil, fmt.Errorf("create geofence: %w", err)
	}

	_ = s.publisher.PublishGeofenceCreated(ctx, fence)

	return toGeofenceResponse(fence), nil
}

func (s *geofenceService) Get(ctx context.Context, id string) (*dto.GeofenceResponse, error) {
	fence, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find geofence: %w", err)
	}
	if fence == nil {
		return nil, fmt.Errorf("geofence not found")
	}
	return toGeofenceResponse(fence), nil
}

func (s *geofenceService) List(ctx context.Context, cityID string) ([]dto.GeofenceResponse, error) {
	fences, err := s.repo.FindByCityID(ctx, cityID)
	if err != nil {
		return nil, fmt.Errorf("list geofences: %w", err)
	}
	return toGeofenceResponseList(fences), nil
}

func (s *geofenceService) Update(ctx context.Context, id string, req dto.UpdateGeofenceRequest) (*dto.GeofenceResponse, error) {
	fence, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find geofence: %w", err)
	}
	if fence == nil {
		return nil, fmt.Errorf("geofence not found")
	}

	applyGeofenceUpdates(fence, req)
	fence.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(ctx, fence); err != nil {
		return nil, fmt.Errorf("update geofence: %w", err)
	}

	_ = s.publisher.PublishGeofenceUpdated(ctx, fence)

	return toGeofenceResponse(fence), nil
}

func (s *geofenceService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete geofence: %w", err)
	}
	return nil
}

func (s *geofenceService) CheckPoint(ctx context.Context, req dto.PointInFenceRequest) (*dto.PointCheckResponse, error) {
	fences, err := s.repo.FindContaining(ctx, req.Lat, req.Lng)
	if err != nil {
		return nil, fmt.Errorf("find containing fences: %w", err)
	}

	matched := toGeofenceResponseList(fences)
	return &dto.PointCheckResponse{
		InFence:       len(matched) > 0,
		MatchedFences: matched,
	}, nil
}

func applyGeofenceUpdates(fence *model.Geofence, req dto.UpdateGeofenceRequest) {
	if req.Name != "" {
		fence.Name = req.Name
	}
	if req.Type != "" {
		fence.Type = model.FenceType(req.Type)
	}
	if len(req.Polygon) > 0 {
		fence.Polygon = toModelCoordinates(req.Polygon)
	}
	if req.CenterLat != nil {
		fence.CenterLat = *req.CenterLat
	}
	if req.CenterLng != nil {
		fence.CenterLng = *req.CenterLng
	}
	if req.RadiusKM != nil {
		fence.RadiusKM = *req.RadiusKM
	}
	if req.Active != nil {
		fence.Active = *req.Active
	}
	if req.Metadata != nil {
		fence.Metadata = req.Metadata
	}
}

func toModelCoordinates(coords []dto.CoordinateDTO) []model.Coordinate {
	result := make([]model.Coordinate, len(coords))
	for i, c := range coords {
		result[i] = model.Coordinate{Lat: c.Lat, Lng: c.Lng}
	}
	return result
}

func toDTOCoordinates(coords []model.Coordinate) []dto.CoordinateDTO {
	result := make([]dto.CoordinateDTO, len(coords))
	for i, c := range coords {
		result[i] = dto.CoordinateDTO{Lat: c.Lat, Lng: c.Lng}
	}
	return result
}

func toGeofenceResponse(f *model.Geofence) *dto.GeofenceResponse {
	return &dto.GeofenceResponse{
		ID:        f.ID,
		CityID:    f.CityID,
		Name:      f.Name,
		Type:      string(f.Type),
		Polygon:   toDTOCoordinates(f.Polygon),
		CenterLat: f.CenterLat,
		CenterLng: f.CenterLng,
		RadiusKM:  f.RadiusKM,
		Active:    f.Active,
		Metadata:  f.Metadata,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

func toGeofenceResponseList(fences []*model.Geofence) []dto.GeofenceResponse {
	resp := make([]dto.GeofenceResponse, len(fences))
	for i, f := range fences {
		resp[i] = *toGeofenceResponse(f)
	}
	return resp
}
