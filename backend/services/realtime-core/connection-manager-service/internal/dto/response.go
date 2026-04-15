package dto

import "time"

type ConnectionResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	ServerNode string    `json:"server_node"`
	Protocol   string    `json:"protocol"`
	CreatedAt  time.Time `json:"created_at"`
}

type NodeStatusResponse struct {
	NodeID            string    `json:"node_id"`
	ActiveConnections int64     `json:"active_connections"`
	Capacity          int       `json:"capacity"`
	Healthy           bool      `json:"healthy"`
	HealthyAt         time.Time `json:"healthy_at"`
}

type UserLocationResponse struct {
	UserID     string `json:"user_id"`
	ServerNode string `json:"server_node"`
	Connected  bool   `json:"connected"`
}
