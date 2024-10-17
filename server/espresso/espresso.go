package espresso

import (
	// "context"
	"fmt"
	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/types"
	"github.com/gin-gonic/gin"
	// "net/http"
	"os"
	// "path/filepath"
	"encoding/json"
    // "path/filepath"
    "time"
    // "github.com/airchains-network/tracks-espresso-client/config"

)

// func TracksEspressoDataLoad(ctx context.Context, mongoDB *database.DB) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		var tracksEspressoStruct types.EspressoSchemaV1
// 		if err := c.ShouldBindJSON(&tracksEspressoStruct); err != nil {
// 			c.JSON(http.StatusBadRequest, types.Response{
// 				Status:      false,
// 				Message:     "Error parsing request body",
// 				Description: fmt.Sprintf("Error parsing request body : %s", err.Error()),
// 			})
// 			return
// 		}

// 		// Convert tracksEspressoStruct to mongoTracksEspressoStruct
// 		// mongoTracksEspressoStruct := types.MongoTracksEspressoStruct{
// 			// StationId:    tracksEspressoStruct.StationId,
// 			// EspressoData: types.MongoEspressoDataStruct{
// 			// 	Transaction: struct {
// 			// 		Namespace int    `bson:"namespace"`
// 			// 		Payload   string `bson:"payload"`
// 			// 	}{
// 			// 		Namespace: tracksEspressoStruct.EspressoData.Transaction.Namespace,
// 			// 		Payload:   tracksEspressoStruct.EspressoData.Transaction.Payload,
// 			// 	},
// 			// 	Hash:  tracksEspressoStruct.EspressoData.Hash,
// 			// 	Index: tracksEspressoStruct.EspressoData.Index,
// 			// 	Proof: struct {
// 			// 		TxIndex            []int `bson:"tx_index"`
// 			// 		PayloadNumTxs      []int `bson:"payload_num_txs"`
// 			// 		PayloadProofNumTxs struct {
// 			// 			Proofs      string        `bson:"proofs"`
// 			// 			PrefixBytes []interface{} `bson:"prefix_bytes"`
// 			// 			SuffixBytes []int         `bson:"suffix_bytes"`
// 			// 		} `bson:"payload_proof_num_txs"`
// 			// 		PayloadTxTableEntries      []int `bson:"payload_tx_table_entries"`
// 			// 		PayloadProofTxTableEntries struct {
// 			// 			Proofs      string `bson:"proofs"`
// 			// 			PrefixBytes []int  `bson:"prefix_bytes"`
// 			// 			SuffixBytes []int  `bson:"suffix_bytes"`
// 			// 		} `bson:"payload_proof_tx_table_entries"`
// 			// 		PayloadProofTx struct {
// 			// 			Proofs      string        `bson:"proofs"`
// 			// 			PrefixBytes []int         `bson:"prefix_bytes"`
// 			// 			SuffixBytes []interface{} `bson:"suffix_bytes"`
// 			// 		} `bson:"payload_proof_tx"`
// 			// 	}{
// 			// 		TxIndex:       tracksEspressoStruct.EspressoData.Proof.TxIndex,
// 			// 		PayloadNumTxs: tracksEspressoStruct.EspressoData.Proof.PayloadNumTxs,
// 			// 		PayloadProofNumTxs: struct {
// 			// 			Proofs      string        `bson:"proofs"`
// 			// 			PrefixBytes []interface{} `bson:"prefix_bytes"`
// 			// 			SuffixBytes []int         `bson:"suffix_bytes"`
// 			// 		}{
// 			// 			Proofs:      tracksEspressoStruct.EspressoData.Proof.PayloadProofNumTxs.Proofs,
// 			// 			PrefixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofNumTxs.PrefixBytes,
// 			// 			SuffixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofNumTxs.SuffixBytes,
// 			// 		},
// 			// 		PayloadTxTableEntries: tracksEspressoStruct.EspressoData.Proof.PayloadTxTableEntries,
// 			// 		PayloadProofTxTableEntries: struct {
// 			// 			Proofs      string `bson:"proofs"`
// 			// 			PrefixBytes []int  `bson:"prefix_bytes"`
// 			// 			SuffixBytes []int  `bson:"suffix_bytes"`
// 			// 		}{
// 			// 			Proofs:      tracksEspressoStruct.EspressoData.Proof.PayloadProofTxTableEntries.Proofs,
// 			// 			PrefixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofTxTableEntries.PrefixBytes,
// 			// 			SuffixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofTxTableEntries.SuffixBytes,
// 			// 		},
// 			// 		PayloadProofTx: struct {
// 			// 			Proofs      string        `bson:"proofs"`
// 			// 			PrefixBytes []int         `bson:"prefix_bytes"`
// 			// 			SuffixBytes []interface{} `bson:"suffix_bytes"`
// 			// 		}{
// 			// 			Proofs:      tracksEspressoStruct.EspressoData.Proof.PayloadProofTx.Proofs,
// 			// 			PrefixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofTx.PrefixBytes,
// 			// 			SuffixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofTx.SuffixBytes,
// 			// 		},
// 			// 	},
// 			// 	BlockHash:   tracksEspressoStruct.EspressoData.BlockHash,
// 			// 	BlockHeight: tracksEspressoStruct.EspressoData.BlockHeight,
// 			// },
// 		// }

// 		// _ = mongoTracksEspressoStruct

// }

// var mu sync.Mutex

// TracksEspressoDataLoad handles the API to load data into main.json and MongoDB
func TracksEspressoDataLoad(filePath string, db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		// Check if the file exists, if not, create it
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// Create the directory if it doesn't exist
			if err := os.MkdirAll("file.data", os.ModePerm); err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				c.JSON(500, gin.H{"message": "Error creating directory"})
				return
			}

			// Create the file and initialize it with an empty JSON array
			if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
				fmt.Printf("Error creating file: %v\n", err)
				c.JSON(500, gin.H{"message": "Error creating file"})
				return
			}
		}

		// Load existing data from the file
		var espressoData []types.EspressoSchemaV1
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			c.JSON(500, gin.H{"message": "Error reading file"})
			return
		}

		if err := json.Unmarshal(fileBytes, &espressoData); err != nil {
			fmt.Printf("Error unmarshalling file: %v\n", err)
			c.JSON(500, gin.H{"message": "Error unmarshalling data"})
			return
		}

		// Bind incoming data from the API request
		var newData types.EspressoSchemaV1
		if err := c.ShouldBindJSON(&newData); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		// Append new data to the existing array
		espressoData = append(espressoData, newData)

		// Marshal the updated data back to JSON and write to the file
		fileData, err := json.MarshalIndent(espressoData, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling data: %v\n", err)
			c.JSON(500, gin.H{"message": "Error saving data"})
			return
		}

		// Write the updated data to the file
		if err := os.WriteFile(filePath, fileData, 0644); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			c.JSON(500, gin.H{"message": "Error saving data"})
			return
		}

		// Insert new data into MongoDB
		fmt.Println("i'm inserting espresso")
		time.Sleep(3 *time.Minute)
		if err := db.InsertMany("espressodata", []interface{}{newData}); err != nil {
			fmt.Printf("Error inserting data into MongoDB: %v\n", err)
			c.JSON(500, gin.H{"message": "Error saving to database"})
			return
		}

		c.JSON(200, gin.H{"message": "Data saved successfully"})
	}
}

// InitBufferFlush initializes a timer to flush data to MongoDB every 5 minutes
func InitBufferFlush(filePath string, db *database.DB) {
	flushInterval := 5 * time.Minute
	ticker := time.NewTicker(flushInterval)

	go func() {
		for {
			<-ticker.C
			flushFileDataToMongo(filePath, db)
		}
	}()
}

// flushFileDataToMongo sends accumulated data from the file to MongoDB every 5 minutes
func flushFileDataToMongo(filePath string, db *database.DB) {
	mu.Lock()
	defer mu.Unlock()

	// Read data from file
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file for flushing: %v\n", err)
		return
	}

	// Unmarshal file data into EspressoSchemaV1 array
	var espressoData []types.EspressoSchemaV1
	if err := json.Unmarshal(fileBytes, &espressoData); err != nil {
		fmt.Printf("Error unmarshalling data for flushing: %v\n", err)
		return
	}

	// If no data, return early
	if len(espressoData) == 0 {
		fmt.Println("No data to flush to MongoDB")
		return
	}

	// Insert data into MongoDB
	var mongoDocuments []interface{}
	for _, data := range espressoData {
		mongoDocuments = append(mongoDocuments, data)
	}

	if err := db.InsertMany("espressodata", mongoDocuments); err != nil {
		fmt.Printf("Error inserting data into MongoDB: %v\n", err)
		return
	}

	// Clear the file after successful insertion
	if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
		fmt.Printf("Error clearing file: %v\n", err)
		return
	}

	fmt.Println("Data successfully flushed to MongoDB and cleared from file.")
}