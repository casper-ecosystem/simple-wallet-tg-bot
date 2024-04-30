package client

import (
	"fmt"
	"math/big"
	"net/http"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetBalanceStaked(rpc string, address string) (map[string]big.Int, int, error) { //TODO: caching
	balance := make(map[string]big.Int)
	if !c.IsAddress(address) {
		return balance, CasperDecimals, NewRpcAdressInvalidError(fmt.Sprintf("address %s is invalid", address))
	}
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := client.GetAuctionInfoLatest(c.ctx)
	if err != nil {
		return balance, CasperDecimals, err
	}

	for _, bid := range res.AuctionState.Bids {
		staked := big.Int{}
		for _, delegator := range bid.Bid.Delegators {
			if delegator.PublicKey.String() == address {
				amount := delegator.StakedAmount.Value()
				staked.Add(&staked, amount)
			}
		}
		if staked.Cmp(big.NewInt(0)) == 1 {
			balance[bid.PublicKey.String()] = staked
		}
	}

	return balance, CasperDecimals, err
}
