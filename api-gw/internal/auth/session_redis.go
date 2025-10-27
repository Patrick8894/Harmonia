// internal/auth/session_redis.go
package auth

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	rdb   *redis.Client
	ttl   time.Duration
	keyNS string
}

func NewRedisStore(rdb *redis.Client, ttl time.Duration, ns string) *RedisStore {
	return &RedisStore{rdb: rdb, ttl: ttl, keyNS: ns}
}

func (s *RedisStore) Create(user string) (string, error) {
	tok, err := randToken(32)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	if err := s.rdb.Set(ctx, s.keyNS+tok, user, s.ttl).Err(); err != nil {
		return "", err
	}
	return tok, nil
}
func (s *RedisStore) Get(tok string) (string, bool) {
	v, err := s.rdb.Get(context.Background(), s.keyNS+tok).Result()
	if err != nil {
		return "", false
	}
	return v, true
}
func (s *RedisStore) Delete(tok string) {
	_ = s.rdb.Del(context.Background(), s.keyNS+tok).Err()
}
