package ports

import "context"

type GatewayAuthorizeResult struct {
	GatewayID     string
	GatewayStatus string
}

type GatewayCaptureResult struct {
	GatewayStatus string
}

type GatewayRefundResult struct {
	GatewayRefundID string
	GatewayStatus   string
}

type PaymentGateway interface {
	Authorize(ctx context.Context, amount float64, currency, method string) (*GatewayAuthorizeResult, error)
	Capture(ctx context.Context, gatewayID string, amount float64) (*GatewayCaptureResult, error)
	Refund(ctx context.Context, gatewayID string, amount float64) (*GatewayRefundResult, error)
}
