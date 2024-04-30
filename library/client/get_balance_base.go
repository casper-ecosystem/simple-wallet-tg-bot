package client

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"

	sdk "github.com/make-software/casper-go-sdk/casper"
	sdkrpc "github.com/make-software/casper-go-sdk/rpc"
	"github.com/make-software/casper-go-sdk/types/keypair"
)

func (c *Client) GetBalanceBase(rpc string, address string) (big.Int, int, error) {
	if !c.IsAddress(address) {
		return big.Int{}, 0, NewRpcAdressInvalidError(fmt.Sprintf("address %s is invalid", address))
	}
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	res, err := client.GetBlockLatest(c.ctx)
	if err != nil {
		return big.Int{}, 0, err
	}
	block_hash := res.Block.Hash
	stateRootHash := res.Block.Header.StateRootHash.String()
	pubkey, err := keypair.NewPublicKey(address)
	if err != nil {
		return big.Int{}, 0, errors.New("Bad pubkey")
	}

	res2, err := client.GetAccountInfoByBlochHash(c.ctx, block_hash.String(), pubkey)

	if err != nil {
		var rpcerr *sdkrpc.RpcError
		if errors.As(err, &rpcerr) {
			if rpcerr.Code == -32003 {
				return *big.NewInt(0), CasperDecimals, nil
			}
		}
		return big.Int{}, 0, err
	}
	purse := res2.Account.MainPurse
	balance, err := client.GetAccountBalance(c.ctx, &stateRootHash, purse.String())
	if err != nil {
		return big.Int{}, 0, err
	}
	balancebi := balance.BalanceValue.Value()
	return *balancebi, CasperDecimals, nil

}
