package pruning

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/airchains-network/tracks-espresso-client/config"
	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/server/espresso"
	"github.com/deadlium/deadlogs"
)

var mu sync.Mutex

// PruneAndStore handles pruning the data from the JSON file and storing it in MongoDB
func PruneAndStore(ctx context.Context, filePath string, dbInstance *database.DB) error {
	mu.Lock()
	defer mu.Unlock() // Ensure the lock is released even if an error occurs

	deadlogs.Info("Pruning and storing data...")

	// Load the data from the JSON file and insert it into MongoDB
	if _, err := espresso.LoadDataFromFile(config.FilePath, dbInstance); err != nil {
		return fmt.Errorf("error during data load and insertion: %s", err)
	}

	// After successful insertion, clear the contents of the JSON file
	if err := clearFile(filePath); err != nil {
		return fmt.Errorf("error clearing file: %s", err)
	}

	deadlogs.Info("Data has been stored in MongoDB and pruned from file.")
	return nil
}

// clearFile removes the content from the specified JSON file
func clearFile(filePath string) error {
	// Truncate the file to clear its contents
	if err := os.Truncate(filePath, 0); err != nil {
		return fmt.Errorf("failed to clear the file: %s", err)
	}
	// Write an empty array to main.json to ensure it is valid JSON
	return os.WriteFile(filePath, []byte("[]"), 0644)
}

// StartPruningScheduler runs the pruning process every 30 minutes
func StartPruningScheduler(ctx context.Context, filePath string, dbInstance *database.DB) {
	for {
		select {
		case <-ctx.Done():
			deadlogs.Info("Pruning scheduler stopped.")
			return
		default:
			time.Sleep(10 * time.Minute) // Adjusted to 30 minutes for the actual implementation
			deadlogs.Info("Starting prune and store process...")
			if err := PruneAndStore(ctx, filePath, dbInstance); err != nil {
				deadlogs.Error(fmt.Sprintf("Error during prune and store: %s", err))
			} else {
				deadlogs.Success("Prune and store completed successfully")
			}
		}
	}
}
