package client

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

// function to get average block time for n blocks
func (c *Client) averageBlockTime(rpc string, block_start uint64, block_end uint64) (time.Duration, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	var total_time time.Duration
	var count int

	var previousTimestamp time.Time

	for i := block_start; i <= block_end; i++ {
		res, err := client.GetBlockByHeight(c.ctx, i)
		if err != nil {
			return time.Duration(0), NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
		}

		currentTimestamp := res.Block.Header.Timestamp.ToTime()

		if i != block_start {
			total_time += currentTimestamp.Sub(previousTimestamp)
			count++
		}

		previousTimestamp = currentTimestamp
	}

	if count == 0 {
		return time.Duration(0), fmt.Errorf("not enough blocks to calculate average block time")
	}

	return total_time / time.Duration(count), nil
}

func (c *Client) GetTimestampByBlock(rpc string, block int) (time.Time, bool, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := client.GetBlockLatest(c.ctx)
	if err != nil {
		return time.Time{}, false, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
	}
	nowheight := res.Block.Header.Height
	if uint64(block) == nowheight {
		return res.Block.Header.Timestamp.ToTime(), false, nil
	}
	if uint64(block) > nowheight {
		diff := uint64(block) - nowheight
		if diff > 1000 {
			return time.Time{}, true, NewBlockTooFarInFutureError("block number is too big")
		}
		avg, err := c.averageBlockTime(rpc, nowheight-10, nowheight)
		if err != nil {
			return time.Time{}, true, err
		}
		return res.Block.Header.Timestamp.ToTime().Add(avg), true, nil
	}
	if uint64(block) < nowheight {
		res, err := client.GetBlockByHeight(c.ctx, uint64(block))
		if err != nil {
			return time.Time{}, false, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
		}
		return res.Block.Header.Timestamp.ToTime(), false, nil
	}
	return time.Time{}, false, errors.New("unknown error")
}
