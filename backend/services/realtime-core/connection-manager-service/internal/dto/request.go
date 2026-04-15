package dto

type RegisterConnectionRequest struct {
	UserID     string `json:"user_id" validate:"required"`
	ServerNode string `json:"server_node" validate:"required"`
	Protocol   string `json:"protocol" validate:"required,oneof=websocket sse grpc"`
}

type LocateUserRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
