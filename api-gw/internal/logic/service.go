package logic

import (
	"context"

	lg "github.com/Patrick8894/harmonia/api-gw/gen/logic/v1"
)

type Service struct {
	c *Client
}

func NewService(c *Client) *Service { return &Service{c: c} }

func (s *Service) Hello(ctx context.Context, name string) (string, error) {
	return s.c.hello(ctx, name)
}

func (s *Service) Evaluate(ctx context.Context, in EvalDTO) (*lg.EvalReply, error) {
	req := &lg.EvalRequest{
		Expression: in.Expression,
		Variables:  in.Variables,
	}
	return s.c.evaluate(ctx, req)
}

func (s *Service) Transform(ctx context.Context, in TransformDTO) (*lg.TransformReply, error) {
	req := &lg.TransformRequest{
		Data:    in.Data,
		Expr:    in.Expr,
		VarName: in.VarName,
		Op:      parseTransformOp(in.Op),
	}
	return s.c.transform(ctx, req)
}

func (s *Service) PlanTasks(ctx context.Context, in PlanDTO) (*lg.PlanReply, error) {
	req := &lg.PlanRequest{
		Goal:     in.Goal,
		Hints:    in.Hints,
		MaxSteps: in.MaxSteps,
	}
	return s.c.planTasks(ctx, req)
}
