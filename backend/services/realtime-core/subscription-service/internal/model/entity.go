package model

import "time"

type Channel string

const (
	ChannelTripTracking   Channel = "trip_tracking"
	ChannelDriverLocation Channel = "driver_location"
	ChannelNotifications  Channel = "notifications"
	ChannelSurgeUpdates   Channel = "surge_updates"
)

type Subscription struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Channel   Channel   `json:"channel"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}
