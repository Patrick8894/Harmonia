package logic

import (
	"context"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	lg "github.com/Patrick8894/harmonia/api-gw/gen/logic/v1"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client { return &Client{addr: addr} }

func (c *Client) dial() (*grpc.ClientConn, lg.LogicServiceClient, error) {
	conn, err := grpc.NewClient(
		c.addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}
	return conn, lg.NewLogicServiceClient(conn), nil
}

func (c *Client) hello(ctx context.Context, name string) (string, error) {
	conn, cli, err := c.dial()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp, err := cli.Hello(ctx, &lg.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	return resp.GetMessage(), nil
}

func (c *Client) evaluate(ctx context.Context, in *lg.EvalRequest) (*lg.EvalReply, error) {
	conn, cli, err := c.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return cli.Evaluate(ctx, in)
}

func (c *Client) transform(ctx context.Context, in *lg.TransformRequest) (*lg.TransformReply, error) {
	conn, cli, err := c.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return cli.Transform(ctx, in)
}

func (c *Client) planTasks(ctx context.Context, in *lg.PlanRequest) (*lg.PlanReply, error) {
	conn, cli, err := c.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Planning can take longer
	ctx, cancel := context.WithTimeout(ctx, 8*time.Second)
	defer cancel()

	return cli.PlanTasks(ctx, in)
}

// helper: map string/number to proto enum
func parseTransformOp(s string) lg.TransformOp {
	if s == "" {
		return lg.TransformOp_TRANSFORM_OP_UNSPECIFIED
	}
	up := strings.ToUpper(strings.TrimSpace(s))
	switch up {
	case "MAP":
		return lg.TransformOp_MAP
	case "FILTER":
		return lg.TransformOp_FILTER
	case "SUM":
		return lg.TransformOp_SUM
	}
	// accept numeric form as fallback
	if n, err := strconv.Atoi(up); err == nil {
		return lg.TransformOp(n)
	}
	return lg.TransformOp_TRANSFORM_OP_UNSPECIFIED
}
