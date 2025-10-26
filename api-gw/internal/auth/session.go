package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
	"time"
)

type Session struct {
	User      string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Store struct {
	mu       sync.RWMutex
	sessions map[string]Session
	ttl      time.Duration
	secret   string
}

func NewStore(ttl time.Duration, secret string) *Store {
	return &Store{
		sessions: make(map[string]Session),
		ttl:      ttl,
		secret:   secret,
	}
}

func randomToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (s *Store) Create(user string) (string, Session, error) {
	token, err := randomToken(32)
	if err != nil {
		return "", Session{}, err
	}
	now := time.Now()
	sess := Session{
		User:      user,
		CreatedAt: now,
		ExpiresAt: now.Add(s.ttl),
	}
	s.mu.Lock()
	s.sessions[token] = sess
	s.mu.Unlock()
	return token, sess, nil
}

func (s *Store) Get(token string) (Session, bool) {
	s.mu.RLock()
	sess, ok := s.sessions[token]
	s.mu.RUnlock()
	if !ok {
		return Session{}, false
	}
	if time.Now().After(sess.ExpiresAt) {
		// expire
		s.Delete(token)
		return Session{}, false
	}
	return sess, true
}

func (s *Store) Delete(token string) {
	s.mu.Lock()
	delete(s.sessions, token)
	s.mu.Unlock()
}

var ErrInvalidCreds = errors.New("invalid username or password")
