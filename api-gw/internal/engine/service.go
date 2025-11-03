package engine

import (
	"context"
	"time"

	eng "github.com/Patrick8894/harmonia/api-gw/gen/engine"
	"github.com/Patrick8894/harmonia/api-gw/internal/cache"
)

type Service struct {
	c   *Client
	kvs cache.Store
	ttl time.Duration
}

func NewService(c *Client, kvs cache.Store, ttl time.Duration) *Service {
	return &Service{c: c, kvs: kvs, ttl: ttl}
}

func (s *Service) Hello(ctx context.Context, name string) (string, error) {
	return s.c.hello(ctx, name)
}

func (s *Service) EstimatePi(ctx context.Context, samples int64) (*eng.PiReply, bool, error) {
	key := cache.Key("engine:pi", struct{ Samples int64 }{samples})
	var cached eng.PiReply
	if ok, _ := s.kvs.Get(ctx, key, &cached); ok {
		return &cached, true, nil
	}
	resp, err := s.c.estimatePi(ctx, samples)
	if err != nil {
		return nil, false, err
	}
	_ = s.kvs.Set(ctx, key, resp, s.ttl)
	return resp, false, nil
}

func (s *Service) MatMul(ctx context.Context, in MatMulDTO) (*eng.MatReply, bool, error) {
	key := cache.Key("engine:matmul", in)
	var cached eng.MatReply
	if ok, _ := s.kvs.Get(ctx, key, &cached); ok {
		return &cached, true, nil
	}
	a := &eng.Matrix{Rows: in.A.Rows, Cols: in.A.Cols, Data: in.A.Data}
	b := &eng.Matrix{Rows: in.B.Rows, Cols: in.B.Cols, Data: in.B.Data}
	resp, err := s.c.matMul(ctx, a, b)
	if err != nil {
		return nil, false, err
	}
	_ = s.kvs.Set(ctx, key, resp, s.ttl)
	return resp, false, nil
}

func (s *Service) ComputeStats(ctx context.Context, in StatsDTO) (*eng.VectorStatsReply, bool, error) {
	key := cache.Key("engine:stats", in)
	var cached eng.VectorStatsReply
	if ok, _ := s.kvs.Get(ctx, key, &cached); ok {
		return &cached, true, nil
	}
	sample := true
	if in.Sample != nil {
		sample = *in.Sample
	}
	resp, err := s.c.computeStats(ctx, in.Data, sample)
	if err != nil {
		return nil, false, err
	}
	_ = s.kvs.Set(ctx, key, resp, s.ttl)
	return resp, false, nil
}
