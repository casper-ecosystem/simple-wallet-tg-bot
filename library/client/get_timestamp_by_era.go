package client

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

// We need using local database to get correct information
func (c *Client) GetTimestampByEra(rpc string, era int) (time.Time, bool, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	nowblock, nowera, err := c.GetCurrentState(rpc)
	if err != nil {
		return time.Time{}, false, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
	}
	avgBlockTime, err := c.averageBlockTime(rpc, uint64(nowblock-10), uint64(nowblock))
	if err != nil {
		return time.Time{}, false, err
	}
	AverageBlocksInEra := 220
	AverageEraTime := time.Duration(int64(AverageBlocksInEra) * int64(avgBlockTime)) //fixme: this is not correct

	if err != nil {
		return time.Time{}, false, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
	}
	if era == nowera {
		res, err := client.GetBlockByHeight(c.ctx, uint64(nowblock))
		if err != nil {
			return time.Time{}, false, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
		}
		return res.Block.Header.Timestamp.ToTime(), false, nil
	}
	if era > nowera {
		diff := era - nowera
		if diff > 100 {
			return time.Time{}, false, NewEraTooFarInFutureError("era number is too big")
		}
		estimatedTime := time.Now().Add(time.Duration(diff) * AverageEraTime)
		return estimatedTime, true, nil
	}
	if era < nowera {
		leftBlock := uint64(0)
		rightBlock := uint64(nowblock)
		for leftBlock <= rightBlock {
			midBlock := (leftBlock + rightBlock) / 2
			res, err := client.GetBlockByHeight(c.ctx, (midBlock))
			if err != nil {
				return time.Time{}, false, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
			}
			if int(res.Block.Header.EraID) == era {
				return res.Block.Header.Timestamp.ToTime(), false, nil
			}
			if int(res.Block.Header.EraID) < era {
				leftBlock = midBlock + 10 //min size of era 10 blocks
			} else {
				rightBlock = midBlock - 10
			}
		}
	}

	return time.Time{}, false, errors.New("unknown error")
}
