package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Config struct {
	Api struct {
		Token     string
		Host      string
		Interface string
	}
}

func Init(cmd *cobra.Command) (*Config, error) {
	netInterface, err := cmd.Flags().GetString("interface")
	if err != nil || len(netInterface) == 0 {
		return nil, fmt.Errorf("no such interface: %s", err.Error())
	}

	apiHost, err := cmd.Flags().GetString("host")
	if err != nil || len(apiHost) == 0 {
		return nil, fmt.Errorf("no such api host: %s", err.Error())
	}

	token, err := cmd.Flags().GetString("token")
	if err != nil || len(token) == 0 {
		return nil, fmt.Errorf("no such token: %s", err.Error())
	}

	cfg := &Config{}
	cfg.Api.Host = apiHost
	cfg.Api.Interface = netInterface
	cfg.Api.Token = token
	return cfg, nil
}
