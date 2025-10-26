package engine

import "context"

type Service struct {
	c *Client
}

func NewService(c *Client) *Service { return &Service{c: c} }

func (s *Service) Hello(ctx context.Context, name string) (string, error) {
	return s.c.hello(ctx, name)
}
