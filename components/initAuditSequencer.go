package components

import (
	"context"
	"fmt"

	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func InitAuditSequencer(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, submitter string, sequencerChecks []types.ExtSequencerCheck) error {
	msgAuditSequencer := &types.MsgAuditSequencer{
		Verifier:        submitter,
		SequencerChecks: sequencerChecks,
	}

	txResp, err := client.BroadcastTx(ctx, account, msgAuditSequencer)
	if err != nil {
		fmt.Println("Error broadcasting transaction:", err)
		return err
	}

	fmt.Printf("Here is the tx response of the tx : %v", txResp)

	return nil
}
