package client

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) GetAPRByERA(rpc string, address string, era_start int, era_end int) (float64, error) {
	rewards, _, err := c.GetRewardsByEra(rpc, address, era_start, era_end)
	rewards_sum := big.NewInt(0)
	for _, reward := range rewards {
		rewards_sum.Add(rewards_sum, reward.Amount)
	}
	if err != nil {
		return 0.0, err
	}
	staked_sum := big.Int{}
	for i := era_start; i <= era_end; i++ {
		staked, _, err := c.getBalanceStakedByEra(rpc, address, i)
		if err != nil {
			return 0.0, err
		}
		for _, v := range staked {
			staked_sum.Add(&staked_sum, &v)
		}
	}
	era_count := (era_end - era_start) + 1
	if era_count <= 0 {
		return 0.0, errors.New("era start must be less or equal than era end")
	}
	average_staked_balance := big.Int{}
	average_staked_balance.Div(&staked_sum, big.NewInt(int64(era_count)))
	average_reward := big.Int{}
	average_reward.Div(rewards_sum, big.NewInt(int64(era_count)))

	//find number of eras in year
	_, nowera, err := c.GetCurrentState(rpc)
	if err != nil {
		return 0.0, err
	}
	era_year_ago, err := c.GetEraByTimestamp(rpc, time.Now().AddDate(-1, 0, 0))
	if err != nil {
		return 0.0, err
	}
	number_of_eras_in_year := nowera - era_year_ago
	// (average_reward / average_staked_balance) * number_of_eras_in_year
	apy := big.Float{}
	if average_staked_balance.Cmp(big.NewInt(0)) == 0 {
		return 0.0, nil
	}
	if average_reward.Cmp(big.NewInt(0)) == 0 {
		return 0.0, nil
	}
	apy.Quo(
		big.NewFloat(0).SetInt(&average_reward),
		big.NewFloat(0).SetInt(&average_staked_balance),
	)
	apy.Mul(&apy, big.NewFloat(float64(number_of_eras_in_year)))
	apy.Mul(&apy, big.NewFloat(100))
	apyres, _ := apy.Float64()
	return apyres, nil
}

func (c *Client) getBalanceStakedByEra(rpc string, address string, era int) (map[string]big.Int, int, error) { //TODO: caching
	balance := make(map[string]big.Int)
	if !c.IsAddress(address) {
		return balance, CasperDecimals, NewRpcAdressInvalidError(fmt.Sprintf("address %s is invalid", address))
	}
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	_, era_end, err := c.GetEraBounds(rpc, era)
	if err != nil {
		return balance, CasperDecimals, err
	}
	res, err := client.GetAuctionInfoByHeight(c.ctx, uint64(era_end))
	if err != nil {
		return balance, CasperDecimals, err
	}

	for _, bid := range res.AuctionState.Bids {
		staked := big.Int{}
		for _, delegator := range bid.Bid.Delegators {
			if delegator.PublicKey.String() == address {
				amount := (delegator.StakedAmount.Value())
				staked.Add(&staked, amount)
			}
		}
		//check if staked is greater than 0
		if staked.Cmp(big.NewInt(0)) == 1 {
			balance[bid.PublicKey.String()] = staked
		}

	}
	return balance, CasperDecimals, err
}
