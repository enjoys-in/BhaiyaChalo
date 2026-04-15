package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/ports"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	SAdd(ctx context.Context, key string, members ...interface{}) error
	SRem(ctx context.Context, key string, members ...interface{}) error
	SCard(ctx context.Context, key string) (int64, error)
}

type redisCache struct {
	client RedisClient
}

func NewRedisCache(client RedisClient) ports.AvailabilityCache {
	return &redisCache{client: client}
}

func (r *redisCache) Set(ctx context.Context, avail *model.DriverAvailability) error {
	data, err := json.Marshal(avail)
	if err != nil {
		return fmt.Errorf("marshal availability: %w", err)
	}

	key := driverKey(avail.DriverID)
	if err := r.client.Set(ctx, key, data, constants.AvailabilityTTL); err != nil {
		return fmt.Errorf("redis set: %w", err)
	}

	setKey := citySetKey(avail.CityID, avail.VehicleType)
	if err := r.client.SAdd(ctx, setKey, avail.DriverID); err != nil {
		return fmt.Errorf("redis sadd: %w", err)
	}

	return nil
}

func (r *redisCache) Get(ctx context.Context, driverID string) (*model.DriverAvailability, error) {
	key := driverKey(driverID)
	data, err := r.client.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("redis get: %w", err)
	}
	if data == "" {
		return nil, nil
	}

	var avail model.DriverAvailability
	if err := json.Unmarshal([]byte(data), &avail); err != nil {
		return nil, fmt.Errorf("unmarshal availability: %w", err)
	}
	return &avail, nil
}

func (r *redisCache) Delete(ctx context.Context, driverID string) error {
	cached, err := r.Get(ctx, driverID)
	if err != nil {
		return err
	}
	if cached != nil {
		setKey := citySetKey(cached.CityID, cached.VehicleType)
		_ = r.client.SRem(ctx, setKey, driverID)
	}
	return r.client.Del(ctx, driverKey(driverID))
}

func (r *redisCache) CountByCity(ctx context.Context, cityID, vehicleType string) (int, error) {
	setKey := citySetKey(cityID, vehicleType)
	count, err := r.client.SCard(ctx, setKey)
	if err != nil {
		return 0, fmt.Errorf("redis scard: %w", err)
	}
	return int(count), nil
}

func driverKey(driverID string) string {
	return fmt.Sprintf("availability:driver:%s", driverID)
}

func citySetKey(cityID, vehicleType string) string {
	return fmt.Sprintf("availability:city:%s:type:%s", cityID, vehicleType)
}
