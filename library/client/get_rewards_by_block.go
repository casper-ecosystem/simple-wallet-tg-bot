package client

import (
	"fmt"
	"net/http"

	"github.com/Simplewallethq/source-code/library/blockchain"
	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetRewardsByBlock(rpc string, address string, blockStart int64, blockEnd int64) ([]blockchain.RewardsByBlock, int, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	rewards := []blockchain.RewardsByBlock{}
	for i := blockStart; i <= blockEnd; i++ {
		res, err := client.GetEraInfoByBlockHeight(c.ctx, uint64(i))
		if err != nil {
			return rewards, CasperDecimals, err
		}
		if res.EraSummary.StoredValue.EraInfo == nil {
			continue
		}

		for _, reward := range res.EraSummary.StoredValue.EraInfo.SeigniorageAllocations {
			//log.Println(reward)
			if reward.Delegator == nil {
				continue
			}
			if reward.Delegator.DelegatorPublicKey.String() == address {
				amount := reward.Delegator.Amount.Value()
				rewards = append(rewards, blockchain.RewardsByBlock{
					Amount:          amount,
					ValidatorPubkey: reward.Delegator.ValidatorPublicKey.String(),
					Block:           uint64(i),
				})
			}
		}
	}
	return rewards, CasperDecimals, nil

}

func (c *Client) GetRewardsByEra(rpc string, address string, era_start int, era_end int) ([]blockchain.RewardsByEra, int, error) {
	rewards := []blockchain.RewardsByEra{}
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := client.GetBlockLatest(c.ctx)
	if err != nil {
		return rewards, CasperDecimals, NewRpcConnectionError(fmt.Sprintf("rpc error: %s", err.Error()))
	}
	nowera := int(res.Block.Header.EraID)
	if era_start > nowera {
		return rewards, CasperDecimals, fmt.Errorf("start era is in the future")
	}
	if era_start == nowera {
		return rewards, CasperDecimals, fmt.Errorf("start era is current")
	}
	if era_end > nowera {
		return rewards, CasperDecimals, fmt.Errorf("end era is in the future")
	}
	if era_end == nowera {
		return rewards, CasperDecimals, fmt.Errorf("end era is current")
	}

	lastBlocks := make([]int64, 0)
	for i := era_start; i <= era_end; i++ {
		_, blockEnd, err := c.GetEraBounds(rpc, i)
		if err != nil {
			return rewards, CasperDecimals, err
		}
		lastBlocks = append(lastBlocks, blockEnd)
	}
	for _, lb := range lastBlocks {
		res, err := client.GetEraInfoByBlockHeight(c.ctx, uint64(lb))
		if err != nil {
			return rewards, CasperDecimals, err
		}
		res_blockinfo, err := client.GetBlockByHeight(c.ctx, uint64(lb))
		if err != nil {
			return rewards, CasperDecimals, err
		}
		for _, reward := range res.EraSummary.StoredValue.EraInfo.SeigniorageAllocations {
			if reward.Delegator == nil {
				continue
			}
			if reward.Delegator.DelegatorPublicKey.String() == address {
				amount := reward.Delegator.Amount.Value()
				rewards = append(rewards, blockchain.RewardsByEra{
					Amount:          amount,
					ValidatorPubkey: reward.Delegator.ValidatorPublicKey.String(),
					Era:             res.EraSummary.EraID,
					Date:            res_blockinfo.Block.Header.Timestamp.ToTime(),
				})
			}
		}
	}
	return rewards, CasperDecimals, nil

}
