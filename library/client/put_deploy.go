package client

import (
	"net/http"

	sdk "github.com/make-software/casper-go-sdk/casper"
	"github.com/make-software/casper-go-sdk/types"
)

func (c *Client) PutDeploy(rpc string, deploy types.Deploy) (string, error) {
	client := sdk.NewRPCClient(sdk.NewRPCHandler(rpc, http.DefaultClient))
	result, err := client.PutDeploy(c.ctx, deploy)
	if err != nil {
		return "", err
	}
	return result.DeployHash.String(), nil
}
