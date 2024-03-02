package grpcclient

import (
	"context"
	"fmt"

	"github.com/FrakenboK/j8-sniffer/internal/sniffer/config"
	"github.com/FrakenboK/j8-sniffer/pkg/protobuf"

	"google.golang.org/grpc"
)

type Client struct {
	grpc protobuf.NetworkSnifferApiClient
	conn *grpc.ClientConn
}

func (c *Client) Close() {
	defer c.conn.Close()
}

func (c *Client) Send(data []byte, dstPort, srcPort string) error {
	_, err := c.grpc.ProcessNetworkData(context.Background(), &protobuf.NetworkData{ApplicationData: data, DstPort: dstPort, SrcPort: srcPort})
	if err != nil {
		return fmt.Errorf("failed to send ProcessNetworkData request: %s", err.Error())
	}
	return nil
}

func Init(cfg *config.Config) (*Client, error) {
	connection, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.Api.Host, 58923), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection: %s", err.Error())
	}
	grpc := protobuf.NewNetworkSnifferApiClient(connection)
	client := &Client{
		grpc: grpc,
		conn: connection,
	}
	return client, nil
}
