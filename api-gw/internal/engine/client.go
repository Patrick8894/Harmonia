package engine

import (
	"context"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	eng "github.com/Patrick8894/harmonia/api-gw/gen/engine"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client { return &Client{addr: addr} }

func (c *Client) hello(ctx context.Context, name string) (string, error) {
	// Buffered transport + binary protocol to match your C++ server
	tf := thrift.NewTBufferedTransportFactory(8192)
	pf := thrift.NewTBinaryProtocolFactoryConf(nil)
	cfg := &thrift.TConfiguration{
		ConnectTimeout: 1 * time.Second,
		SocketTimeout:  2 * time.Second,
	}

	sock := thrift.NewTSocketConf(c.addr, cfg)
	if sock == nil {
		return "", thrift.NewTTransportException(thrift.NOT_OPEN, "failed to create socket")
	}

	transport, err := tf.GetTransport(sock)
	if err != nil {
		return "", err
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		return "", err
	}

	iprot := pf.GetProtocol(transport)
	oprot := pf.GetProtocol(transport)
	tclient := thrift.NewTStandardClient(iprot, oprot)
	client := eng.NewEngineServiceClient(tclient)

	resp, err := client.Hello(ctx, &eng.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	return resp.GetMessage(), nil
}
