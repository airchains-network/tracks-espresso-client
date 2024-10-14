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

        // Load existing data from file if it exists
        var espressoData []types.EspressoSchemaV1
        if fileBytes, err := os.ReadFile(filePath); err == nil {
            if err := json.Unmarshal(fileBytes, &espressoData); err != nil {
                fmt.Printf("Error unmarshalling file: %v", err)
                c.JSON(500, gin.H{"message": "Error reading data"})
                return
            }
        } else if os.IsNotExist(err) {
            // If the file does not exist, create an empty array
            espressoData = []types.EspressoSchemaV1{}
        } else {
            fmt.Printf("Error reading file: %v", err)
            c.JSON(500, gin.H{"message": "Error reading file"})
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

        // Marshal the updated data back to JSON
        fileData, err := json.MarshalIndent(espressoData, "", "  ")
        if err != nil {
            fmt.Printf("Error marshalling data: %v", err)
            c.JSON(500, gin.H{"message": "Error saving data"})
            return
        }

        // Write the updated data to the file (create file if it doesn't exist)
        if err := os.WriteFile(filePath, fileData, 0644); err != nil {
            fmt.Printf("Error writing to file: %v", err)
            c.JSON(500, gin.H{"message": "Error saving data"})
            return
        }

        // Prepare documents for MongoDB insertion
        var mongoDocuments []interface{}
        mongoDocuments = append(mongoDocuments, newData)

        // Insert data into MongoDB
        if err := db.InsertMany("espressodata", mongoDocuments); err != nil {
            fmt.Printf("Error inserting data into MongoDB: %v", err)
            c.JSON(500, gin.H{"message": "Error saving to database"})
            return
        }

        c.JSON(200, gin.H{"message": "Data saved successfully"})
    }
}

