package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/ports"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisLocationCache(client *redis.Client) ports.LocationCache {
	return &redisCache{client: client}
}

func (r *redisCache) Set(ctx context.Context, driverID string, lat, lng float64, ttlSeconds int) error {
	pipe := r.client.Pipeline()

	pipe.GeoAdd(ctx, constants.RedisGeoKey, &redis.GeoLocation{
		Name:      driverID,
		Longitude: lng,
		Latitude:  lat,
	})

	locData, _ := json.Marshal(map[string]float64{"lat": lat, "lng": lng})
	key := constants.RedisLocationPrefix + driverID
	pipe.Set(ctx, key, locData, 0)

	_, err := pipe.Exec(ctx)
	return err
}

func (r *redisCache) Get(ctx context.Context, driverID string) (float64, float64, error) {
	positions, err := r.client.GeoPos(ctx, constants.RedisGeoKey, driverID).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("geo pos: %w", err)
	}
	if len(positions) == 0 || positions[0] == nil {
		return 0, 0, fmt.Errorf("location not found for driver: %s", driverID)
	}
	return positions[0].Latitude, positions[0].Longitude, nil
}

func (r *redisCache) GeoRadius(ctx context.Context, lat, lng, radiusKM float64) ([]ports.GeoLocation, error) {
	results, err := r.client.GeoSearch(ctx, constants.RedisGeoKey, &redis.GeoSearchQuery{
		Longitude:  lng,
		Latitude:   lat,
		Radius:     radiusKM,
		RadiusUnit: "km",
		Sort:       "ASC",
	}).Result()
	if err != nil {
		return nil, fmt.Errorf("geo search: %w", err)
	}

	locations := make([]ports.GeoLocation, 0, len(results))
	for _, name := range results {
		pos, err := r.client.GeoPos(ctx, constants.RedisGeoKey, name).Result()
		if err != nil || len(pos) == 0 || pos[0] == nil {
			continue
		}
		dist, _ := r.client.GeoDist(ctx, constants.RedisGeoKey, name, name, "km").Result()
		locations = append(locations, ports.GeoLocation{
			DriverID: name,
			Lat:      pos[0].Latitude,
			Lng:      pos[0].Longitude,
			DistKM:   dist,
		})
	}
	return locations, nil
}
