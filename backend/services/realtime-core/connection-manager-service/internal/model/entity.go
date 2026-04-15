package model

import "time"

type Connection struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	ServerNode string    `json:"server_node"`
	Protocol   string    `json:"protocol"`
	CreatedAt  time.Time `json:"created_at"`
}

type NodeStatus struct {
	NodeID            string    `json:"node_id"`
	ActiveConnections int64     `json:"active_connections"`
	Capacity          int       `json:"capacity"`
	HealthyAt         time.Time `json:"healthy_at"`
}
