package client

import (
	"log"
	"net/http"

	"github.com/Simplewallethq/library/blockchain"
	newsdk "github.com/make-software/casper-go-sdk/casper"
	sdkrpc "github.com/make-software/casper-go-sdk/rpc"
)

func (c *Client) GetBalanceBeingUnstaked(rpc string, address string) ([]blockchain.UnstakingBalance, int, error) {
	result := []blockchain.UnstakingBalance{}
	newclient := newsdk.NewRPCClient(newsdk.NewRPCHandler(rpc, http.DefaultClient))

	latestBlockHeight, nowEra, err := c.GetCurrentState(rpc)
	if err != nil {
		return result, CasperDecimals, err
	}
	log.Println(nowEra)

	deployHashes, err := c.collectDeployHashesByEra(newclient, latestBlockHeight, nowEra, nowEra-UndelegateDelay)
	if err != nil {
		return result, CasperDecimals, err
	}
	//log.Println("deploy hashes:", deployHashes)

	UserWithdraws, err := c.collectUserWithdrawsUndelegate(newclient, deployHashes, address)
	if err != nil {
		return result, CasperDecimals, err
	}
	//log.Println("user withdraws:", UserWithdraws)
	nowblock, err := newclient.GetBlockLatest(c.ctx)
	if err != nil {
		return result, CasperDecimals, err
	}

	result, err = c.processUserWithdrawsUnstaking(newclient, UserWithdraws, nowblock.Block.Hash.String())
	if err != nil {
		return result, CasperDecimals, err
	}

	return result, CasperDecimals, nil
}

func (c *Client) collectDeployHashesByEra(client sdkrpc.Client, latestBlockHeight int, nowEra int, eraEnd int) ([]UserDeploysData, error) {
	deployHashes := []UserDeploysData{}
	currentEra := nowEra
	currentHeight := latestBlockHeight
	for currentEra != eraEnd {
		block, err := client.GetBlockByHeight(c.ctx, uint64(currentHeight))
		if err != nil {
			return deployHashes, err
		}
		currentEra = int(block.Block.Header.EraID)
		currentHeight = int(block.Block.Header.Height - 1)
		for _, dephash := range block.Block.Body.DeployHashes {
			deployHashes = append(deployHashes, UserDeploysData{Hash: dephash.String(), Height: uint64(currentHeight)})
		}
	}
	return deployHashes, nil
}

type UncheckedWithdraws struct {
	amount       string
	withdrawHash string
	validator    string
	unbounder    string
	era          int
	height       uint64
}

func (c *Client) collectUserWithdrawsUndelegate(client sdkrpc.Client, deployHashes []UserDeploysData, address string) ([]UncheckedWithdraws, error) {
	UserWithdraws := []UncheckedWithdraws{}

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
		if deploy.Deploy.Session.StoredContractByHash.EntryPoint != "undelegate" {
			continue
		}
		for _, execres := range deploy.ExecutionResults {
			if execres.Result.Success == nil {
				continue
			}
			for _, transform := range execres.Result.Success.Effect.Transforms {
				if !transform.Transform.IsWriteWithdraw() {
					continue
				}
				with, err := transform.Transform.ParseAsWriteWithdraws()
				if err != nil {
					continue
				}
				for _, tran := range with {
					if tran.UnbonderPublicKey.String() != address {
						continue
					}
					//log.Println("hash", transform.Key.Withdraw.String())
					var promres = UncheckedWithdraws{}
					promres.withdrawHash = "withdraw-" + transform.Key.Withdraw.String()
					promres.validator = tran.ValidatorPublicKey.String()
					promres.unbounder = tran.UnbonderPublicKey.String()
					promres.amount = tran.Amount.Value().String()
					promres.era = int(tran.EraOfCreation)
					promres.height = depHash.Height
					UserWithdraws = append(UserWithdraws, promres)
				}
			}
		}
	}

	return UserWithdraws, nil
}

func (c *Client) processUserWithdrawsUnstaking(client sdkrpc.Client, UserWithdraws []UncheckedWithdraws, blockhash string) ([]blockchain.UnstakingBalance, error) {
	result := []blockchain.UnstakingBalance{}

	for _, withdraw := range UserWithdraws {
		res, err := client.QueryGlobalStateByBlockHash(c.ctx,
			blockhash,
			withdraw.withdrawHash,
			[]string{},
		)
		if err != nil {
			return result, err
		}
		//log.Println("res:", res.StoredValue.Withdraw)
		if len(res.StoredValue.Withdraw) == 0 {
			continue
		}
		for _, with := range res.StoredValue.Withdraw {
			//log.Println("res.StoredValue.Withdraw:", with)
			if with.UnbonderPublicKey.String() != withdraw.unbounder {
				continue
			}
			log.Println("res.StoredValue.Withdraw:", with)

			result = append(result, blockchain.UnstakingBalance{
				Amount:          with.Amount.Value().String(),
				ValidatorPubkey: with.ValidatorPublicKey.String(),
				Era:             int(with.EraOfCreation) + UndelegateDelay,
			})

		}
	}
	return result, nil
}
