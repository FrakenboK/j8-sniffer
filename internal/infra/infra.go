package infra

import (
	grpcclient "github.com/FrakenboK/j8-sniffer/internal/grpc_client"
	"github.com/FrakenboK/j8-sniffer/internal/sniffer"
	"github.com/FrakenboK/j8-sniffer/internal/sniffer/config"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) {
	cfg, err := config.Init(cmd)
	if err != nil {
		cmd.Help()
		return
	}

	client, err := grpcclient.Init(cfg)
	if err != nil {
		panic("failed to create grpc client")
	}

	sniff := sniffer.Init(cfg, client)
	sniff.Run(cmd, args)
}
