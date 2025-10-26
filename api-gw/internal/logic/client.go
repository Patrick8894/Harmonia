package logic

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	lg "github.com/Patrick8894/harmonia/api-gw/gen/logic/v1"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client { return &Client{addr: addr} }

func (c *Client) hello(ctx context.Context, name string) (string, error) {
	// keep your original NewClient usage for compatibility
	conn, err := grpc.NewClient(
		c.addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := lg.NewLogicServiceClient(conn)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp, err := client.Hello(ctx, &lg.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	return resp.GetMessage(), nil
}
