package client

import (
	"errors"
	"net/http"

	"github.com/Simplewallethq/library/blockchain"
	sdk "github.com/make-software/casper-go-sdk/casper"
	sdkrpc "github.com/make-software/casper-go-sdk/rpc"
)

func (c *Client) GetHistoryDelegate(rpc string, address string, blockStart int, blockEnd int) ([]blockchain.HistroyStake, int, error) {
	var result []blockchain.HistroyStake
	if blockStart-blockEnd > 1000 {
		return result, CasperDecimals, errors.New("blockStart-blockEnd > 1000, max 1000 blocks")
	}
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	deployHashes, err := c.collectDeployHashesByBlocks(client, blockStart, blockEnd)
	if err != nil {
		return result, CasperDecimals, err
	}
	//log.Println(deployHashes)

	delegates, err := c.collectUserDelegates(client, deployHashes, address)
	if err != nil {
		return result, CasperDecimals, err
	}
	//log.Println("delegates", delegates)
	nowblockres, err := client.GetBlockLatest(c.ctx)
	if err != nil {
		return result, CasperDecimals, err
	}

	for _, delegate := range delegates {
		var finished bool
		if int(nowblockres.Block.Header.EraID) > delegate.era {
			finished = true
		} else {
			finished = false
		}
		result = append(result, blockchain.HistroyStake{
			Amount:          delegate.amount,
			ValidatorPubkey: delegate.validator,
			Era:             delegate.era,
			Finished:        finished,
			Height:          delegate.height,
		})
	}

	return result, CasperDecimals, nil
}

type Delegates struct {
	validator string
	amount    string
	delegator string
	era       int
	height    uint64
}

func (c *Client) collectUserDelegates(client sdkrpc.Client, deployHashes []UserDeploysData, address string) ([]Delegates, error) {
	UserWithdraws := []Delegates{}

	for _, depHash := range deployHashes {
		deploy, err := client.GetDeploy(c.ctx, depHash.Hash)
		if err != nil {
			continue
		}
		if deploy.Deploy.Header.Account.String() != address {
			continue
		}
		if deploy.Deploy.Session.StoredContractByHash == nil {
			continue
		}
		if deploy.Deploy.Session.StoredContractByHash.EntryPoint != "delegate" {
			continue
		}
		for _, execres := range deploy.ExecutionResults {
			promres := Delegates{}
			if execres.Result.Success == nil {
				continue
			}
			block, err := client.GetBlockByHash(c.ctx, execres.BlockHash.String())
			if err != nil {
				return UserWithdraws, err
			}
			promres.era = int(block.Block.Header.EraID)

			//fmt.Println(deploy.Deploy.Session.StoredContractByHash.Args)
			delegator, err := deploy.Deploy.Session.StoredContractByHash.Args.Find("delegator")
			if err != nil {
				continue
			} else {
				value, err := delegator.Value()
				if err != nil {
					continue
				}
				promres.delegator = value.String()
			}
			validator, err := deploy.Deploy.Session.StoredContractByHash.Args.Find("validator")
			if err != nil {
				continue
			} else {
				value, err := validator.Value()
				if err != nil {
					continue
				}
				promres.validator = value.String()
			}
			amount, err := deploy.Deploy.Session.StoredContractByHash.Args.Find("amount")
			if err != nil {
				continue
			} else {
				value, err := amount.Value()
				if err != nil {
					continue
				}
				promres.amount = value.String()
			}
			promres.height = depHash.Height
			UserWithdraws = append(UserWithdraws, promres)
		}
	}
	return UserWithdraws, nil
}
