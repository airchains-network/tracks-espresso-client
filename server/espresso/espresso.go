package espresso

import (
	"encoding/json"
	"fmt"
	"github.com/airchains-network/tracks-espresso-client/config"
	"github.com/airchains-network/tracks-espresso-client/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
)

var fileLock sync.Mutex // Global mutex lock

// TracksEspressoDataLoad API handler to load data into JSON file
func TracksEspressoDataLoad() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tracksEspressoStruct types.EspressoSchemaV1
		if err := c.ShouldBindJSON(&tracksEspressoStruct); err != nil {
			c.JSON(http.StatusBadRequest, types.Response{
				Status:      false,
				Message:     "Error parsing request body",
				Description: fmt.Sprintf("Error parsing request body : %s", err.Error()),
			})
			return
		}

		// Attempt to acquire the file lock using TryLock()
		if !fileLock.TryLock() {
			// If the file is locked, respond with a message
			c.JSON(http.StatusTooEarly, types.Response{
				Status:      false,
				Message:     "File is currently locked, please try again later.",
				Description: "Data is being processed, try again after a while.",
			})
			return
		}

		// Ensure the lock is released after processing
		defer fileLock.Unlock()

		// Load existing data from the file
		var espressoData []types.MongoTracksEspressoStruct
		fileBytes, err := os.ReadFile(config.FilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.Response{
				Status:      false,
				Message:     "Error reading data",
				Description: fmt.Sprintf("Error reading data : %s", err.Error()),
			})
			return
		}

		if err := json.Unmarshal(fileBytes, &espressoData); err != nil {
			c.JSON(http.StatusInternalServerError, types.Response{
				Status:      false,
				Message:     "Error unmarshalling data",
				Description: fmt.Sprintf("Error unmarshalling data : %s", err.Error()),
			})
			return
		}

		// Append new data to the existing data
		espressoDataLoad := append(espressoData, types.MongoTracksEspressoStruct{
			EspressoStationId: fmt.Sprintf("%s-%v", tracksEspressoStruct.StationID, tracksEspressoStruct.PodNumber),
			EspressoTxResponseV1: types.EspressoTxResponseV1{
				Transaction: struct {
					Namespace int    `json:"namespace"`
					Payload   string `json:"payload"`
				}{
					Namespace: tracksEspressoStruct.EspressoTxResponseV1.Transaction.Namespace,
					Payload:   tracksEspressoStruct.EspressoTxResponseV1.Transaction.Payload,
				},
				Hash:  tracksEspressoStruct.EspressoTxResponseV1.Hash,
				Index: tracksEspressoStruct.EspressoTxResponseV1.Index,
				Proof: struct {
					TxIndex            string `json:"tx_index"`
					PayloadNumTxs      string `json:"payload_num_txs"`
					PayloadProofNumTxs struct {
						Proofs      string `json:"proofs"`
						PrefixBytes string `json:"prefix_bytes"`
						SuffixBytes string `json:"suffix_bytes"`
					} `json:"payload_proof_num_txs"`
					PayloadTxTableEntries      string `json:"payload_tx_table_entries"`
					PayloadProofTxTableEntries struct {
						Proofs      string `json:"proofs"`
						PrefixBytes string `json:"prefix_bytes"`
						SuffixBytes string `json:"suffix_bytes"`
					} `json:"payload_proof_tx_table_entries"`
					PayloadProofTx struct {
						Proofs      string `json:"proofs"`
						PrefixBytes string `json:"prefix_bytes"`
						SuffixBytes string `json:"suffix_bytes"`
					} `json:"payload_proof_tx"`
				}{
					TxIndex:       tracksEspressoStruct.EspressoTxResponseV1.Proof.TxIndex,
					PayloadNumTxs: tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadNumTxs,
					PayloadProofNumTxs: struct {
						Proofs      string `json:"proofs"`
						PrefixBytes string `json:"prefix_bytes"`
						SuffixBytes string `json:"suffix_bytes"`
					}{
						Proofs:      tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofNumTxs.Proofs,
						PrefixBytes: tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofNumTxs.PrefixBytes,
						SuffixBytes: tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofNumTxs.SuffixBytes,
					},
					PayloadTxTableEntries: tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadTxTableEntries,
					PayloadProofTxTableEntries: struct {
						Proofs      string `json:"proofs"`
						PrefixBytes string `json:"prefix_bytes"`
						SuffixBytes string `json:"suffix_bytes"`
					}{
						Proofs:      tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofTxTableEntries.Proofs,
						PrefixBytes: tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofTxTableEntries.PrefixBytes,
						SuffixBytes: tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofTxTableEntries.SuffixBytes,
					},
					PayloadProofTx: struct {
						Proofs      string `json:"proofs"`
						PrefixBytes string `json:"prefix_bytes"`
						SuffixBytes string `json:"suffix_bytes"`
					}{
						Proofs:      tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofTx.Proofs,
						PrefixBytes: tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofTx.PrefixBytes,
						SuffixBytes: tracksEspressoStruct.EspressoTxResponseV1.Proof.PayloadProofTx.SuffixBytes,
					},
				},
				BlockHash:   tracksEspressoStruct.EspressoTxResponseV1.BlockHash,
				BlockHeight: tracksEspressoStruct.EspressoTxResponseV1.BlockHeight,
			},
			StationId: tracksEspressoStruct.StationID,
			PodNumber: tracksEspressoStruct.PodNumber,
		})

		// Marshal the updated data back to JSON and write to the file
		fileData, err := json.MarshalIndent(espressoDataLoad, "", "  ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.Response{
				Status:      false,
				Message:     "Error marshalling data",
				Description: fmt.Sprintf("Error marshalling data : %s", err.Error()),
			})
			return
		}

		// Write the updated data to the file
		if err := os.WriteFile(config.FilePath, fileData, 0644); err != nil {
			c.JSON(http.StatusInternalServerError, types.Response{
				Status:      false,
				Message:     "Error writing to file",
				Description: fmt.Sprintf("Error writing to file : %s", err.Error()),
			})
			return
		}

		// Respond with success
		c.JSON(http.StatusOK, types.Response{
			Status:      true,
			Message:     "Data load successful",
			Description: "Data successfully saved to file",
		})
	}
}
