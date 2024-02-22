package service

import (
	"context"

	"github.com/FrakenboK/j8-sniffer/internal/config"
	"github.com/FrakenboK/j8-sniffer/pkg/protobuf"
)

type NetworkSnifferService struct {
	protobuf.NetworkSnifferApiServer
	cfg *config.Config
}

func (ns *NetworkSnifferService) ProcessNetworkData(ctx context.Context, req *protobuf.NetworkData) (*protobuf.EmptyResponse, error) {
	return &protobuf.EmptyResponse{}, nil
}

func NewNetworkSnifferService(cfg *config.Config) *NetworkSnifferService {
	ns := &NetworkSnifferService{
		cfg: cfg,
	}
	return ns
}
