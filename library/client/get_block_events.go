package client

import (
	"net/http"
	"strconv"

	"github.com/Simplewallethq/source-code/library/blockchain"
	sdk "github.com/make-software/casper-go-sdk/casper"
	sdkrpc "github.com/make-software/casper-go-sdk/rpc"
)

func (c *Client) GetBlockAllEvents(rpc string, height uint64, transfers bool, delegates bool, undelegates bool, rewards bool) (blockchain.BlockEvents, error) {
	result := blockchain.BlockEvents{}
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	block_info, err := client.GetBlockByHeight(c.ctx, height)
	if err != nil {
		return result, err
	}
	result.Era = block_info.Block.Header.EraID
	result.Decimals = CasperDecimals
	if transfers {
		transfers, _, err := c.GetBlockAllTransfers(rpc, height)
		if err != nil {
			return result, err
		}
		result.Transfers = transfers
	}
	if delegates || undelegates {
		deployHashes, err := c.collectDeployHashesByBlocks(client, int(height), int(height))
		if err != nil {
			return result, err
		}
		if delegates {
			delegates, _, err := c.GetBlockAllDelegates(rpc, height, deployHashes)
			if err != nil {
				return result, err
			}
			result.Delegates = delegates
		}
		if undelegates {
			undelegates, _, err := c.GetBlockAllUndelegate(rpc, height, deployHashes)
			if err != nil {
				return result, err
			}
			result.Undelegates = undelegates
		}

	}
	if rewards {
		rewards, _, err := c.GetBlockAllRewards(rpc, height)
		if err != nil {
			return result, err
		}
		result.Rewards = rewards
	}
	result.Date = block_info.Block.Header.Timestamp.ToTime()
	return result, nil
}

func (c *Client) GetBlockAllTransfers(rpc string, height uint64) ([]blockchain.TransferResponse, int, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	result := []blockchain.TransferResponse{}

	blockTransfers, err := client.GetBlockTransfersByHeight(c.ctx, height)
	if err != nil {
		return []blockchain.TransferResponse{}, 0, err
	}
	for _, transfer := range blockTransfers.Transfers {
		var frompubkey, topubkey string
		var to string
		if transfer.To != nil {
			to = transfer.To.String()
		} else {
			to = topubkey
		}

		deploy, err := client.GetDeploy(c.ctx, transfer.DeployHash.String())
		if err != nil {
			return []blockchain.TransferResponse{}, 0, err
		}
		frompubkey = deploy.Deploy.Header.Account.String()
		if deploy.Deploy.Session.Transfer != nil {
			target, err := deploy.Deploy.Session.Transfer.Args.Find("target")
			if err != nil {
				topubkey = ""
			} else {
				targetParsed, err := target.Value()
				if err != nil {
					return []blockchain.TransferResponse{}, 0, err
				}

				topubkey = targetParsed.String()
			}

		}
		result = append(result, blockchain.TransferResponse{
			DeployHash: transfer.DeployHash.String(),
			From:       transfer.From.String(),
			Memo:       transfer.ID,
			FromPubKey: frompubkey,
			To:         to,
			ToPubKey:   topubkey,
			Amount:     transfer.Amount.Value().String(),
			Gas:        strconv.FormatUint(uint64(transfer.Gas), 10),
		})

	}
	return result, CasperDecimals, nil
}

func (c *Client) GetBlockAllDelegates(rpc string, height uint64, deployHashes []UserDeploysData) ([]blockchain.DelegateData, int, error) {
	var result []blockchain.DelegateData
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))

	delegates, err := c.collectUserDelegatesAll(client, deployHashes)
	if err != nil {
		return result, CasperDecimals, err
	}
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
		result = append(result, blockchain.DelegateData{
			Address:         delegate.delegator,
			Amount:          delegate.amount,
			ValidatorPubkey: delegate.validator,
			Era:             delegate.era,
			Finished:        finished,
		})
	}
	return result, CasperDecimals, nil
}

func (c *Client) GetBlockAllUndelegate(rpc string, height uint64, deployHashes []UserDeploysData) ([]blockchain.UndelegateData, int, error) {
	result := []blockchain.UndelegateData{}
	newclient := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))

	undelegates, err := c.collectUserUndelegateAll(newclient, deployHashes)
	if err != nil {
		return result, CasperDecimals, err
	}

	for _, withdraw := range undelegates {
		tempres := blockchain.UndelegateData{
			Address:         withdraw.delegator,
			Amount:          withdraw.amount,
			ValidatorPubkey: withdraw.validator,
			Era:             withdraw.era,
			Finished:        false,
		}
		result = append(result, tempres)
	}

	return result, CasperDecimals, nil
}

func (c *Client) collectUserDelegatesAll(client sdkrpc.Client, deployHashes []UserDeploysData) ([]Delegates, error) {
	UserWithdraws := []Delegates{}

	for _, depHash := range deployHashes {
		deploy, err := client.GetDeploy(c.ctx, depHash.Hash)
		if err != nil {
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
			UserWithdraws = append(UserWithdraws, promres)
		}
	}
	return UserWithdraws, nil
}

func (c *Client) collectUserUndelegateAll(client sdkrpc.Client, deployHashes []UserDeploysData) ([]Delegates, error) {
	UserWithdraws := []Delegates{}

	for _, depHash := range deployHashes {
		deploy, err := client.GetDeploy(c.ctx, depHash.Hash)
		if err != nil {
			continue
		}
		if deploy.Deploy.Session.StoredContractByHash == nil {
			continue
		}
		if deploy.Deploy.Session.StoredContractByHash.EntryPoint != "undelegate" {
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
			UserWithdraws = append(UserWithdraws, promres)
		}
	}
	return UserWithdraws, nil
}

func (c *Client) GetBlockAllRewards(rpc string, height uint64) ([]blockchain.RewardsData, int, error) {
	result := []blockchain.RewardsData{}
	newclient := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := newclient.GetEraInfoByBlockHeight(c.ctx, uint64(height))
	if err != nil {
		return result, CasperDecimals, err
	}

	if res.EraSummary.StoredValue.EraInfo == nil {
		return result, CasperDecimals, nil
	}

	for _, reward := range res.EraSummary.StoredValue.EraInfo.SeigniorageAllocations {
		if reward.Delegator == nil {
			continue
		}
		result = append(result, blockchain.RewardsData{
			Delagator: reward.Delegator.DelegatorPublicKey.String(),
			Validator: reward.Delegator.ValidatorPublicKey.String(),
			Amount:    reward.Delegator.Amount.Value(),
		})
	}

	return result, CasperDecimals, nil
}
