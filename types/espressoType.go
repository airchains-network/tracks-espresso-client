package types

type EspressoSchemaV1 struct {
	EspressoTxResponseV1 EspressoTxResponseV1 `json:"espresso_tx_response_v_1"`
	StationID            string               `json:"station_id"`
	PodNumber            int                  `json:"pod_number"`
}

type EspressoTxResponseV1 struct {
	Transaction struct {
		Namespace int    `json:"namespace"`
		Payload   string `json:"payload"`
	} `json:"transaction"`
	Hash  string `json:"hash"`
	Index int    `json:"index"`
	Proof struct {
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
	} `json:"proof"`
	BlockHash   string `json:"block_hash"`
	BlockHeight int    `json:"block_height"`
}
