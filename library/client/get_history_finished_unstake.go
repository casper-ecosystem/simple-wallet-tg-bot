package client

import (
	"errors"
	"net/http"

	"github.com/Simplewallethq/library/blockchain"
	newsdk "github.com/make-software/casper-go-sdk/casper"
	sdkrpc "github.com/make-software/casper-go-sdk/rpc"
)

func (c *Client) GetHistoryUndelegate(rpc string, address string, blockStart int, blockEnd int) ([]blockchain.HistroyUnstake, int, error) {
	result := []blockchain.HistroyUnstake{}
	if blockStart-blockEnd > 1000 {
		return result, CasperDecimals, errors.New("blockEnd - blockStart > 1000, max 1000 blocks")
	}
	newclient := newsdk.NewRPCClient(newsdk.NewRPCHandler(rpc, http.DefaultClient))
	deployHashes, err := c.collectDeployHashesByBlocks(newclient, blockStart, blockEnd)
	if err != nil {
		return result, CasperDecimals, err
	}
	//log.Println("deploy hashes:", deployHashes)
	withdraws, err := c.collectUserWithdrawsUndelegate(newclient, deployHashes, address)
	if err != nil {
		return result, CasperDecimals, err
	}
	//log.Println("withdraws:", withdraws)
	nowblock, err := newclient.GetBlockLatest(c.ctx)
	if err != nil {
		return result, CasperDecimals, err
	}
	for _, withdraw := range withdraws {
		tempres := blockchain.HistroyUnstake{
			Amount:          withdraw.amount,
			ValidatorPubkey: withdraw.validator,
			Era:             withdraw.era,
			Finished:        false,
			Height:          withdraw.height,
		}
		res, err := newclient.QueryGlobalStateByBlockHash(c.ctx,
			nowblock.Block.Hash.String(),
			withdraw.withdrawHash,
			[]string{},
		)
		if err != nil {
			return result, CasperDecimals, err
		}
		if len(res.StoredValue.Withdraw) == 0 {
			tempres.Finished = true
		}
		result = append(result, tempres)
	}

	return result, CasperDecimals, nil
}

type UserDeploysData struct {
	Hash   string
	Height uint64
}

func (c *Client) collectDeployHashesByBlocks(client sdkrpc.Client, block_start int, block_end int) ([]UserDeploysData, error) {
	deployHashes := []UserDeploysData{}
	currentHeight := block_start
	for currentHeight != block_end+1 {
		block, err := client.GetBlockByHeight(c.ctx, uint64(currentHeight))
		if err != nil {
			return deployHashes, err
		}
		currentHeight = int(block.Block.Header.Height) + 1
		for _, dephash := range block.Block.Body.DeployHashes {
			deployHashes = append(deployHashes, UserDeploysData{Hash: dephash.String(), Height: uint64(currentHeight)})
		}
	}

	return deployHashes, nil
}
