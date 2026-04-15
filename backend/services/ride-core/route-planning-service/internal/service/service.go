package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/event"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/ports"
)

type routeService struct {
	repo     ports.RouteRepository
	engine   ports.RoutingEngine
	producer *event.Producer
}

func NewRouteService(repo ports.RouteRepository, engine ports.RoutingEngine, producer *event.Producer) ports.RouteService {
	return &routeService{
		repo:     repo,
		engine:   engine,
		producer: producer,
	}
}

func (s *routeService) PlanRoute(ctx context.Context, req dto.PlanRouteRequest) (*model.Route, error) {
	waypoints := buildWaypoints(req)

	route, err := s.engine.Route(ctx, waypoints)
	if err != nil {
		return nil, err
	}

	route.ID = uuid.New().String()
	route.BookingID = req.BookingID
	route.CreatedAt = time.Now().UTC()

	if err := s.repo.Save(ctx, route); err != nil {
		return nil, err
	}

	if err := s.producer.PublishRoutePlanned(ctx, route); err != nil {
		return nil, err
	}

	return route, nil
}

func (s *routeService) GetRoute(ctx context.Context, bookingID string) (*model.Route, error) {
	return s.repo.FindByBookingID(ctx, bookingID)
}

func buildWaypoints(req dto.PlanRouteRequest) []model.Waypoint {
	waypoints := make([]model.Waypoint, 0, len(req.Waypoints)+2)

	waypoints = append(waypoints, model.Waypoint{
		Lat:   req.Origin.Lat,
		Lng:   req.Origin.Lng,
		Order: 0,
		Label: "origin",
	})

	for i, wp := range req.Waypoints {
		waypoints = append(waypoints, model.Waypoint{
			Lat:   wp.Lat,
			Lng:   wp.Lng,
			Order: i + 1,
			Label: "waypoint",
		})
	}

	waypoints = append(waypoints, model.Waypoint{
		Lat:   req.Destination.Lat,
		Lng:   req.Destination.Lng,
		Order: len(waypoints),
		Label: "destination",
	})

	return waypoints
}
