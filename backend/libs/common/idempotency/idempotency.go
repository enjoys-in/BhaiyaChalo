package idempotency

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Store is the interface for idempotency key storage.
type Store interface {
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, ttl time.Duration) error
}

// Key generates a deterministic idempotency key from parts.
func Key(parts ...string) string {
	h := sha256.New()
	for _, p := range parts {
		h.Write([]byte(p))
		h.Write([]byte("|"))
	}
	return hex.EncodeToString(h.Sum(nil))
}
