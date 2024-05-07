package client

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Simplewallethq/source-code/library/blockchain"
	sdk "github.com/make-software/casper-go-sdk/casper"
	"github.com/make-software/casper-go-sdk/types/keypair"
)

func (c *Client) GetHistoryTransfers(rpc string, address string, blockStart int, blockEnd int) ([]blockchain.TransferResponse, int, error) { //TODO account hash locally
	result := []blockchain.TransferResponse{}

	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	pubkey, err := keypair.NewPublicKey(address)
	if err != nil {
		return []blockchain.TransferResponse{}, 0, errors.New("Bad pubkey")
	}
	account_hash := pubkey.AccountHash()
	//log.Println(account_hash)
	// account_hash := "account-hash-" + keypair.AccountHash()

	for i := blockStart; i <= blockEnd; i++ {
		blockTransfers, err := client.GetBlockTransfersByHeight(c.ctx, uint64(i))
		if err != nil {
			log.Println("ERROR HERE: ", i)
			return []blockchain.TransferResponse{}, 0, err
		}
		for _, transfer := range blockTransfers.Transfers {
			var frompubkey, topubkey string
			var to string
			if transfer.To != nil {
				to = transfer.To.String()
			} else {
				to = topubkey
			}

			if transfer.From.String() == account_hash.String() || to == account_hash.String() {
				deploy, err := client.GetDeploy(c.ctx, transfer.DeployHash.String())
				if err != nil {
					return []blockchain.TransferResponse{}, 0, err
				}
				frompubkey = deploy.Deploy.Header.Account.String()
				if deploy.Deploy.Session.Transfer != nil {
					target, err := deploy.Deploy.Session.Transfer.Args.Find("target")
					if err != nil {
						topubkey = ""
					} else {
						targetParsed, err := target.Value()
						if err != nil {
							return []blockchain.TransferResponse{}, 0, err
						}

						topubkey = targetParsed.String()
					}

				}

				result = append(result, blockchain.TransferResponse{
					DeployHash: transfer.DeployHash.String(),
					From:       transfer.From.String(),
					FromPubKey: frompubkey,
					To:         to,
					ToPubKey:   topubkey,
					Amount:     transfer.Amount.Value().String(),
					Gas:        strconv.FormatUint(uint64(transfer.Gas), 10),
					Height:     uint64(i),
				})
			}

		}

	}
	return result, CasperDecimals, nil
}
