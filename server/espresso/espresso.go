package espresso

import (
	"context"
	"fmt"
	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TracksEspressoDataLoad(ctx context.Context, mongoDB *database.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var tracksEspressoStruct types.TracksEspressoStruct
		if err := c.ShouldBindJSON(&tracksEspressoStruct); err != nil {
			c.JSON(http.StatusBadRequest, types.Response{
				Status:      false,
				Message:     "Error parsing request body",
				Description: fmt.Sprintf("Error parsing request body : %s", err.Error()),
			})
			return
		}

		// Convert tracksEspressoStruct to mongoTracksEspressoStruct
		mongoTracksEspressoStruct := types.MongoTracksEspressoStruct{
			StationId:    tracksEspressoStruct.StationId,
			EspressoVmId: tracksEspressoStruct.EspressoVmId,
			EspressoData: types.MongoEspressoDataStruct{
				Transaction: struct {
					Namespace int    `bson:"namespace"`
					Payload   string `bson:"payload"`
				}{
					Namespace: tracksEspressoStruct.EspressoData.Transaction.Namespace,
					Payload:   tracksEspressoStruct.EspressoData.Transaction.Payload,
				},
				Hash:  tracksEspressoStruct.EspressoData.Hash,
				Index: tracksEspressoStruct.EspressoData.Index,
				Proof: struct {
					TxIndex            []int `bson:"tx_index"`
					PayloadNumTxs      []int `bson:"payload_num_txs"`
					PayloadProofNumTxs struct {
						Proofs      string        `bson:"proofs"`
						PrefixBytes []interface{} `bson:"prefix_bytes"`
						SuffixBytes []int         `bson:"suffix_bytes"`
					} `bson:"payload_proof_num_txs"`
					PayloadTxTableEntries      []int `bson:"payload_tx_table_entries"`
					PayloadProofTxTableEntries struct {
						Proofs      string `bson:"proofs"`
						PrefixBytes []int  `bson:"prefix_bytes"`
						SuffixBytes []int  `bson:"suffix_bytes"`
					} `bson:"payload_proof_tx_table_entries"`
					PayloadProofTx struct {
						Proofs      string        `bson:"proofs"`
						PrefixBytes []int         `bson:"prefix_bytes"`
						SuffixBytes []interface{} `bson:"suffix_bytes"`
					} `bson:"payload_proof_tx"`
				}{
					TxIndex:       tracksEspressoStruct.EspressoData.Proof.TxIndex,
					PayloadNumTxs: tracksEspressoStruct.EspressoData.Proof.PayloadNumTxs,
					PayloadProofNumTxs: struct {
						Proofs      string        `bson:"proofs"`
						PrefixBytes []interface{} `bson:"prefix_bytes"`
						SuffixBytes []int         `bson:"suffix_bytes"`
					}{
						Proofs:      tracksEspressoStruct.EspressoData.Proof.PayloadProofNumTxs.Proofs,
						PrefixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofNumTxs.PrefixBytes,
						SuffixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofNumTxs.SuffixBytes,
					},
					PayloadTxTableEntries: tracksEspressoStruct.EspressoData.Proof.PayloadTxTableEntries,
					PayloadProofTxTableEntries: struct {
						Proofs      string `bson:"proofs"`
						PrefixBytes []int  `bson:"prefix_bytes"`
						SuffixBytes []int  `bson:"suffix_bytes"`
					}{
						Proofs:      tracksEspressoStruct.EspressoData.Proof.PayloadProofTxTableEntries.Proofs,
						PrefixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofTxTableEntries.PrefixBytes,
						SuffixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofTxTableEntries.SuffixBytes,
					},
					PayloadProofTx: struct {
						Proofs      string        `bson:"proofs"`
						PrefixBytes []int         `bson:"prefix_bytes"`
						SuffixBytes []interface{} `bson:"suffix_bytes"`
					}{
						Proofs:      tracksEspressoStruct.EspressoData.Proof.PayloadProofTx.Proofs,
						PrefixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofTx.PrefixBytes,
						SuffixBytes: tracksEspressoStruct.EspressoData.Proof.PayloadProofTx.SuffixBytes,
					},
				},
				BlockHash:   tracksEspressoStruct.EspressoData.BlockHash,
				BlockHeight: tracksEspressoStruct.EspressoData.BlockHeight,
			},
		}

		_ = mongoTracksEspressoStruct

	}
}
