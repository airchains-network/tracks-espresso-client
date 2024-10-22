package batches

import (
	"context"
	// "fmt"
	"log"

	// "math/rand"
	"strconv"
	"time"

	junctiontypes "github.com/airchains-network/junction/x/trackgate/types"
	"github.com/airchains-network/tracks-espresso-client/components"
	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Constants for Cosmos account information
const (
	addressPrefix = "air"
	accountPath   = "./accounts"
	accountName   = "charlie"
)

// Batch handles batch operations including fetching data from MongoDB, constructing the sequencer check,
// and initiating the sequencer audit. It also contains a placeholder for pruning logic.
func Batch(ctx context.Context, mongo database.DB, client cosmosclient.Client, account cosmosaccount.Account) {
	// func Batch( ctx context.Context , mongo database.DB) {

	// Set the MongoDB collection and context
	collection := mongo.Collection("espresso-data")

	// Define the projection to only retrieve specific fields
	projection := bson.D{
		{Key: "station_id", Value: 1},
		{Key: "pod_number", Value: 1},
		{Key: "espresso_tx_response_v_1", Value: 1},
	}

	option := options.Find().SetProjection(projection).
	SetLimit(100).
	SetSort(bson.D{{Key :"created_at", Value: 1}})

	// Find documents in MongoDB with projection
	cursor, err := collection.Find(ctx, bson.D{}, option)
	if err != nil {
		log.Fatalf("Failed to fetch data from MongoDB: %v", err)
	}
	defer cursor.Close(ctx)

	// Store the fetched data into batch slice
	var batch []types.MongoTracksEspressoStruct
	if batchErr := cursor.All(ctx, &batch); batchErr != nil {
		log.Fatalf("Failed to decode data: %v", batchErr)
	}

	// Construct the sequencer check based on the fetched data
	var seqcheck []junctiontypes.ExtSequencerCheck
	for _, b := range batch {
		// Convert pod_number and namespace to string
		pn := strconv.Itoa(b.PodNumber)
		ns := strconv.Itoa(b.EspressoTxResponseV1.Transaction.Namespace)

		// Append to seqcheck
		seqcheck = append(seqcheck, junctiontypes.ExtSequencerCheck{
			PodNumber:          pn,
			ExtTrackStationId:  b.StationId,
			VerificationStatus: true, // Assuming verification status is true
			Namespace:          ns,
		})
	}

	// Get the submitter address from the Cosmos account
	submitterAddress, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatalf("Failed to get the address for account %s: %v", accountName, err)
	}

	// Initialize the response check flag
	responseCheck := false

	// Start the sequencer audit loop
	for {
		// Call the InitAuditSequencer function to initiate the sequencer audit
		auditSequencerResponseErr := components.InitAuditSequencer(client, ctx, account, submitterAddress, seqcheck)
		if auditSequencerResponseErr != nil {
			log.Printf("Failed to init auditSequencerResponseErr: %v", auditSequencerResponseErr)
			time.Sleep(time.Second * 3) // Retry after 3 seconds
			log.Printf("Retrying after 3sec...")
			continue
		} else {
			//pruning logic 
			var idsToDelete []string

        // Check each sequencer in seqcheck for verification status
        for i, check := range seqcheck {
            if check.VerificationStatus {
                // If VerificationStatus is true, prepare to delete this entry
				uniqueId := check.ExtTrackStationId + check.PodNumber 
                idsToDelete = append(idsToDelete, uniqueId)
                // Optionally remove the entry from seqcheck
                seqcheck = append(seqcheck[:i], seqcheck[i+1:]...) // Remove the entry from seqcheck
            }
        }

        // Delete the corresponding records from MongoDB
        if len(idsToDelete) > 0 {
            // Assuming ExtTrackStationId is a unique identifier for the records
            _, err := collection.DeleteMany(ctx, bson.D{{Key: "EspressoStationId", Value: bson.D{{Key: "$in", Value: idsToDelete}}}})
            if err != nil {
                log.Printf("Failed to delete records from MongoDB: %v", err)
            } else {
                log.Printf("Deleted %d records from MongoDB", len(idsToDelete))
				responseCheck = true
            }
        }

		}

		//If responseCheck is set to true, break the loop
		if responseCheck {
			break
		}
	}
	// fmt.Println("This is batch",batch)
	// fmt.Println("this is sequence" ,seqcheck)
}

