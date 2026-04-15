package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/ports"
)

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) ports.SessionRepository {
	return &redisRepo{client: client}
}

func (r *redisRepo) Register(ctx context.Context, session *model.SocketSession) error {
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("marshal session: %w", err)
	}

	pipe := r.client.Pipeline()
	sessionKey := constants.RedisSessionPrefix + session.ID
	pipe.Set(ctx, sessionKey, data, 24*time.Hour)
	pipe.SAdd(ctx, constants.RedisUserSessions+session.UserID, session.ID)
	pipe.SAdd(ctx, constants.RedisServerSessions+session.ServerID, session.ID)
	pipe.Incr(ctx, constants.RedisActiveCount)

	_, err = pipe.Exec(ctx)
	return err
}

func (r *redisRepo) Unregister(ctx context.Context, sessionID string) error {
	sessionKey := constants.RedisSessionPrefix + sessionID

	data, err := r.client.Get(ctx, sessionKey).Bytes()
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}

	var session model.SocketSession
	if err := json.Unmarshal(data, &session); err != nil {
		return fmt.Errorf("unmarshal session: %w", err)
	}

	now := time.Now().UTC()
	session.Active = false
	session.DisconnectedAt = &now

	updated, _ := json.Marshal(session)

	pipe := r.client.Pipeline()
	pipe.Set(ctx, sessionKey, updated, 1*time.Hour)
	pipe.SRem(ctx, constants.RedisUserSessions+session.UserID, sessionID)
	pipe.SRem(ctx, constants.RedisServerSessions+session.ServerID, sessionID)
	pipe.Decr(ctx, constants.RedisActiveCount)

	_, err = pipe.Exec(ctx)
	return err
}

func (r *redisRepo) FindByUserID(ctx context.Context, userID string) ([]*model.SocketSession, error) {
	ids, err := r.client.SMembers(ctx, constants.RedisUserSessions+userID).Result()
	if err != nil {
		return nil, fmt.Errorf("get user sessions: %w", err)
	}
	return r.fetchSessions(ctx, ids)
}

func (r *redisRepo) FindActiveByServer(ctx context.Context, serverID string) ([]*model.SocketSession, error) {
	ids, err := r.client.SMembers(ctx, constants.RedisServerSessions+serverID).Result()
	if err != nil {
		return nil, fmt.Errorf("get server sessions: %w", err)
	}
	return r.fetchSessions(ctx, ids)
}

func (r *redisRepo) CountActive(ctx context.Context) (int64, error) {
	count, err := r.client.Get(ctx, constants.RedisActiveCount).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return count, err
}

func (r *redisRepo) fetchSessions(ctx context.Context, ids []string) ([]*model.SocketSession, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = constants.RedisSessionPrefix + id
	}

	vals, err := r.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("mget sessions: %w", err)
	}

	sessions := make([]*model.SocketSession, 0, len(vals))
	for _, v := range vals {
		if v == nil {
			continue
		}
		var s model.SocketSession
		if err := json.Unmarshal([]byte(v.(string)), &s); err == nil {
			sessions = append(sessions, &s)
		}
	}
	return sessions, nil
}
