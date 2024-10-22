package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/airchains-network/tracks-espresso-client/batches"
	"github.com/airchains-network/tracks-espresso-client/client"
	"github.com/airchains-network/tracks-espresso-client/config"
	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/server"
	"github.com/airchains-network/tracks-espresso-client/server/espresso"
	// "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/deadlium/deadlogs"
	// "github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/airchains-network/tracks-espresso-client/components"
	"log"

)

func main() {
	ctx := context.Background()

	// Load environment variables
	if err := config.EnvConfig(); err != nil {
		deadlogs.Error(fmt.Sprintf("Error configuring environment: %s", err.Error()))
		return
	}
	deadlogs.Success("Environment configured successfully")

	// Initialize chain RPC connection
	clientInit, err := client.InitClient(ctx)
	if err != nil {
		deadlogs.Error(fmt.Sprintf("Error initializing client: %s", err.Error()))
		return
	}
	deadlogs.Success("Client initialized successfully")

	// Initialize MongoDB connection
	dbInstance, err := database.InitConnection() // Establish MongoDB connection
	if err != nil {
		deadlogs.Error(fmt.Sprintf("Error initializing database: %s", err.Error()))
		return
	}
	deadlogs.Success("Database initialized successfully")
	
	// batches.Batch()
	// Create a WaitGroup for the server and data load
	var wg sync.WaitGroup

	// Start the data load function in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		espresso.DataLoadFunction(dbInstance)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		// batches.Batch(ctx ,*dbInstance ,*clientInit.Client ,cosmosaccount.Account{
		// 	Name: "bob",
		// })
		gasFees := "2000air" // Specify the appropriate gas fees
	_, _, account, err := components.SetupAccountClient(ctx, gasFees)
	if err != nil {
		log.Fatalf("Failed to set up account client: %v", err)
	}
		batches.Batch(ctx ,*dbInstance ,*clientInit.Client ,account)

	}()
	// Start the server in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		serverInstance := server.InitServer(ctx, dbInstance, clientInit)
		uri := fmt.Sprintf("0.0.0.0:%s", config.ServerPort)
		if err := serverInstance.Run(uri); err != nil {
			deadlogs.Error(err.Error())
		} else {
			deadlogs.Success("Server started successfully")
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
