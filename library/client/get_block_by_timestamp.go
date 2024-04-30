package client

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetBlockByTimestamp(rpc string, timestamp time.Time) (int, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := client.GetBlockLatest(c.ctx)
	if err != nil {
		return 0, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
	}
	nowblock := res.Block.Header.Height
	nowblocktime := res.Block.Header.Timestamp.ToTime()

	if timestamp.After(nowblocktime) {
		return 0, errors.New("timestamp is in the future")
	}
	res2, err := client.GetBlockByHeight(c.ctx, uint64(1))
	if err != nil {
		return 0, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
	}
	if timestamp.Before(res2.Block.Header.Timestamp.ToTime()) {
		return 0, errors.New("timestamp is less than timestamp of first available block")
	}

	leftBlock := int64(0)
	rightBlock := int64(nowblock)
	foundBlock := int64(-1)

	//find block with largest timestamp smaller than provided
	for leftBlock <= rightBlock {
		midBlock := (leftBlock + rightBlock) / 2
		res, err := client.GetBlockByHeight(c.ctx, uint64(midBlock))
		if err != nil {
			return 0, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
		}
		midBlockTime := res.Block.Header.Timestamp.ToTime()

		if midBlockTime.Before(timestamp) {
			leftBlock = midBlock + 1
			foundBlock = midBlock
		} else {
			rightBlock = midBlock - 1
		}
	}

	if foundBlock != -1 {
		return int(foundBlock), nil
	}

	return 0, errors.New("cannot determine closest block")
}
