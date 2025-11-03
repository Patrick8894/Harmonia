package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	rdb   *redis.Client
	keyNS string
}

func NewRedisStore(rdb *redis.Client, ns string) *RedisStore {
	return &RedisStore{rdb: rdb, keyNS: ns}
}

func (r *RedisStore) full(key string) string {
	return r.keyNS + key
}

func (r *RedisStore) Get(ctx context.Context, key string, dst any) (bool, error) {
	s, err := r.rdb.Get(ctx, r.full(key)).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return json.Unmarshal([]byte(s), dst) == nil, nil
}

func (r *RedisStore) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return r.rdb.Set(ctx, r.full(key), b, ttl).Err()
}

func (r *RedisStore) Delete(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, r.full(key)).Err()
}
