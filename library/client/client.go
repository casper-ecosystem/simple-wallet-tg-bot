package client

import "context"

const CasperDecimals = 9
const UndelegateDelay = 8

type Client struct {
	ctx context.Context
}

func NewClient() *Client {
	return &Client{
		ctx: context.Background(),
	}
}

func NewClientWithContext(ctx context.Context) *Client {
	return &Client{
		ctx: ctx,
	}

}
