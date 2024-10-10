package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func EnvConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file : %s", err.Error())
	}

	// ! RPC URL CONFIG
	TendermintRpcUrl = os.Getenv("TENDERMINT_RPC_URL")
	if TendermintRpcUrl == "" {
		return fmt.Errorf("TENDERMINT_RPC_URL is not set")
	}

	TendermintApiUrl = os.Getenv("TENDERMINT_API_URL")
	if TendermintApiUrl == "" {
		return fmt.Errorf("TENDERMINT_API_URL is not set")
	}

	// ! SERVER CONFIG
	ServerPort = os.Getenv("SERVER_PORT")
	if ServerPort == "" {
		return fmt.Errorf("SERVER_PORT is not set")
	}
	GinEvn = os.Getenv("GIN_ENV")
	if ServerPort == "" {
		return fmt.Errorf("GIN_ENV is not set")
	}

	// ! DATABASE CONFIG
	MongoUrl = os.Getenv("MONGO_URL")
	if MongoUrl == "" {
		return fmt.Errorf("MONGO_URL is not set")
	}

	return nil
}
