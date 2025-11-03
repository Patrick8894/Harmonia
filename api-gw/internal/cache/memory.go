package cache

import (
	"context"
	"encoding/json"
	"sync"
	"time"
)

type memEntry struct {
	raw    []byte
	expiry time.Time
}

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]memEntry
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string]memEntry)}
}

func (m *MemoryStore) Get(_ context.Context, key string, dst any) (bool, error) {
	m.mu.RLock()
	e, ok := m.data[key]
	m.mu.RUnlock()
	if !ok || time.Now().After(e.expiry) {
		if ok {
			m.Delete(context.Background(), key)
		}
		return false, nil
	}
	return json.Unmarshal(e.raw, dst) == nil, nil
}

func (m *MemoryStore) Set(_ context.Context, key string, val any, ttl time.Duration) error {
	raw, err := json.Marshal(val)
	if err != nil {
		return err
	}
	m.mu.Lock()
	m.data[key] = memEntry{raw: raw, expiry: time.Now().Add(ttl)}
	m.mu.Unlock()
	return nil
}

func (m *MemoryStore) Delete(_ context.Context, key string) error {
	m.mu.Lock()
	delete(m.data, key)
	m.mu.Unlock()
	return nil
}
