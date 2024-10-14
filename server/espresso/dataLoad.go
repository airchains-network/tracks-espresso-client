package espresso

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/types"
	"github.com/deadlium/deadlogs"
)

// Define a mutex for thread safety
var mu sync.Mutex

// LoadDataFromFile loads the data from main.json and inserts it into MongoDB
func LoadDataFromFile(filePath string, db *database.DB) ([]types.EspressoSchemaV1, error) {
	mu.Lock() 

	// Open the JSON file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		mu.Unlock() // Unlock the mutex before returning
		return nil, fmt.Errorf("failed to open JSON file: %s", err)
	}

	// Read the file content
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		jsonFile.Close() // Close the file explicitly
		mu.Unlock() // Unlock the mutex before returning
		return nil, fmt.Errorf("failed to read JSON file: %s", err)
	}

	// Check if the file is empty
	if len(byteValue) == 0 {
		deadlogs.Warn("JSON file is empty, nothing to load.")
		jsonFile.Close() 
		mu.Unlock() 
		return nil, nil 
	}

	// Unmarshal the JSON content into a slice of EspressoDataStruct
	var espressoData []types.EspressoSchemaV1
	err = json.Unmarshal(byteValue, &espressoData)
	if err != nil {
		jsonFile.Close() // Close the file before returning
		mu.Unlock() // Unlock the mutex before returning
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}

	// Close the file after reading and processing
	jsonFile.Close()

	// Prepare data for MongoDB insertion
	var mongoDocuments []interface{}
	for _, data := range espressoData {
		mongoDocuments = append(mongoDocuments, data)
	}

	// Insert data into MongoDB
	if err := db.InsertMany("espressodata", mongoDocuments); err != nil {
		mu.Unlock() // Unlock the mutex before returning
		return nil, fmt.Errorf("failed to insert data into MongoDB: %s", err)
	}

	fmt.Println("Data loaded and inserted into MongoDB successfully.")
	mu.Unlock() // Unlock the mutex before returning successfully
	return espressoData, nil
}
