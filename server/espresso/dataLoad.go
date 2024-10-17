package espresso

import (
	"encoding/json"
	"fmt"
	"github.com/airchains-network/tracks-espresso-client/config"
	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/deadlium/deadlogs"
	"os"
	"time"
)

// DataLoadFunction Load data from JSON file to MongoDB
func DataLoadFunction(mongo *database.DB) {
	for {
		var espressoData []interface{}

		// Acquire the file lock to prevent concurrent access
		fileLock.Lock()

		// Check if the file exists, if not, create it
		if _, err := os.Stat(config.FilePath); os.IsNotExist(err) {
			// Create the directory if it doesn't exist
			if err := os.MkdirAll("file.data", os.ModePerm); err != nil {
			}
			// Create the file and initialize it with an empty JSON array
			if err := os.WriteFile(config.FilePath, []byte("[]"), 0644); err != nil {
			}
		}

		// Read file contents
		fileBytes, err := os.ReadFile(config.FilePath)
		if err != nil {
			deadlogs.Warn(fmt.Sprintf("Error reading file: %s", err.Error()))
			fileLock.Unlock()
		}

		// Unmarshal JSON data into espressoData slice
		if err = json.Unmarshal(fileBytes, &espressoData); err != nil {
			deadlogs.Warn(fmt.Sprintf("Error unmarshalling file data: %s", err.Error()))
			fileLock.Unlock()
		}

		// Insert the data into MongoDB
		err = mongo.InsertMany(espressoData)
		if err != nil {
			deadlogs.Warn(fmt.Sprintf("Error inserting into MongoDB: %s", err.Error()))
			fileLock.Unlock()
		}

		// Purge data from file after successful MongoDB insertion
		err = os.WriteFile(config.FilePath, []byte("[]"), 0644)
		if err != nil {
			deadlogs.Warn(fmt.Sprintf("Error clearing file after insertion: %s", err.Error()))
			fileLock.Unlock()
		}

		// Release the file lock
		fileLock.Unlock()

		deadlogs.Debug("Data loaded successfully, file purged, retrying in 30 seconds")

		// Wait 30 seconds before retrying
		time.Sleep(time.Second * 30)
	}
}
