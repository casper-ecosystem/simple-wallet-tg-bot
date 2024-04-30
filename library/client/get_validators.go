package client

import (
	"net/http"

	"github.com/Simplewallethq/library/blockchain"
	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetValidators(rpc string) (blockchain.ValidatorsResponse, error) {
	resp := blockchain.ValidatorsResponse{}
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := client.GetAuctionInfoLatest(c.ctx)
	if err != nil {
		return resp, err
	}
	for _, bid := range res.AuctionState.Bids {
		resp.Validators = append(resp.Validators, blockchain.Validator{
			Address: bid.PublicKey.String(),
			Fee:     bid.Bid.DelegationRate,
			Active:  !bid.Bid.Inactive,
		})
	}
	resp.Decimals = CasperDecimals
	return resp, nil

}
