package client

import (
	"net/http"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetCurrentState(rpc string) (int, int, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))

	res, err := client.GetBlockLatest(c.ctx)
	if err != nil {
		return 0, 0, err
	}
	return int(res.Block.Header.Height), int(res.Block.Header.EraID), nil
}
