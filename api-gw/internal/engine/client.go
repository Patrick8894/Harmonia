package engine

import (
	"context"
	"time"

	eng "github.com/Patrick8894/harmonia/api-gw/gen/engine"
	"github.com/apache/thrift/lib/go/thrift"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client { return &Client{addr: addr} }

func (c *Client) dial() (thrift.TTransport, *eng.EngineServiceClient, error) {
	tf := thrift.NewTBufferedTransportFactory(8192)
	pf := thrift.NewTBinaryProtocolFactoryConf(nil)
	cfg := &thrift.TConfiguration{
		ConnectTimeout: 1 * time.Second,
		SocketTimeout:  3 * time.Second,
	}

	sock := thrift.NewTSocketConf(c.addr, cfg)
	if sock == nil {
		return nil, nil, thrift.NewTTransportException(thrift.NOT_OPEN, "failed to create socket")
	}
	transport, err := tf.GetTransport(sock)
	if err != nil {
		return nil, nil, err
	}
	if err := transport.Open(); err != nil {
		transport.Close()
		return nil, nil, err
	}
	iprot := pf.GetProtocol(transport)
	oprot := pf.GetProtocol(transport)
	tclient := thrift.NewTStandardClient(iprot, oprot)
	cli := eng.NewEngineServiceClient(tclient)
	return transport, cli, nil
}

func (c *Client) hello(ctx context.Context, name string) (string, error) {
	transport, cli, err := c.dial()
	if err != nil {
		return "", err
	}
	defer transport.Close()

	resp, err := (*cli).Hello(ctx, &eng.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	return resp.GetMessage(), nil
}

func (c *Client) estimatePi(ctx context.Context, samples int64) (*eng.PiReply, error) {
	transport, cli, err := c.dial()
	if err != nil {
		return nil, err
	}
	defer transport.Close()

	return (*cli).EstimatePi(ctx, &eng.PiRequest{Samples: samples})
}

func (c *Client) matMul(ctx context.Context, a, b *eng.Matrix) (*eng.MatReply, error) {
	transport, cli, err := c.dial()
	if err != nil {
		return nil, err
	}
	defer transport.Close()

	return (*cli).MatMul(ctx, &eng.MatMulRequest{A: a, B: b})
}

func (c *Client) computeStats(ctx context.Context, data []float64, sample bool) (*eng.VectorStatsReply, error) {
	transport, cli, err := c.dial()
	if err != nil {
		return nil, err
	}
	defer transport.Close()

	return (*cli).ComputeStats(ctx, &eng.VectorStatsRequest{Data: data, Sample: sample})
}
