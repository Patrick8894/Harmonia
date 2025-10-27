package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

// SessionStore is the abstraction used by controllers/middleware.
type SessionStore interface {
	Create(user string) (token string, err error)
	Get(token string) (user string, ok bool)
	Delete(token string)
}

// -------- In-Memory implementation (dev/fallback) --------

type memoryStore struct {
	mu  sync.RWMutex
	ttl time.Duration
	// token -> (user, expiry)
	data map[string]struct {
		user   string
		expiry time.Time
	}
}

func NewMemoryStore(ttl time.Duration) SessionStore {
	return &memoryStore{
		ttl: ttl,
		data: make(map[string]struct {
			user   string
			expiry time.Time
		}),
	}
}

func randToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (m *memoryStore) Create(user string) (string, error) {
	tok, err := randToken(32)
	if err != nil {
		return "", err
	}
	m.mu.Lock()
	m.data[tok] = struct {
		user   string
		expiry time.Time
	}{user: user, expiry: time.Now().Add(m.ttl)}
	m.mu.Unlock()
	return tok, nil
}

func (m *memoryStore) Get(tok string) (string, bool) {
	m.mu.RLock()
	rec, ok := m.data[tok]
	m.mu.RUnlock()
	if !ok || time.Now().After(rec.expiry) {
		if ok {
			m.Delete(tok)
		}
		return "", false
	}
	return rec.user, true
}

func (m *memoryStore) Delete(tok string) {
	m.mu.Lock()
	delete(m.data, tok)
	m.mu.Unlock()
}
