package casper

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/Simplewallethq/library/blockchain"
	"github.com/gin-gonic/gin"
	"github.com/make-software/casper-go-sdk/types"
	"github.com/sirupsen/logrus"
)

type Casper struct {
	client blockchain.Client
	logger *logrus.Logger
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func New(client blockchain.Client, logger *logrus.Logger) *Casper {
	return &Casper{
		client: client,
		logger: logger,
	}
}

func IpToUrl(input string) (string, error) {
	//sanitizing
	ip := net.ParseIP(input)
	if ip == nil {
		return "", errors.New("invalid rpc_node IP address")
	}
	return "http://" + ip.String() + ":7777/rpc", nil

}

// @Summary Get the current state of the Casper network
// @Description Retrieves the current block height and era of the Casper network by querying the given rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(89.58.55.157)
// @Success 200 {object} casper.GetStateHandler.Response "A JSON object containing the blockHeight and currentEra of the Casper network"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node parameter is missing or an invalid IP address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/state [get]
// @Security ApiKeyAuth
func (cn *Casper) GetStateHandler(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}

	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}

	blockHeight, currentEra, err := cn.client.GetCurrentState(rpcurl)
	if err != nil {
		c.JSON(500, ErrorResponse{"error querying casper network"})
		cn.logger.Error(err)
		return
	}

	type Response struct {
		BlockHeight int `json:"block_height"`
		CurrentEra  int `json:"era_id"`
	}
	c.JSON(200, Response{
		BlockHeight: blockHeight,
		CurrentEra:  currentEra,
	})
}

// @Summary Check if a public key is a valid Casper address
// @Description Determines if the provided public key is a valid Casper address by querying the given rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(89.58.55.157)
// @Param public_key query string true "The public key of the Casper address to be checked for validity" default(011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635)
// @Success 200 {object} casper.IsAddressHandler.Response "A JSON object containing the validity of the provided public key"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node or pubkey parameter is missing or an invalid IP address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/is_address [get]
// @Security ApiKeyAuth
func (cn *Casper) IsAddressHandler(c *gin.Context) {

	pubkey := c.Query("public_key")
	if pubkey == "" {
		c.JSON(400, ErrorResponse{"public_key is required"})
		return
	}
	valid := cn.client.IsAddress(pubkey)

	type Response struct {
		Valid bool `json:"valid"`
	}
	c.JSON(200, Response{
		Valid: valid,
	})

}

// @Summary Get liquid account balance for a Casper address
// @Description Retrieves the account balance for a given Casper address by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The public key of the Casper address to retrieve the balance for" default(011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635)
// @Success 200 {object} casper.GetAccountBalanceHandler.Response "A JSON object containing the balance and decimals values of the provided Casper address"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node or pubkey parameter is missing or an invalid IP address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_balance_main [get]
// @Security ApiKeyAuth
func (cn *Casper) GetAccountBalanceHandler(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	pubkey := c.Query("address")
	if pubkey == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}
	balance, dec, err := cn.client.GetBalanceBase(rpcurl, pubkey)
	if err != nil {
		c.JSON(500, ErrorResponse{"error querying casper network"})
		cn.logger.Error(err)
		return
	}
	type Response struct {
		Balance  string `json:"balance"`
		Decimals int    `json:"decimals"`
	}
	c.JSON(200, Response{
		Balance:  balance.String(),
		Decimals: dec,
	})
}

// @Summary Get delegated account balance for a Casper address
// @Description Retrieves the delegated account balance for a given Casper address by querying the specified rpc_node and returns the balance associated with each validator
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The public key of the Casper address to retrieve the delegated balance for" default(011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635)
// @Success 200 {object} casper.GetDelegatedBalanceHandler.Response "A JSON object containing the delegated balance and decimals values of the provided Casper address"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node or pubkey parameter is missing or an invalid IP address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_balance_delegated [get]
// @Security ApiKeyAuth
func (cn *Casper) GetDelegatedBalanceHandler(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	pubkey := c.Query("address")
	if pubkey == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}
	balance, dec, err := cn.client.GetBalanceStaked(rpcurl, pubkey)
	if err != nil {
		c.JSON(500, ErrorResponse{"error querying casper network"})
		cn.logger.Error(err)
		return
	}
	type Response struct {
		Data []struct {
			Validator string `json:"validator"`
			Balance   string `json:"amount"`
		} `json:"data"`
		Decimals int `json:"decimals"`
	}
	var resp Response
	for k, v := range balance {
		resp.Data = append(resp.Data, struct {
			Validator string `json:"validator"`
			Balance   string `json:"amount"`
		}{Validator: k, Balance: v.String()})
	}
	resp.Decimals = dec
	c.JSON(200, resp)
}

type TransferResponse struct {
	DeployHash string `json:"deploy_hash"`
	From       string `json:"from"`
	FromPubKey string `json:"from_pubkey"`
	To         string `json:"to"`
	ToPubKey   string `json:"to_pubkey"`
	Amount     string `json:"amount"`
	Gas        string `json:"gas"`
	Height     uint64 `json:"height"`
	Memo       uint64 `json:"memo"`
}

// @Summary Get the transaction history for a Casper address
// @Description Retrieves the transaction history for a given Casper address by querying the specified rpc_node and returns the transactions between the specified block range
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The public key of the Casper address to retrieve the transaction history for" default(020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292)
// @Param start_block query string true "The starting block height for the transaction history" default(1638650)
// @Param end_block query string true "The ending block height for the transaction history" default(1638660)
// @Success 200 {object} casper.GetHistoryTransfers.Response "A JSON object containing the transaction history and decimals values of the provided Casper address"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, pubkey, start_block, or end_block parameter is missing, or an invalid IP address, public key, or block height is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_history_transfers [get]
// @Security ApiKeyAuth
func (cn *Casper) GetHistoryTransfers(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	pubkey := c.Query("address")
	if pubkey == "" {
		c.JSON(400, ErrorResponse{"pubkey is required"})
		return
	}
	start_block := c.Query("start_block")
	if start_block == "" {
		c.JSON(400, ErrorResponse{"start_block is required"})
		return
	}
	start_block_int, err := strconv.Atoi(start_block)
	if err != nil {
		c.JSON(400, ErrorResponse{"start_block is invalid"})
		return
	}
	end_block := c.Query("end_block")
	if end_block == "" {
		c.JSON(400, ErrorResponse{"end_block is required"})
		return
	}
	end_block_int, err := strconv.Atoi(end_block)
	if err != nil {
		c.JSON(400, ErrorResponse{"end_block is invalid"})
		return
	}
	transfers, dec, err := cn.client.GetHistoryTransfers(rpcurl, pubkey, start_block_int, end_block_int)
	if err != nil {
		c.JSON(500, ErrorResponse{"error querying casper network"})
		cn.logger.Error(err)
		return
	}

	type Response struct {
		Transfers []TransferResponse `json:"data"`
		Decimals  int                `json:"decimals"`
	}
	var resp Response
	tr := make([]TransferResponse, 0, len(transfers))
	for _, v := range transfers {
		tr = append(tr, TransferResponse{
			DeployHash: v.DeployHash,
			From:       v.From,
			FromPubKey: v.FromPubKey,
			To:         v.To,
			ToPubKey:   v.ToPubKey,
			Amount:     v.Amount,
			Gas:        v.Gas,
			Height:     v.Height,
		})
	}

	resp.Transfers = tr
	resp.Decimals = dec
	c.JSON(200, resp)
}

type RewardsByBlockPayload struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator"`
	Block           uint64 `json:"block"`
}

type ResponseRewardsByBlock struct {
	Rewards  []RewardsByBlockPayload `json:"rewards"`
	Decimals int                     `json:"decimals"`
}

// @Summary Get the rewards earned by a Casper address
// @Description Retrieves the rewards earned by a given Casper address by querying the specified rpc_node and returns the rewards between the specified block range
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The public key of the Casper address to retrieve the rewards for" default(011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635)
// @Param start_block query string true "The starting block height for the rewards" default(1639836)
// @Param end_block query string true "The ending block height for the rewards" default(1639836)
// @Success 200 {object} casper.ResponseRewardsByBlock "A JSON object containing the rewards earned by the provided Casper address"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, pubkey, start_block, or end_block parameter is missing, or an invalid IP address, public key, or block height is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_rewards_by_blocks [get]
// @Security ApiKeyAuth
func (cn *Casper) GetRewardsByBlock(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	pubkey := c.Query("address")
	if pubkey == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}
	start_block := c.Query("start_block")
	if start_block == "" {
		c.JSON(400, ErrorResponse{"start_block is required"})
		return
	}
	start_block_int, err := strconv.Atoi(start_block)
	if err != nil {
		c.JSON(400, ErrorResponse{"start_block is invalid"})
		return
	}
	end_block := c.Query("end_block")
	if end_block == "" {
		c.JSON(400, ErrorResponse{"end_block is required"})
		return
	}
	end_block_int, err := strconv.Atoi(end_block)
	if err != nil {
		c.JSON(400, ErrorResponse{"end_block is invalid"})
		return
	}
	rewards, dec, err := cn.client.GetRewardsByBlock(rpcurl, pubkey, int64(start_block_int), int64(end_block_int))
	if err != nil {
		c.JSON(500, ErrorResponse{"error querying casper network"})
		cn.logger.Error(err)
		return
	}
	var resp ResponseRewardsByBlock
	for _, reward := range rewards {
		resp.Rewards = append(resp.Rewards, RewardsByBlockPayload{
			Amount:          reward.Amount.String(),
			ValidatorPubkey: reward.ValidatorPubkey,
			Block:           reward.Block,
		})
	}
	resp.Decimals = dec
	c.JSON(200, resp)

}

type RewardsByEraPayload struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator"`
	Era             uint32 `json:"era"`
	Timestamp       string `json:"timestamp"`
}

type ResponseRewardsByEra struct {
	Rewards  []RewardsByEraPayload `json:"rewards"`
	Decimals int                   `json:"decimals"`
}

// @Summary Get the rewards earned by a Casper address
// @Description Retrieves the rewards earned by a given Casper address by querying the specified rpc_node and returns the rewards between the specified block range
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The public key of the Casper address to retrieve the rewards for" default(020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292)
// @Param start_era query string true "The starting era for the rewards" default(8902)
// @Param end_era query string true "The ending era for the rewards" default(8903)
// @Success 200 {object} casper.ResponseRewardsByEra "A JSON object containing the rewards earned by the provided Casper address"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, pubkey, start_block, or end_block parameter is missing, or an invalid IP address, public key, or block height is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_rewards_by_era [get]
// @Security ApiKeyAuth
func (cn *Casper) GetRewardsByEra(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	pubkey := c.Query("address")
	if pubkey == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}
	start_era := c.Query("start_era")
	if start_era == "" {
		c.JSON(400, ErrorResponse{"start_era is required"})
		return
	}
	start_era_int, err := strconv.Atoi(start_era)
	if err != nil {
		c.JSON(400, ErrorResponse{"start_era is invalid"})
		return
	}
	end_era := c.Query("end_era")
	if end_era == "" {
		c.JSON(400, ErrorResponse{"end_era is required"})
		return
	}
	end_era_int, err := strconv.Atoi(end_era)
	if err != nil {
		c.JSON(400, ErrorResponse{"end_era is invalid"})
		return
	}
	rewards, dec, err := cn.client.GetRewardsByEra(rpcurl, pubkey, start_era_int, end_era_int)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}
	var resp ResponseRewardsByEra
	for _, reward := range rewards {
		resp.Rewards = append(resp.Rewards, RewardsByEraPayload{
			Amount:          reward.Amount.String(),
			ValidatorPubkey: reward.ValidatorPubkey,
			Era:             reward.Era,
			Timestamp:       reward.Date.UTC().Format(time.RFC1123),
		})
	}
	resp.Decimals = dec
	c.JSON(200, resp)

}

// @Summary Get the timestamp of a Casper block
// @Description Retrieves the timestamp of a given Casper block by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param block_height query string true "The block height to retrieve the timestamp for" default(1639836)
// @Success 200 {object} casper.GetTimestampByBlock.Response "A JSON object containing the timestamp of the provided Casper block"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, block_height parameter is missing, or an invalid IP address or block height is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_timestamp_by_block [get]
// @Security ApiKeyAuth
func (cn *Casper) GetTimestampByBlock(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	type Response struct {
		BlockHeight int    `json:"block_height"`
		Timestamp   string `json:"timestamp"`
		Estimated   bool   `json:"estimated"`
	}
	var resp Response
	blockHeight := c.Query("block_height")
	if blockHeight == "" {
		c.JSON(400, ErrorResponse{"block_height is required"})
		return
	}
	blockHeightInt, err := strconv.Atoi(blockHeight)
	if err != nil {
		c.JSON(400, ErrorResponse{"block_height is invalid"})
		return
	}
	timestamp, estimated, err := cn.client.GetTimestampByBlock(rpcurl, blockHeightInt)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}
	resp.BlockHeight = blockHeightInt
	resp.Timestamp = timestamp.UTC().Format(time.RFC1123)
	resp.Estimated = estimated
	c.JSON(200, resp)
}

// @Summary Get the timestamp of a Casper era
// @Description Retrieves the timestamp of a given Casper era by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param era_id query string true "The era ID to retrieve the timestamp for" default(8822)
// @Success 200 {object} casper.GetTimestampByEra.Response "A JSON object containing the timestamp of the provided Casper era"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, era_id parameter is missing, or an invalid IP address or era ID is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_timestamp_by_era [get]
// @Security ApiKeyAuth
func (cn *Casper) GetTimestampByEra(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	type Response struct {
		EraId     int    `json:"era_id"`
		Timestamp string `json:"timestamp_era_start"`
		Estimated bool   `json:"estimated"`
	}
	var resp Response
	eraId := c.Query("era_id")
	if eraId == "" {
		c.JSON(400, ErrorResponse{"era_id is required"})
		return
	}
	eraIdInt, err := strconv.Atoi(eraId)
	if err != nil {
		c.JSON(400, ErrorResponse{"era_id is invalid"})
		return
	}
	timestamp, estimated, err := cn.client.GetTimestampByEra(rpcurl, eraIdInt)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}
	resp.EraId = eraIdInt
	resp.Timestamp = timestamp.UTC().Format(time.RFC1123)
	resp.Estimated = estimated
	c.JSON(200, resp)
}

// @Summary Get the timestamp of a Casper block by timestamp
// @Description Retrieves the timestamp of a given Casper block by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param timestamp query string true "The timestamp to retrieve the block height for" default(2023-04-11T11:18:04+00:00)
// @Success 200 {object} casper.GetBlockByTimestamp.Response "A JSON object containing the block height of the provided Casper timestamp"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, timestamp parameter is missing, or an invalid IP address or timestamp is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_block_by_timestamp [get]
// @Security ApiKeyAuth
func (cn *Casper) GetBlockByTimestamp(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	type Response struct {
		BlockHeight int    `json:"block_height"`
		Timestamp   string `json:"timestamp"`
	}
	var resp Response
	timestamp := c.Query("timestamp")
	if timestamp == "" {
		c.JSON(400, ErrorResponse{"timestamp is required"})
		return
	}
	//parse timestamp to time.Time
	timestampTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		c.JSON(400, ErrorResponse{"timestamp is invalid"})
		cn.logger.Error(err)
		return
	}
	//log.Print(timestampTime.Unix())
	blockHeight, err := cn.client.GetBlockByTimestamp(rpcurl, timestampTime)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}
	resp.BlockHeight = blockHeight
	resp.Timestamp = timestampTime.UTC().Format(time.RFC1123)
	c.JSON(200, resp)
}

// @Summary Get the timestamp of a Casper era by timestamp
// @Description Retrieves the timestamp of a given Casper era by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param timestamp query string true "The timestamp to retrieve the block height for" default(2023-04-11T11:18:04+00:00)
// @Success 200 {object} casper.GetEraByTimestamp.Response "A JSON object containing the era ID of the provided Casper timestamp"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, timestamp parameter is missing, or an invalid IP address or timestamp is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_era_by_timestamp [get]
// @Security ApiKeyAuth
func (cn *Casper) GetEraByTimestamp(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	type Response struct {
		EraId     int    `json:"era_id"`
		Timestamp string `json:"timestamp"`
	}
	var resp Response
	timestamp := c.Query("timestamp")
	if timestamp == "" {
		c.JSON(400, ErrorResponse{"timestamp is required"})
		return
	}
	//parse timestamp to time.Time
	timestampTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		c.JSON(400, ErrorResponse{"timestamp is invalid"})
		cn.logger.Error(err)
		return
	}
	era, err := cn.client.GetEraByTimestamp(rpcurl, timestampTime)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}
	log.Println(era)
	resp.EraId = era
	resp.Timestamp = timestampTime.UTC().Format(time.RFC1123)
	c.JSON(200, resp)
}

type PayloadBalanceBeingUndelegated struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator"`
	Era             int    `json:"era_undelegation_finished"`
}

// @Summary Get the balance being undelegated for an address
// @Description Retrieves the balance being undelegated for a given address by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The address for which to retrieve the unstaking balance" default(020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292)
// @Success 200 {array} casper.GetBalanceBeingUndelegated.Response{Data=casper.PayloadBalanceBeingUndelegated} "A JSON array containing the unstaking balances for the provided address"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, address parameter is missing, or an invalid IP address or address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_balance_being_undelegated [get]
// @Security ApiKeyAuth
func (cn *Casper) GetBalanceBeingUndelegated(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	address := c.Query("address")
	if address == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}

	type Response struct {
		Data     []PayloadBalanceBeingUndelegated `json:"data"`
		Decimals int                              `json:"decimals"`
	}
	var resp Response

	unstakingBalances, dec, err := cn.client.GetBalanceBeingUnstaked(rpcurl, address)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}

	for _, balance := range unstakingBalances {
		resp.Data = append(resp.Data, PayloadBalanceBeingUndelegated{
			Amount:          balance.Amount,
			ValidatorPubkey: balance.ValidatorPubkey,
			Era:             balance.Era,
		})
	}
	resp.Decimals = dec

	c.JSON(200, resp)
}

type PayloadHistoryUndelegate struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator"`
	Era             int    `json:"era"`
	Finished        bool   `json:"is_finished"`
	Height          uint64 `json:"height"`
}

// @Summary Get history of undelegate deploys for an address
// @Description Retrieves the history of undelegate for a given address by querying the specified rpc_node within a range of block heights
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The address for which to retrieve the history of finished unstake" default(020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292)
// @Param block_start query int true "The starting block height for the range" default(1660930)
// @Param block_end query int true "The ending block height for the range" default(1660950)
// @Success 200 {array} casper.GetHistoryUndelegate.Response "A JSON array containing the history of finished unstakes for the provided address within the specified range"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, address, block_start, or block_end parameter is missing, or an invalid IP address or address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_history_undelegate [get]
// @Security ApiKeyAuth
func (cn *Casper) GetHistoryUndelegate(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	address := c.Query("address")
	if address == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}
	blockStartStr := c.Query("block_start")
	if blockStartStr == "" {
		c.JSON(400, ErrorResponse{"block_start is required"})
		return
	}
	blockStart, err := strconv.Atoi(blockStartStr)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid block_start value"})
		return
	}
	blockEndStr := c.Query("block_end")
	if blockEndStr == "" {
		c.JSON(400, ErrorResponse{"block_end is required"})
		return
	}
	blockEnd, err := strconv.Atoi(blockEndStr)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid block_end value"})
		return
	}

	type Response struct {
		Data     []PayloadHistoryUndelegate `json:"data"`
		Decimals int                        `json:"decimals"`
	}
	var resp Response

	historyUnstakes, dec, err := cn.client.GetHistoryUndelegate(rpcurl, address, blockStart, blockEnd)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}

	for _, unstake := range historyUnstakes {
		resp.Data = append(resp.Data, PayloadHistoryUndelegate{
			Amount:          unstake.Amount,
			ValidatorPubkey: unstake.ValidatorPubkey,
			Era:             unstake.Era,
			Finished:        unstake.Finished,
			Height:          unstake.Height,
		})
	}
	resp.Decimals = dec

	c.JSON(200, resp)
}

// @Summary Get the price of the main coin (CSPR)
// @Description Retrieves the current price of the CSPR coin from a supported exchange
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Success 200 {object} float64 "A JSON object containing the price of the main coin (CSPR)"
// @Failure 500 {object} ErrorResponse "Returned when there is an error retrieving the price from the exchange or processing the response"
// @Router /api/v1/{cspr_chain}/get_price_main_coin [get]
// @Security ApiKeyAuth
func (cn *Casper) GetPriceMainCoin(c *gin.Context) {
	price, err := cn.client.GetPriceMainCoin()
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}
	c.JSON(200, gin.H{
		"price": price,
	})
}

type PayloadHistoryDelegate struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator"`
	Era             int    `json:"era"`
	Finished        bool   `json:"is_finished"`
	Height          uint64 `json:"height"`
}

// @Summary Get history of delegates for an address
// @Description Retrieves the history of delegates for a given address by querying the specified rpc_node within a range of block heights
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The address for which to retrieve the history of finished stake" default(020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292)
// @Param block_start query int true "The starting block height for the range" default(1660930)
// @Param block_end query int true "The ending block height for the range" default(1660950)
// @Success 200 {array} casper.GetHistoryDelegate.Response "A JSON array containing the history of finished stakes for the provided address within the specified range"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, address, block_start, or block_end parameter is missing, or an invalid IP address or address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_history_delegate [get]
// @Security ApiKeyAuth
func (cn *Casper) GetHistoryDelegate(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	address := c.Query("address")
	if address == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}
	blockStartStr := c.Query("block_start")
	if blockStartStr == "" {
		c.JSON(400, ErrorResponse{"block_start is required"})
		return
	}
	blockStart, err := strconv.Atoi(blockStartStr)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid block_start value"})
		return
	}
	blockEndStr := c.Query("block_end")
	if blockEndStr == "" {
		c.JSON(400, ErrorResponse{"block_end is required"})
		return
	}
	blockEnd, err := strconv.Atoi(blockEndStr)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid block_end value"})
		return
	}

	type Response struct {
		Data     []PayloadHistoryDelegate `json:"data"`
		Decimals int                      `json:"decimals"`
	}
	var resp Response

	historyStakes, dec, err := cn.client.GetHistoryDelegate(rpcurl, address, blockStart, blockEnd)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}
	resp.Data = []PayloadHistoryDelegate{}

	for _, stake := range historyStakes {
		resp.Data = append(resp.Data, PayloadHistoryDelegate{
			Amount:          stake.Amount,
			ValidatorPubkey: stake.ValidatorPubkey,
			Era:             stake.Era,
			Finished:        stake.Finished,
			Height:          stake.Height,
		})
	}
	resp.Decimals = dec

	c.JSON(200, resp)
}

// @Summary Get APR for a delegator by era range
// @Description Retrieves the Annual Percentage Rate (APR) for a delegator address by querying the specified rpc_node within a range of eras
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The address of the delegator for which to retrieve the APR" default(020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292)
// @Param era_start query int true "The starting era for the range" default(8924)
// @Param era_end query int true "The ending era for the range" default(8927)
// @Success 200 {object} casper.GetAPRByEra.Response "A JSON object containing the APY for the provided delegator address within the specified range"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node, address, era_start, or era_end parameter is missing, or an invalid IP address or address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_apr_by_era [get]
// @Security ApiKeyAuth
func (cn *Casper) GetAPRByEra(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	address := c.Query("address")
	if address == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}
	eraStartStr := c.Query("era_start")
	if eraStartStr == "" {
		c.JSON(400, ErrorResponse{"era_start is required"})
		return
	}
	eraStart, err := strconv.Atoi(eraStartStr)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid era_start value"})
		return
	}
	eraEndStr := c.Query("era_end")
	if eraEndStr == "" {
		c.JSON(400, ErrorResponse{"era_end is required"})
		return
	}
	eraEnd, err := strconv.Atoi(eraEndStr)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid era_end value"})
		return
	}
	type Response struct {
		APR float64 `json:"apr"`
	}

	apr, err := cn.client.GetAPRByERA(rpcurl, address, eraStart, eraEnd)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}

	resp := Response{
		APR: apr,
	}

	c.JSON(200, resp)
}

type PayloadBalanceBeingdelegated struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator"`
	Era             int    `json:"era_delegation_finished"`
}

// @Summary Get balance being delegated for a delegator
// @Description Retrieves the balance being delegated for a delegator address by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param address query string true "The address of the delegator for which to retrieve the balance being delegated" default(020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292)
// @Success 200 {array} casper.GetBalanceBeingdelegated.Response "An array of JSON objects containing the balance being delegated for the provided delegator address"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node or address parameter is missing, or an invalid IP address or address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_balance_being_delegated [get]
// @Security ApiKeyAuth
func (cn *Casper) GetBalanceBeingdelegated(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	address := c.Query("address")
	if address == "" {
		c.JSON(400, ErrorResponse{"address is required"})
		return
	}
	beingdelegated, dec, err := cn.client.GetBalanceBeingStaked(rpcurl, address)

	type Response struct {
		Data []PayloadBalanceBeingdelegated `json:"data"`
		Dec  int                            `json:"decimals"`
	}
	var resp Response
	for _, stake := range beingdelegated {
		resp.Data = append(resp.Data, PayloadBalanceBeingdelegated{
			Amount:          stake.Amount,
			ValidatorPubkey: stake.ValidatorPubkey,
			Era:             stake.Era,
		})
	}
	resp.Dec = dec

	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}

	c.JSON(200, resp)
}

// @Summary Calculate current chain APY
// @Description Calculates the current Annual Percentage Yield (APY) for the specified chain by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain   path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Success 200 {object} casper.CalculateCurrentChainAPY.Response "A JSON object containing the current APY for the specified chain"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node or chain parameter is missing, or an invalid IP address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/calculate_current_chain_apy [get]
// @Security ApiKeyAuth
func (cn *Casper) CalculateCurrentChainAPY(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	chain_context := c.Value("chain")
	var chain string
	switch chain_context {
	case "casper":
		chain = "mainnet"
	case "casper-test":
		chain = "testnet"
	}
	type Response struct {
		APY float64 `json:"apy"`
	}

	apy, err := cn.client.CalculateCurrentChainAPY(rpcurl, chain)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}

	resp := Response{
		APY: apy,
	}

	c.JSON(200, resp)
}

type Validator struct {
	Address string  `json:"address"`
	Fee     float32 `json:"fee"`
	Active  bool    `json:"active"`
}
type ValidatorsResponse struct {
	Validators []Validator `json:"validators"`
	Decimals   int         `json:"decimals"`
}

// @Summary Get validators list
// @Description Retrieves the list of validators for the specified chain by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Produce json
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Success 200 {object} casper.ValidatorsResponse "A JSON object containing the validators and the decimal for the specified chain"
// @Failure 400 {object} casper.ErrorResponse "Returned when the rpc_node or chain parameter is missing, or an invalid IP address is provided"
// @Failure 500 {object} casper.ErrorResponse "Returned when there is an error querying the Casper network or processing the response"
// @Router /api/v1/{cspr_chain}/get_validators [get]
// @Security ApiKeyAuth
func (cn *Casper) GetValidators(c *gin.Context) {
	rpcNode := c.Query("rpc_node")
	if rpcNode == "" {
		c.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcurl, err := IpToUrl(rpcNode)
	if err != nil {
		c.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	validatorsResp, err := cn.client.GetValidators(rpcurl)
	if err != nil {
		c.JSON(500, ErrorResponse{err.Error()})
		cn.logger.Error(err)
		return
	}
	resp := ValidatorsResponse{Decimals: validatorsResp.Decimals}
	for _, val := range validatorsResp.Validators {
		resp.Validators = append(resp.Validators, Validator{Address: val.Address, Fee: val.Fee, Active: val.Active})
	}
	c.JSON(200, resp)
}

type DelegateData struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator_pubkey"`
	Era             int    `json:"block_height"`
	Finished        bool   `json:"is_finished"`
}

type UndelegateData struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator_pubkey"`
	Era             int    `json:"era"`
	Finished        bool   `json:"is_finished"`
}

type RewardsData struct {
	Delagator string `json:"delagator"`
	Validator string `json:"validator"`
	Amount    string `json:"amount"`
}

type BlockEventsResponse struct {
	Transfers   []TransferResponse `json:"transfers"`
	Delegates   []DelegateData     `json:"delegates"`
	Undelegates []UndelegateData   `json:"undelegates"`
	Rewards     []RewardsData      `json:"rewards"`
	EraID       uint32             `json:"EraID"`
	Decimals    int                `json:"decimals"`
	Date        string             `json:"timestamp"`
}

// @Summary Get all events from a specific block
// @Description Retrieves all events (transfers, delegates, undelegates) from a given block by querying the specified rpc_node
// @Tags Casper
// @Param cspr_chain path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Param height query uint64 true "The block height to retrieve events from" default(1808739)
// @Param transfers query bool true "Include transfer events" default(true)
// @Param delegates query bool true "Include delegate events" default(true)
// @Param undelegates query bool true "Include undelegate events" default(true)
// @Param rewards query bool true "Include rewards events" default(true)
// @Produce json
// @Success 200 {object} BlockEventsResponse "A JSON object containing all events from the provided block"
// @Failure 400 {object} ErrorResponse "Returned when the rpc_node, height, transfers, delegates, or undelegates parameter is missing, or an invalid IP address or block height is provided"
// @Failure 500 {object} ErrorResponse "Returned when there is an error querying the network or processing the response"
// @Router /api/v1/{cspr_chain}/get_block_events [get]
// @Security ApiKeyAuth
func (cn *Casper) GetBlockEvents(ctx *gin.Context) {
	rpc := ctx.Query("rpc_node")
	if rpc == "" {
		ctx.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcURL, err := IpToUrl(rpc)
	if err != nil {
		ctx.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}
	heightStr := ctx.Query("height")
	if heightStr == "" {
		ctx.JSON(400, ErrorResponse{"height is required"})
		return
	}
	height, err := strconv.ParseUint(heightStr, 10, 64)
	if err != nil {
		ctx.JSON(400, ErrorResponse{"height is invalid"})
		return
	}
	transfers, err := strconv.ParseBool(ctx.DefaultQuery("transfers", "true"))
	if err != nil {
		ctx.JSON(400, ErrorResponse{"transfers is invalid"})
		return
	}
	delegates, err := strconv.ParseBool(ctx.DefaultQuery("delegates", "true"))
	if err != nil {
		ctx.JSON(400, ErrorResponse{"delegates is invalid"})
		return
	}
	undelegates, err := strconv.ParseBool(ctx.DefaultQuery("undelegates", "true"))
	if err != nil {
		ctx.JSON(400, ErrorResponse{"undelegates is invalid"})
		return
	}

	rewards, err := strconv.ParseBool(ctx.DefaultQuery("rewards", "true"))
	if err != nil {
		ctx.JSON(400, ErrorResponse{"rewards is invalid"})
		return
	}
	blockEvents, err := cn.client.GetBlockAllEvents(rpcURL, height, transfers, delegates, undelegates, rewards)
	if err != nil {
		ctx.JSON(500, ErrorResponse{"error querying network"})
		return
	}

	resp := BlockEventsResponse{}
	for _, tr := range blockEvents.Transfers {
		resp.Transfers = append(resp.Transfers, TransferResponse{
			DeployHash: tr.DeployHash,
			From:       tr.From,
			FromPubKey: tr.FromPubKey,
			To:         tr.To,
			ToPubKey:   tr.ToPubKey,
			Amount:     tr.Amount,
			Gas:        tr.Gas,
			Memo:       tr.Memo,
		})
	}
	for _, del := range blockEvents.Delegates {
		resp.Delegates = append(resp.Delegates, DelegateData{
			Address:         del.Address,
			Amount:          del.Amount,
			ValidatorPubkey: del.ValidatorPubkey,
			Era:             del.Era,
			Finished:        del.Finished,
		})
	}
	for _, undel := range blockEvents.Undelegates {
		resp.Undelegates = append(resp.Undelegates, UndelegateData{
			Address:         undel.Address,
			Amount:          undel.Amount,
			ValidatorPubkey: undel.ValidatorPubkey,
			Era:             undel.Era,
			Finished:        undel.Finished,
		})
	}

	for _, rew := range blockEvents.Rewards {
		resp.Rewards = append(resp.Rewards, RewardsData{
			Delagator: rew.Delagator,
			Validator: rew.Validator,
			Amount:    rew.Amount.String(),
		})
	}
	resp.Decimals = blockEvents.Decimals
	resp.EraID = blockEvents.Era
	resp.Date = blockEvents.Date.UTC().Format(time.RFC1123)
	ctx.JSON(200, resp)
}

// @Summary Put Deploy to the blockchain
// @Description Put Deploy to the blockchain
// @Tags Casper
// @Param cspr_chain path string true "cspr/cspr-testnet" default(cspr-testnet)
// @Param rpc_node query string true "The IP address of the Casper RPC node" default(65.21.238.180)
// @Produce json
// @Success 200 {object} BlockEventsResponse "A JSON object containing all events from the provided block"
// @Failure 400 {object} ErrorResponse "Returned when the rpc_node, height, transfers, delegates, or undelegates parameter is missing, or an invalid IP address or block height is provided"
// @Failure 500 {object} ErrorResponse "Returned when there is an error querying the network or processing the response"
// @Router /api/v1/{cspr_chain}/put_deploy [get]
// @Security ApiKeyAuth
func (cn *Casper) PutDeploy(ctx *gin.Context) {
	rpc := ctx.Query("rpc_node")
	if rpc == "" {
		ctx.JSON(400, ErrorResponse{"rpc_node is required"})
		return
	}
	rpcURL, err := IpToUrl(rpc)
	if err != nil {
		ctx.JSON(400, ErrorResponse{"Invalid rpc_node IP address"})
		return
	}

	var newDeploy types.Deploy
	jsonDataBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, ErrorResponse{"Invalid body"})
		return
	}
	err = json.Unmarshal(jsonDataBytes, &newDeploy)
	if err != nil {
		ctx.JSON(400, ErrorResponse{"Invalid body"})
		return
	}
	deployHash, err := cn.client.PutDeploy(rpcURL, newDeploy)

	if err != nil {
		cn.logger.Errorf("error put deploy: %s", err.Error())
		ctx.JSON(500, ErrorResponse{"error querying network"})
		return
	}
	ctx.JSON(200, deployHash)
}
