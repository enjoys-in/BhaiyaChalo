package client

type AuthClient struct {
	Addr string
}

func NewAuthClient(addr string) *AuthClient {
	return &AuthClient{Addr: addr}
}
