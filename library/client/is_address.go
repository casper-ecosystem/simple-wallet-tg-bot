package client

import "github.com/make-software/casper-go-sdk/types/keypair"

// func (c *Client) IsAddress(rpc string, chain string, address string) (bool, error) { //TODO
// 	client := sdk.NewRpcClient(rpc)
// 	_, err := client.GetAccountInfoLatestBlock(address)
// 	log.Println(err)
// 	if err != nil {
// 		if customerr, ok := err.(*sdk.RpcError); ok {
// 			if customerr.Code == -32003 {
// 				return false, nil
// 			}
// 		}
// 		return false, err
// 	}

// 	return true, nil
// }

func (c *Client) IsAddress(address string) bool { //TODO not contain unicode
	_, err := keypair.NewPublicKey(address)
	return err == nil
}
