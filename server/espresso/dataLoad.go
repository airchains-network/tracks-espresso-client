package espresso

import (
	"encoding/json"
	"fmt"
	"github.com/airchains-network/tracks-espresso-client/config"
	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/deadlium/deadlogs"
	"github.com/airchains-network/tracks-espresso-client/types"

	"os"
	"time"
)

// DataLoadFunction Load data from JSON file to MongoDB
func DataLoadFunction(mongo *database.DB) {
	for {
		var espressoData []interface{}
		if err := ensureFileExists(config.FilePath); err != nil {
			// Handle error, you can log it or take necessary actions
			return
		}

		// Acquire the file lock to prevent concurrent access
		fileLock.Lock()

		// // Check if the file exists, if not, create it
		// if _, err := os.Stat(config.FilePath); os.IsNotExist(err) {
		// 	// Create the directory if it doesn't exist
		// 	if err := os.MkdirAll("file.data", os.ModePerm); err != nil {
		// 	}
		// 	// Create the file and initialize it with an empty JSON array
		// 	if err := os.WriteFile(config.FilePath, []byte("[]"), 0644); err != nil {
		// 	}
		// }

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
		
		// batchesData := BatchData(espressoData)

		var espressoDataInterface []interface{}

		// Convert []interface{} to []types.EspressoSchemaV1
		var espressoDataload []types.EspressoSchemaV1
		
		for _, item := range espressoDataInterface {
			if espresso, ok := item.(types.EspressoSchemaV1); ok {
				espressoData = append(espressoData, espresso)
			} else {
				// Handle the case where the type assertion fails, if necessary
				// You can log an error or skip this item
			}
		}
		
		// Now call BatchData with the correct type
		batchesData := BatchData(espressoDataload)
		fmt.Println("batchdata", batchesData)

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

		deadlogs.Success("Data loaded successfully, file purged, retrying in 1 min")

		// Wait 30 seconds before retrying
		time.Sleep(time.Second * 30)
	}
}
func ensureFileExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		emptyData := []byte("[]")
		if err := os.MkdirAll("file.data", os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %s", err)
		}
		if err := os.WriteFile(filePath, emptyData, 0644); err != nil {
			return fmt.Errorf("failed to create empty JSON file: %s", err)
		}
	}
	return nil
}
