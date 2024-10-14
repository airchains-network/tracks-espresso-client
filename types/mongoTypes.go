package types

type MongoTracksEspressoStruct struct {
	StationId    string                  `bson:"station_id"`
	EspressoData MongoEspressoDataStruct `bson:"espresso_data"`
}

type MongoEspressoDataStruct struct {
	Transaction struct {
		Namespace int    `bson:"namespace"`
		Payload   string `bson:"payload"`
	} `bson:"transaction"`
	Hash  string `bson:"hash"`
	Index int    `bson:"index"`
	Proof struct {
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
	} `bson:"proof"`
	BlockHash   string `bson:"block_hash"`
	BlockHeight int    `bson:"block_height"`
}
