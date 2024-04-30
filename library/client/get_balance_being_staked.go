package client

import (
	"net/http"

	"github.com/Simplewallethq/library/blockchain"
	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetBalanceBeingStaked(rpc string, address string) ([]blockchain.BeingStaked, int, error) {
	var result []blockchain.BeingStaked
	// client := sdk.NewRpcClient(rpc)
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	nowblock, nowera, err := c.GetCurrentState(rpc)
	if err != nil {
		return result, CasperDecimals, err
	}
	PrevEra := nowera - 1
	startblock, _, err := c.GetEraBounds(rpc, PrevEra)
	if err != nil {
		return result, CasperDecimals, err
	}
	deploys, err := c.collectDeployHashesByBlocks(client, int(startblock), int(nowblock))
	if err != nil {
		return result, CasperDecimals, err
	}
	delegates, err := c.collectUserDelegates(client, deploys, address)
	if err != nil {
		return result, CasperDecimals, err
	}
	for _, delegate := range delegates {
		result = append(result, blockchain.BeingStaked{
			Amount:          delegate.amount,
			ValidatorPubkey: delegate.validator,
			Era:             delegate.era + 2,
		})
	}

	return result, CasperDecimals, nil
}
