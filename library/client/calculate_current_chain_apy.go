package client

import (
	"errors"
	"math/big"
	"net/http"

	sdk "github.com/make-software/casper-go-sdk/casper"
)

func (c *Client) CalculateCurrentChainAPY(rpc string, chain string) (float64, error) {
	testnetUREF := "uref-5d7b1b23197cda53dec593caf30836a5740afa2279b356fae74bf1bdc2b2e725-007"
	mainnetUREF := "uref-8032100a1dcc56acf84d5fc9c968ce8caa5f2835ed665a2ae2186141e9946214-007"
	var UREF string
	totalSupply := big.NewInt(0)
	totalStake := big.NewInt(0)
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	nowblock, err := client.GetBlockLatest(c.ctx)
	if err != nil {
		return 0.0, err
	}
	switch chain {
	case "testnet":
		UREF = testnetUREF
		//fmt.Println("testnet", testnetUREF)
	case "mainnet":
		UREF = mainnetUREF
		//fmt.Println("mainnet", mainnetUREF)
	default:
		return 0.0, errors.New("unknown chain")
	}
	svres, err := client.QueryGlobalStateByBlockHash(c.ctx,
		nowblock.Block.Hash.String(),
		UREF, []string{})
	if err != nil {
		return 0.0, err
	}
	if svres.StoredValue.CLValue != nil {
		v, err := svres.StoredValue.CLValue.Value()
		if err != nil {
			return 0.0, err
		}
		totalSupply.SetString(v.String(), 10)
	}

	//fmt.Println(totalSupply.String())

	validators, res := client.GetAuctionInfoLatest(c.ctx)
	if res != nil {
		return 0.0, res
	}

	for _, bid := range validators.AuctionState.Bids {
		if bid.Bid.Inactive {
			continue
		}
		var tmp big.Int
		tmp = (*bid.Bid.StakedAmount.Value())
		totalStake.Add(totalStake, &tmp)
		for _, del := range bid.Bid.Delegators {
			var tmp2 big.Int
			tmp2 = (*del.StakedAmount.Value())
			totalStake.Add(totalStake, &tmp2)
		}
	}
	//(0.08 * TOTAL_SUPPLY) / TOTAL_STAKE AMOUNT
	percentage := big.NewFloat(0.08)
	hundred := big.NewFloat(100)

	// Calculate the result
	result := new(big.Float).Mul(percentage, new(big.Float).SetInt(totalSupply))
	result = new(big.Float).Mul(result, hundred)
	result = new(big.Float).Quo(result, new(big.Float).SetInt(totalStake))

	// Convert the result to float64
	resultFloat64, _ := result.Float64()
	//fmt.Println(totalStake)
	//fmt.Println(totalSupply)
	return resultFloat64, nil

}
