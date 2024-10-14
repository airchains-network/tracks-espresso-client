package main

import (
	"context"
	"fmt"
	"github.com/airchains-network/tracks-espresso-client/client"
	"github.com/airchains-network/tracks-espresso-client/config"

	"github.com/deadlium/deadlogs"

	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/server"
	"github.com/airchains-network/tracks-espresso-client/server/espresso"
	"log"
	"github.com/airchains-network/tracks-espresso-client/pruning"
)

func main() {
	ctx := context.Background()

	// env variable load
	if err := config.EnvConfig(); err != nil {
		deadlogs.Error(fmt.Sprintf("error to config environment: %s", err.Error()))
	} else {
		deadlogs.Success("config environment set to success")
	}

	// chain rpc connection
	clientInit, err := client.InitClient(ctx)
	if err != nil {
		deadlogs.Error(fmt.Sprintf("error to init client: %s", err.Error()))
	} else {
		deadlogs.Success("client initialized")
	}

	// Initialize MongoDB connection
	dbInstance, err := database.InitConnection() // This will establish MongoDB connection
	if err != nil {
		deadlogs.Error(fmt.Sprintf("error to init database: %s", err.Error()))
	} else {
		deadlogs.Success("database initialized")
	}
	
	// Load the data from the JSON file and insert it into MongoDB
	if _ , err := espresso.LoadDataFromFile(config.FilePath, dbInstance); err != nil {
		log.Fatalf("Error loading and inserting data: %s", err)
	} else {
		deadlogs.Success("Data loaded and inserted successfully")
	}	

	// Print the loaded data for debugging (optional)
	// fmt.Printf("Loaded Espresso Data: %+v\n", espressoData)
	// Print the loaded data (you can modify this as per your requirement)
	// fmt.Printf("Espresso Data: %+v\n", espressoData)

	pruning.StartPruningScheduler(ctx, config.FilePath, dbInstance)
	// gin server init
	serverInstance := server.InitServer(ctx, dbInstance, clientInit)
	uri := fmt.Sprintf("0.0.0.0:%s", config.ServerPort)
	runnerErr := serverInstance.Run(uri)
	if runnerErr != nil {
		deadlogs.Error(runnerErr.Error())
	}else{
		deadlogs.Success("server started")
	}
}
