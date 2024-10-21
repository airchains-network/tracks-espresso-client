package batches

import (
	"context"
	"log"
	"strconv"
	

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/types"
	junctiontypes  "github.com/airchains-network/junction/x/trackgate/types"
)

func Batch(ctx context.Context, mongo database.DB) {
    // Set the MongoDB collection and context
    collection := mongo.Collection("espresso-data")

    // Define the projection to only retrieve specific fields
	projection := bson.D{
		{Key: "station_id", Value: 1},
		{Key: "pod_number", Value: 1},
		{Key: "espresso_tx_response_v_1", Value: 1},
	}
	
	// fmt.Print(projection)
    // Find documents in MongoDB with projection
    cursor, err := collection.Find(ctx, bson.D{}, options.Find().SetProjection(projection))
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)

	var batch []types.MongoTracksEspressoStruct
	if batchErr := cursor.All(ctx, &batch); batchErr != nil {
		// return nil, nil, nil, fmt.Errorf("failed to decode vote data: %v", batchErr)
	}

	var seqcheck  []junctiontypes.ExtSequencerCheck
	for _ , b := range batch {

		pn := strconv.Itoa(b.PodNumber)
		ns := strconv.Itoa(b.EspressoTxResponseV1.Transaction.Namespace)

		seqcheck = append(seqcheck, junctiontypes.ExtSequencerCheck{
			PodNumber  : pn ,
	ExtTrackStationId :b.StationId ,
	VerificationStatus : true,
	Namespace : ns ,
		})
	}



}
