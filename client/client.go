package client

import (
	"context"
	"fmt"
	"github.com/airchains-network/tracks-espresso-client/config"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

type Client struct {
	*cosmosclient.Client
}

func InitClient(ctx context.Context) (*Client, error) {
	c, err := cosmosclient.New(ctx, cosmosclient.WithNodeAddress(config.TendermintRpcUrl))
	if err != nil {
		return nil, fmt.Errorf("error initializing client: %w", err)
	}

	return &Client{
		Client: &c,
	}, nil
}
