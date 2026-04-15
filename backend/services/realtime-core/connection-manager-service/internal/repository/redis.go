package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/ports"
)

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) ports.ConnectionRepository {
	return &redisRepo{client: client}
}

func (r *redisRepo) Register(ctx context.Context, conn *model.Connection) error {
	data, err := json.Marshal(conn)
	if err != nil {
		return fmt.Errorf("marshal connection: %w", err)
	}

	pipe := r.client.Pipeline()
	connKey := constants.RedisConnectionPrefix + conn.ID
	pipe.Set(ctx, connKey, data, 24*time.Hour)
	pipe.Set(ctx, constants.RedisUserConnection+conn.UserID, conn.ID, 24*time.Hour)
	pipe.SAdd(ctx, constants.RedisNodeConnections+conn.ServerNode, conn.ID)
	pipe.Incr(ctx, constants.RedisNodeHealth+conn.ServerNode+":count")

	_, err = pipe.Exec(ctx)
	return err
}

func (r *redisRepo) Remove(ctx context.Context, connectionID string) error {
	connKey := constants.RedisConnectionPrefix + connectionID

	data, err := r.client.Get(ctx, connKey).Bytes()
	if err != nil {
		return fmt.Errorf("connection not found: %w", err)
	}

	var conn model.Connection
	if err := json.Unmarshal(data, &conn); err != nil {
		return fmt.Errorf("unmarshal connection: %w", err)
	}

	pipe := r.client.Pipeline()
	pipe.Del(ctx, connKey)
	pipe.Del(ctx, constants.RedisUserConnection+conn.UserID)
	pipe.SRem(ctx, constants.RedisNodeConnections+conn.ServerNode, connectionID)
	pipe.Decr(ctx, constants.RedisNodeHealth+conn.ServerNode+":count")

	_, err = pipe.Exec(ctx)
	return err
}

func (r *redisRepo) FindByUserID(ctx context.Context, userID string) (*model.Connection, error) {
	connID, err := r.client.Get(ctx, constants.RedisUserConnection+userID).Result()
	if err != nil {
		return nil, fmt.Errorf("user connection not found: %w", err)
	}

	data, err := r.client.Get(ctx, constants.RedisConnectionPrefix+connID).Bytes()
	if err != nil {
		return nil, fmt.Errorf("connection data not found: %w", err)
	}

	var conn model.Connection
	if err := json.Unmarshal(data, &conn); err != nil {
		return nil, fmt.Errorf("unmarshal connection: %w", err)
	}
	return &conn, nil
}

func (r *redisRepo) FindByNode(ctx context.Context, nodeID string) ([]*model.Connection, error) {
	ids, err := r.client.SMembers(ctx, constants.RedisNodeConnections+nodeID).Result()
	if err != nil {
		return nil, fmt.Errorf("get node connections: %w", err)
	}

	if len(ids) == 0 {
		return nil, nil
	}

	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = constants.RedisConnectionPrefix + id
	}

	vals, err := r.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("mget connections: %w", err)
	}

	conns := make([]*model.Connection, 0, len(vals))
	for _, v := range vals {
		if v == nil {
			continue
		}
		var c model.Connection
		if err := json.Unmarshal([]byte(v.(string)), &c); err == nil {
			conns = append(conns, &c)
		}
	}
	return conns, nil
}

func (r *redisRepo) GetNodeStatus(ctx context.Context, nodeID string) (*model.NodeStatus, error) {
	count, err := r.client.Get(ctx, constants.RedisNodeHealth+nodeID+":count").Int64()
	if err == redis.Nil {
		count = 0
	} else if err != nil {
		return nil, fmt.Errorf("get node count: %w", err)
	}

	return &model.NodeStatus{
		NodeID:            nodeID,
		ActiveConnections: count,
		HealthyAt:         time.Now().UTC(),
	}, nil
}
