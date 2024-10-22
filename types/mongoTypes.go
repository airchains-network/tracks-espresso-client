package types

import (
	"time"
)

type MongoTracksEspressoStruct struct {
	EspressoStationId    string               `json:"espresso_station_id"`
	EspressoTxResponseV1 EspressoTxResponseV1 `json:"espresso_tx_response_v_1"`
	StationId            string               `json:"station_id"`
	PodNumber            int                  `json:"pod_number"`
	CreatedAt            time.Time                `json:"created_at"`
}
