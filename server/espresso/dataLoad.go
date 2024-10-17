// 
package espresso

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/types"
	"github.com/deadlium/deadlogs"
)

// Define a mutex for thread safety
var mu sync.Mutex

// LoadDataFromFile keeps checking for data in the JSON file until data is found, then inserts it into MongoDB.
func LoadDataFromFile(filePath string, db *database.DB) ([]types.EspressoSchemaV1, error) {
	mu.Lock()
	defer mu.Unlock()

	var espressoData []types.EspressoSchemaV1

	// Ensure the file exists; if not, create it
	if err := ensureFileExists(filePath); err != nil {
		return nil, err
	}

	// Continuous loop until data is found
	for {
		// Open the JSON file
		jsonFile, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open JSON file: %s", err)
		}

		// Read the file content
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			jsonFile.Close() // Close the file explicitly
			return nil, fmt.Errorf("failed to read JSON file: %s", err)
		}
		jsonFile.Close()

		// Check if the file is empty
		if len(byteValue) == 0 {
			deadlogs.Warn("JSON file is empty, waiting before trying again...")
			time.Sleep(time.Second * 5) // Wait before checking again
			continue // Retry if the file is empty
		}

		// Unmarshal the JSON content into a slice of EspressoSchemaV1
		err = json.Unmarshal(byteValue, &espressoData)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
		}

		// If data was successfully loaded, break out of the loop
		if len(espressoData) > 0 {
			break
		}

		// If still no data, wait before trying again
		deadlogs.Warn("No data loaded, waiting before trying again...")
		time.Sleep(time.Minute * 5) // Wait before checking again
	}

	// Prepare data for MongoDB insertion
	var mongoDocuments []interface{}
	for _, data := range espressoData {
		mongoDocuments = append(mongoDocuments, data)
	}
	// time.Sleep(time.Minute * 4)
	// Insert data into MongoDB
	fmt.Println("im inserting dataload" )
	if err := db.InsertMany("espressodata", mongoDocuments); err != nil {
		return nil, fmt.Errorf("failed to insert data into MongoDB: %s", err)
	}

	deadlogs.Info("Data loaded and inserted into MongoDB successfully.")
	return espressoData, nil
}

// ensureFileExists checks if the file exists; if not, it creates it
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
