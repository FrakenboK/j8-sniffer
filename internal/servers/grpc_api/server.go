package grpc_api

import (
	"fmt"
	"net"

	"github.com/FrakenboK/j8-sniffer/internal/config"
	"github.com/FrakenboK/j8-sniffer/internal/servers/grpc_api/service"
	"github.com/FrakenboK/j8-sniffer/pkg/protobuf"
	"google.golang.org/grpc"
)

type Server struct {
	cfg  *config.Config
	grpc *grpc.Server
}

func (srv *Server) ListenAndServe() error {
	addr := fmt.Sprintf("%s:%s", srv.cfg.Server.Host, srv.cfg.Server.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	defer func() {
		_ = listener.Close()
	}()

	service := service.NewNetworkSnifferService(srv.cfg)
	protobuf.RegisterNetworkSnifferApiServer(srv.grpc, service)

	return srv.grpc.Serve(listener)
}

func (srv *Server) GracefulStop() {
	srv.grpc.GracefulStop()
}

func NewServer(cfg *config.Config) *Server {
	grpc := grpc.NewServer()

	srv := &Server{
		cfg:  cfg,
		grpc: grpc,
	}
	return srv
}
