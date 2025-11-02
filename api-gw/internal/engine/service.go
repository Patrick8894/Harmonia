package engine

import (
	"context"

	eng "github.com/Patrick8894/harmonia/api-gw/gen/engine"
)

type Service struct {
	c *Client
}

func NewService(c *Client) *Service { return &Service{c: c} }

func (s *Service) Hello(ctx context.Context, name string) (string, error) {
	return s.c.hello(ctx, name)
}

func (s *Service) EstimatePi(ctx context.Context, samples int64) (*eng.PiReply, error) {
	return s.c.estimatePi(ctx, samples)
}

func (s *Service) MatMul(ctx context.Context, in MatMulDTO) (*eng.MatReply, error) {
	a := &eng.Matrix{Rows: in.A.Rows, Cols: in.A.Cols, Data: in.A.Data}
	b := &eng.Matrix{Rows: in.B.Rows, Cols: in.B.Cols, Data: in.B.Data}
	return s.c.matMul(ctx, a, b)
}

func (s *Service) ComputeStats(ctx context.Context, in StatsDTO) (*eng.VectorStatsReply, error) {
	sample := true
	if in.Sample != nil {
		sample = *in.Sample
	}
	return s.c.computeStats(ctx, in.Data, sample)
}
