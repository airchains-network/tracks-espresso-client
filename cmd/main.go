package main

import (
	"context"
	"fmt"
	// "path/filepath"

	"github.com/airchains-network/tracks-espresso-client/client"
	"github.com/airchains-network/tracks-espresso-client/config"

	"github.com/deadlium/deadlogs"

	// "log"
	// "os"
	"sync"

	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/pruning"
	"github.com/airchains-network/tracks-espresso-client/server"
	"github.com/airchains-network/tracks-espresso-client/server/espresso"
	"time"
	// "path/filepath"
)
var dataLock sync.Mutex
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
	// if _ , err := espresso.LoadDataFromFile(config.FilePath, dbInstance); err != nil {
	// 	log.Fatalf("Error loading and inserting data: %s", err)
	// } else {
	// 	deadlogs.Success("Data loaded and inserted successfully")
	// }	

	// Print the loaded data for debugging (optional)
	// fmt.Printf("Loaded Espresso Data: %+v\n", espressoData)
	// Print the loaded data (you can modify this as per your requirement)
	// fmt.Printf("Espresso Data: %+v\n", espressoData)
	go func() {
		for {
			deadlogs.Info("Checking and loading data from JSON file...")

			// Lock the data to prevent collision with pruning
			dataLock.Lock()
			time.Sleep(time.Minute *1)
			if _, err := espresso.LoadDataFromFile(config.FilePath, dbInstance); err != nil {
				deadlogs.Error(fmt.Sprintf("Error loading data: %s", err))
			} else {
				deadlogs.Success("Data loaded and inserted successfully")
			}
			dataLock.Unlock()

			// Wait before the next check
			time.Sleep(2 * time.Minute)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		pruning.StartPruningScheduler(ctx,config.FilePath ,dbInstance)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		serverInstance := server.InitServer(ctx, dbInstance, clientInit)
		uri := fmt.Sprintf("0.0.0.0:%s", config.ServerPort)
		runnerErr := serverInstance.Run(uri)
		if runnerErr != nil {
			deadlogs.Error(runnerErr.Error())
		} else {
			deadlogs.Success("server started")
		}
		wg.Done()
	}()

	wg.Wait()
}
