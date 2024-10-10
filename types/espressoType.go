package types

type TracksEspressoStruct struct {
	StationId    string             `json:"station_id"`
	EspressoVmId string             `json:"espresso_vm_id"`
	EspressoData EspressoDataStruct `json:"espresso_data"`
}

type EspressoDataStruct struct {
	Transaction struct {
		Namespace int    `json:"namespace"`
		Payload   string `json:"payload"`
	} `json:"transaction"`
	Hash  string `json:"hash"`
	Index int    `json:"index"`
	Proof struct {
		TxIndex            []int `json:"tx_index"`
		PayloadNumTxs      []int `json:"payload_num_txs"`
		PayloadProofNumTxs struct {
			Proofs      string        `json:"proofs"`
			PrefixBytes []interface{} `json:"prefix_bytes"`
			SuffixBytes []int         `json:"suffix_bytes"`
		} `json:"payload_proof_num_txs"`
		PayloadTxTableEntries      []int `json:"payload_tx_table_entries"`
		PayloadProofTxTableEntries struct {
			Proofs      string `json:"proofs"`
			PrefixBytes []int  `json:"prefix_bytes"`
			SuffixBytes []int  `json:"suffix_bytes"`
		} `json:"payload_proof_tx_table_entries"`
		PayloadProofTx struct {
			Proofs      string        `json:"proofs"`
			PrefixBytes []int         `json:"prefix_bytes"`
			SuffixBytes []interface{} `json:"suffix_bytes"`
		} `json:"payload_proof_tx"`
	} `json:"proof"`
	BlockHash   string `json:"block_hash"`
	BlockHeight int    `json:"block_height"`
}
