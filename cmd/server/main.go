package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/FrakenboK/j8-sniffer/internal/config"
	"github.com/FrakenboK/j8-sniffer/internal/servers/grpc_api"
)

func main() {
	cfg := config.NewConfig()
	fmt.Printf("Server token: %s\n", cfg.Server.Token)

	srv := grpc_api.NewServer(cfg)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-stopChan:
		srv.GracefulStop()

	case <-errChan:
		// TODO: Add slog
		panic("error occured")
	}
}
