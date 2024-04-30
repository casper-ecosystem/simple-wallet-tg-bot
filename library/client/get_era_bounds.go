package client

import (
	"errors"
	"fmt"
	"net/http"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetEraBounds(rpc string, era int) (int64, int64, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := client.GetBlockLatest(c.ctx)
	if err != nil {
		return 0, 0, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
	}
	nowblock := res.Block.Header.Height
	nowEra := res.Block.Header.EraID

	if era > int(nowEra) {
		return 0, 0, errors.New("requested era is in the future")
	}
	if era == int(nowEra) {
		return 0, 0, errors.New("requested era is current")
	}

	avgBlocksPerEra := 220
	estimatedStartBlock := int64(nowblock) - int64(int(nowEra)-era)*int64(avgBlocksPerEra)

	leftBlock := int64(0)
	if estimatedStartBlock > 0 {
		leftBlock = estimatedStartBlock
	}
	rightBlock := int64(nowblock)
	firstBlock := int64(-1)

	// Find the first block
	for leftBlock <= rightBlock {
		midBlock := (leftBlock + rightBlock) / 2
		res, err := client.GetBlockByHeight(c.ctx, uint64(midBlock))
		if err != nil {
			if err.Error() == "block not found" {
				return 0, 0, errors.New("cannot determine era bounds")
			}
			return 0, 0, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
		}
		midBlockEra := int(res.Block.Header.EraID)

		if midBlockEra < era {
			leftBlock = midBlock + 1
		} else {
			if midBlockEra == era {
				firstBlock = midBlock
			}
			rightBlock = midBlock - 1
		}
	}

	if firstBlock == -1 {
		return 0, 0, errors.New("cannot determine era bounds")
	}

	// Find the last block
	leftBlock = firstBlock
	//rightBlock = int64(nowblock)
	rightBlock = int64(nowblock) + int64(avgBlocksPerEra)
	lastBlock := int64(-1)

	for leftBlock <= rightBlock {
		midBlock := (leftBlock + rightBlock) / 2
		res, err := client.GetBlockByHeight(c.ctx, uint64(midBlock))
		if err != nil {
			return 0, 0, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
		}
		midBlockEra := int(res.Block.Header.EraID)

		if midBlockEra > era {
			rightBlock = midBlock - 1
		} else {
			if midBlockEra == era {
				lastBlock = midBlock
			}
			leftBlock = midBlock + 1
		}
	}

	return firstBlock, lastBlock, nil
}
