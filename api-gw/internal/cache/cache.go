package cache

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Store interface {
	Get(ctx context.Context, key string, dst any) (bool, error)
	Set(ctx context.Context, key string, val any, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

func Key(prefix string, input any) string {
	b, _ := json.Marshal(input)
	sum := sha256.Sum256(b)
	return prefix + ":" + hex.EncodeToString(sum[:])
}
