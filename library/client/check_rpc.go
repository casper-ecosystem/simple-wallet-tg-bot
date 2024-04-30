package client

import (
	"net/http"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) CheckChain(rpc string, chain string) (bool, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := client.GetStatus(c.ctx)
	if err != nil {
		return false, err
	}
	return res.ChainSpecName == chain, nil
}
