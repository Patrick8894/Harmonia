package logic

import (
	"context"
	"time"

	lg "github.com/Patrick8894/harmonia/api-gw/gen/logic/v1"
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

func (s *Service) Evaluate(ctx context.Context, in EvalDTO) (*lg.EvalReply, bool, error) {
	key := cache.Key("logic:eval", in)
	var cached lg.EvalReply
	if ok, _ := s.kvs.Get(ctx, key, &cached); ok {
		return &cached, true, nil
	}
	resp, err := s.c.evaluate(ctx, &lg.EvalRequest{
		Expression: in.Expression,
		Variables:  in.Variables,
	})
	if err != nil {
		return nil, false, err
	}
	_ = s.kvs.Set(ctx, key, resp, s.ttl)
	return resp, false, nil
}

func (s *Service) Transform(ctx context.Context, in TransformDTO) (*lg.TransformReply, bool, error) {
	key := cache.Key("logic:xform", in)
	var cached lg.TransformReply
	if ok, _ := s.kvs.Get(ctx, key, &cached); ok {
		return &cached, true, nil
	}
	resp, err := s.c.transform(ctx, &lg.TransformRequest{
		Data: in.Data, Expr: in.Expr, VarName: in.VarName, Op: parseTransformOp(in.Op),
	})
	if err != nil {
		return nil, false, err
	}
	_ = s.kvs.Set(ctx, key, resp, s.ttl)
	return resp, false, nil
}

func (s *Service) PlanTasks(ctx context.Context, in PlanDTO) (*lg.PlanReply, bool, error) {
	key := cache.Key("logic:plan", in)
	var cached lg.PlanReply
	if ok, _ := s.kvs.Get(ctx, key, &cached); ok {
		return &cached, true, nil
	}
	resp, err := s.c.planTasks(ctx, &lg.PlanRequest{
		Goal: in.Goal, Hints: in.Hints, MaxSteps: in.MaxSteps,
	})
	if err != nil {
		return nil, false, err
	}
	_ = s.kvs.Set(ctx, key, resp, s.ttl)
	return resp, false, nil
}
