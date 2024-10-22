package components

import (
	"context"
	"fmt"
	"log"

	"github.com/airchains-network/tracks-espresso-client/config"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

// CheckIfAccountExists verifies if an account exists in the registry and returns its address if it does.
func CheckIfAccountExists(accountName string, client cosmosclient.Client, addressPrefix string, accountPath string) (bool, string) {
	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		log.Printf("Error creating account registry: %v", err)
		return false, ""
	}

	account, err := registry.GetByName(accountName)
	if err != nil {
		log.Printf("Account %s not found: %v", accountName, err)
		return false, ""
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Printf("Failed to get the address for account %s: %v", accountName, err)
		return false, ""
	}

	return true, addr
}

// CreateAccount creates a new account in the registry and prints its address and mnemonic.
func CreateAccount(accountName string, accountPath string) {
	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		log.Printf("Error creating account registry: %v", err)
		return
	}

	account, mnemonic, err := registry.Create(accountName)
	if err != nil {
		log.Printf("Error creating account %s: %v", accountName, err)
		return
	}

	accountAddr, err := account.Address("air")
	if err != nil {
		log.Printf("Error retrieving address for account %s: %v", accountName, err)
		return
	}

	log.Printf("Rollup Account Created: %s", accountAddr)
	log.Printf("Mnemonic: %s", mnemonic)
}

func SetupAccountClient(ctx context.Context, gasFees string) (*cosmosclient.Client, string, cosmosaccount.Account, error) {
	const (
		addressPrefix = "air"
		accountPath   = "./accounts"
		accountName   = "charlie"
	)

	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		return nil, "", cosmosaccount.Account{}, fmt.Errorf("error creating account registry: %v", err)
	}

	newTempAccount, err := registry.GetByName(accountName)
	if err != nil {
		return nil, "", cosmosaccount.Account{}, fmt.Errorf("error getting account: %v", err)
	}

	newTempAddr, err := newTempAccount.Address(addressPrefix)
	if err != nil {
		return nil, "", cosmosaccount.Account{}, fmt.Errorf("error getting address: %v", err)
	}

	accountClient, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(config.TendermintApiUrl), cosmosclient.WithHome(accountPath), cosmosclient.WithGas("auto"), cosmosclient.WithFees(gasFees))
	if err != nil {
		return nil, "", cosmosaccount.Account{}, fmt.Errorf("error creating account client: %v", err)
	}

	return &accountClient, newTempAddr, newTempAccount, nil
}