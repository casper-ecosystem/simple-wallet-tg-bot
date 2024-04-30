package client

import (
	"fmt"
	"net/http"
	"time"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetEraByTimestamp(rpc string, timestamp time.Time) (int, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	closestblock, err := c.GetBlockByTimestamp(rpc, timestamp)
	if err != nil {
		return 0, err
	}
	res, err := client.GetBlockByHeight(c.ctx, uint64(closestblock))
	if err != nil {
		return 0, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
	}
	return int(res.Block.Header.EraID), nil

}
