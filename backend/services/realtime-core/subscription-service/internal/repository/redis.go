package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/ports"
)

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) ports.SubscriptionRepository {
	return &redisRepo{client: client}
}

func (r *redisRepo) Subscribe(ctx context.Context, sub *model.Subscription) error {
	data, err := json.Marshal(sub)
	if err != nil {
		return fmt.Errorf("marshal subscription: %w", err)
	}

	channelKey := constants.RedisChannelPrefix + string(sub.Channel) + ":" + sub.Topic

	pipe := r.client.Pipeline()
	pipe.Set(ctx, constants.RedisSubscriptionPrefix+sub.ID, data, 24*time.Hour)
	pipe.SAdd(ctx, constants.RedisUserSubscriptions+sub.UserID, sub.ID)
	pipe.SAdd(ctx, channelKey, sub.ID)
	pipe.Incr(ctx, constants.RedisChannelCount+string(sub.Channel))

	_, err = pipe.Exec(ctx)
	return err
}

func (r *redisRepo) Unsubscribe(ctx context.Context, subscriptionID string) error {
	subKey := constants.RedisSubscriptionPrefix + subscriptionID

	data, err := r.client.Get(ctx, subKey).Bytes()
	if err != nil {
		return fmt.Errorf("subscription not found: %w", err)
	}

	var sub model.Subscription
	if err := json.Unmarshal(data, &sub); err != nil {
		return fmt.Errorf("unmarshal subscription: %w", err)
	}

	channelKey := constants.RedisChannelPrefix + string(sub.Channel) + ":" + sub.Topic

	pipe := r.client.Pipeline()
	pipe.Del(ctx, subKey)
	pipe.SRem(ctx, constants.RedisUserSubscriptions+sub.UserID, subscriptionID)
	pipe.SRem(ctx, channelKey, subscriptionID)
	pipe.Decr(ctx, constants.RedisChannelCount+string(sub.Channel))

	_, err = pipe.Exec(ctx)
	return err
}

func (r *redisRepo) FindByUser(ctx context.Context, userID string) ([]*model.Subscription, error) {
	ids, err := r.client.SMembers(ctx, constants.RedisUserSubscriptions+userID).Result()
	if err != nil {
		return nil, fmt.Errorf("get user subscriptions: %w", err)
	}
	return r.fetchSubscriptions(ctx, ids)
}

func (r *redisRepo) FindByChannel(ctx context.Context, channel model.Channel, topic string) ([]*model.Subscription, error) {
	channelKey := constants.RedisChannelPrefix + string(channel) + ":" + topic
	ids, err := r.client.SMembers(ctx, channelKey).Result()
	if err != nil {
		return nil, fmt.Errorf("get channel subscriptions: %w", err)
	}
	return r.fetchSubscriptions(ctx, ids)
}

func (r *redisRepo) CountByChannel(ctx context.Context, channel model.Channel) (int64, error) {
	count, err := r.client.Get(ctx, constants.RedisChannelCount+string(channel)).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return count, err
}

func (r *redisRepo) fetchSubscriptions(ctx context.Context, ids []string) ([]*model.Subscription, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = constants.RedisSubscriptionPrefix + id
	}

	vals, err := r.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("mget subscriptions: %w", err)
	}

	subs := make([]*model.Subscription, 0, len(vals))
	for _, v := range vals {
		if v == nil {
			continue
		}
		var s model.Subscription
		if err := json.Unmarshal([]byte(v.(string)), &s); err == nil {
			subs = append(subs, &s)
		}
	}
	return subs, nil
}
