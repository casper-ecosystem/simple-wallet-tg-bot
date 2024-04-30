package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/make-software/casper-go-sdk/types"
)

type Client struct {
	Host    string
	hclient *http.Client
}

func NewClient(host string) *Client {
	return &Client{
		Host:    host,
		hclient: &http.Client{},
	}
}

func NewClientWithToken(host string, token string) *Client {
	c := http.Client{Transport: &transport{underlyingTransport: http.DefaultTransport, token: token}}
	return &Client{
		Host:    host,
		hclient: &c,
	}
}

type transport struct {
	underlyingTransport http.RoundTripper
	token               string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Token", t.token)
	return t.underlyingTransport.RoundTrip(req)
}

//IsAdressRequest
// 'http://65.108.2.174/rest/api/v1/cspr-testnet/is_address?rpc_node=89.58.55.157&pubkey=011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635'
//return json {
//   "valid": true
// }

func (c *Client) IsAddress(rpcNode string, pubkey string) (bool, error) {
	req := c.Host + "/is_address?rpc_node=" + rpcNode + "&public_key=" + pubkey
	resp, err := c.hclient.Get(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	type Response struct {
		Valid bool   `json:"valid"`
		Error string `json:"error"`
	}
	var response Response
	//read body to bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return false, err
	}
	if response.Error != "" {
		return false, errors.New("rest error: " + response.Error)
	}
	if response.Valid {
		return true, nil
	}
	return false, nil
}

func (c *Client) GetBalance(rpcNode string, pubkey string) (float64, error) {
	req := c.Host + "/get_balance_main?rpc_node=" + rpcNode + "&address=" + pubkey
	resp, err := c.hclient.Get(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()
	type Response struct {
		Balance  string `json:"balance"`
		Decimals int    `json:"decimals"`
		Error    string `json:"error"`
	}
	var response Response
	//read body to bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return 0.0, err
	}
	if response.Error != "" {
		return 0.0, errors.New("rest error: " + response.Error)
	}
	bigBalance, success := new(big.Int).SetString(response.Balance, 10)
	if !success {
		return 0.0, errors.New("invalid balance format")
	}
	// Convert balance to big.Float
	floatBalance := new(big.Float).SetInt(bigBalance)
	//create divisor based on decimals
	//log.Println("Decimals: ", response.Decimals)
	divisor := new(big.Float).SetFloat64(math.Pow10(response.Decimals))
	//divide balance by divisor
	balance := new(big.Float).Quo(floatBalance, divisor)
	//log.Println("Balance: ", balance.String())
	//log.Println("divisor: ", divisor.String())
	f64res, _ := balance.Float64()

	return f64res, nil
}

func (c *Client) GetPrice() (float64, error) {
	req := c.Host + "/get_price_main_coin"
	resp, err := c.hclient.Get(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()
	type Response struct {
		Price float64 `json:"price"`
	}
	var response Response
	//read body to bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return 0.0, err
	}

	// balance, err := strconv.ParseFloat(response.Price, 64)
	// if err != nil {
	// 	return 0.0, err
	// }
	return response.Price, nil
}

type DelegatedBalanceData struct {
	Validator string `json:"validator"`
	Amount    float64
}

type DelegatedBalance struct {
	Data     []DelegatedBalanceData `json:"data"`
	Decimals int                    `json:"decimals"`
	Error    string                 `json:"error"`
}

func (c *Client) GetBalanceDelegated(rpcNode string, pubkey string) (DelegatedBalance, error) {
	req := c.Host + "/get_balance_delegated?rpc_node=" + rpcNode + "&address=" + pubkey
	resp, err := c.hclient.Get(req)
	if err != nil {
		return DelegatedBalance{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DelegatedBalance{}, err
	}

	type rawDelegatedBalance struct { // used to get the raw string amount
		Data []struct {
			Validator string `json:"validator"`
			Amount    string `json:"amount"`
		} `json:"data"`
		Decimals int    `json:"decimals"`
		Error    string `json:"error"`
	}

	var rawBalance rawDelegatedBalance
	err = json.Unmarshal(body, &rawBalance)
	if err != nil {
		return DelegatedBalance{}, err
	}

	if rawBalance.Error != "" {
		return DelegatedBalance{}, errors.New("rest error: " + rawBalance.Error)
	}

	var balance DelegatedBalance
	balance.Decimals = rawBalance.Decimals
	balance.Error = rawBalance.Error
	divisor := new(big.Float).SetFloat64(math.Pow10(rawBalance.Decimals))

	for _, data := range rawBalance.Data {
		bigAmount, success := new(big.Int).SetString(data.Amount, 10)
		if !success {
			return DelegatedBalance{}, errors.New("invalid amount format")
		}

		floatAmount := new(big.Float).SetInt(bigAmount)
		adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
		f64b, _ := adjustedAmount.Float64()

		balance.Data = append(balance.Data, DelegatedBalanceData{
			Validator: data.Validator,
			Amount:    f64b,
		})
	}

	return balance, nil
}

type BeingDelegatedBalanceData struct {
	Amount                float64 `json:"amount"`
	EraDelegationFinished int     `json:"era_delegation_finished"`
	ValidatorPubkey       string  `json:"validator"`
}

type BeingDelegatedBalance struct {
	Data     []BeingDelegatedBalanceData `json:"data"`
	Decimals int                         `json:"decimals"`
}

func (c *Client) GetBalanceBeingDelegated(rpcNode string, pubkey string) (BeingDelegatedBalance, error) {
	req := c.Host + "/get_balance_being_delegated?rpc_node=" + rpcNode + "&address=" + pubkey
	resp, err := c.hclient.Get(req)
	if err != nil {
		return BeingDelegatedBalance{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return BeingDelegatedBalance{}, err
	}

	type rawBeingDelegatedBalance struct { // used to get the raw string amount
		Data []struct {
			Amount                string `json:"amount"`
			EraDelegationFinished int    `json:"era_delegation_finished"`
			ValidatorPubkey       string `json:"validator"`
		} `json:"data"`
		Decimals int `json:"decimals"`
	}

	var rawBalance rawBeingDelegatedBalance
	err = json.Unmarshal(body, &rawBalance)
	if err != nil {
		return BeingDelegatedBalance{}, err
	}

	var balance BeingDelegatedBalance
	balance.Decimals = rawBalance.Decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(rawBalance.Decimals))

	for _, data := range rawBalance.Data {
		bigAmount, success := new(big.Int).SetString(data.Amount, 10)
		if !success {
			return BeingDelegatedBalance{}, errors.New("invalid amount format")
		}

		floatAmount := new(big.Float).SetInt(bigAmount)
		adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
		f64b, _ := adjustedAmount.Float64()

		balance.Data = append(balance.Data, BeingDelegatedBalanceData{
			Amount:                f64b,
			EraDelegationFinished: data.EraDelegationFinished,
			ValidatorPubkey:       data.ValidatorPubkey,
		})
	}

	return balance, nil
}

type BeingUndelegatedBalanceData struct {
	Amount                  float64 `json:"amount"`
	EraUndelegationFinished int     `json:"era_undelegation_finished"`
	ValidatorPubkey         string  `json:"validator_pubkey"`
}

type BeingUndelegatedBalance struct {
	Data     []BeingUndelegatedBalanceData `json:"data"`
	Decimals int                           `json:"decimals"`
}

func (c *Client) GetBalanceBeingUndelegated(rpcNode string, pubkey string) (BeingUndelegatedBalance, error) {
	req := c.Host + "/get_balance_being_undelegated?rpc_node=" + rpcNode + "&address=" + pubkey
	resp, err := c.hclient.Get(req)
	if err != nil {
		return BeingUndelegatedBalance{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return BeingUndelegatedBalance{}, err
	}

	type rawBeingUndelegatedBalance struct {
		Data []struct {
			Amount                  string `json:"amount"`
			EraUndelegationFinished int    `json:"era_undelegation_finished"`
			ValidatorPubkey         string `json:"validator_pubkey"`
		} `json:"data"`
		Decimals int `json:"decimals"`
	}

	var rawBalance rawBeingUndelegatedBalance
	err = json.Unmarshal(body, &rawBalance)
	if err != nil {
		return BeingUndelegatedBalance{}, err
	}

	var balance BeingUndelegatedBalance
	balance.Decimals = rawBalance.Decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(rawBalance.Decimals))

	for _, data := range rawBalance.Data {
		bigAmount, success := new(big.Int).SetString(data.Amount, 10)
		if !success {
			return BeingUndelegatedBalance{}, errors.New("invalid amount format")
		}

		floatAmount := new(big.Float).SetInt(bigAmount)
		adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
		f64b, _ := adjustedAmount.Float64()

		balance.Data = append(balance.Data, BeingUndelegatedBalanceData{
			Amount:                  f64b,
			EraUndelegationFinished: data.EraUndelegationFinished,
			ValidatorPubkey:         data.ValidatorPubkey,
		})
	}

	return balance, nil
}

type HistoryTransfersData struct {
	DeployHash string  `json:"deploy_hash"`
	From       string  `json:"from"`
	FromPubkey string  `json:"from_pubkey"`
	To         string  `json:"to"`
	ToPubkey   string  `json:"to_pubkey"`
	Amount     float64 `json:"amount"`
	Gas        string  `json:"gas"`
	Height     uint64  `json:"height"`
}

type HistoryTransfers struct {
	Data     []HistoryTransfersData `json:"data"`
	Decimals int                    `json:"decimals"`
}

func (c *Client) GetHistoryTransfers(rpcNode string, address string, startBlock int64, endBlock int64) (HistoryTransfers, error) {
	req := fmt.Sprintf("%s/get_history_transfers?rpc_node=%s&address=%s&start_block=%d&end_block=%d",
		c.Host, rpcNode, address, startBlock, endBlock)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return HistoryTransfers{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return HistoryTransfers{}, err
	}

	type rawHistoryTransfers struct {
		Data []struct {
			DeployHash string `json:"deploy_hash"`
			From       string `json:"from"`
			FromPubkey string `json:"from_pubkey"`
			To         string `json:"to"`
			ToPubkey   string `json:"to_pubkey"`
			Amount     string `json:"amount"`
			Gas        string `json:"gas"`
			Height     uint64 `json:"height"`
		} `json:"data"`
		Decimals int `json:"decimals"`
	}

	var rawTransfers rawHistoryTransfers
	err = json.Unmarshal(body, &rawTransfers)
	if err != nil {
		return HistoryTransfers{}, err
	}

	var transfers HistoryTransfers
	transfers.Decimals = rawTransfers.Decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(rawTransfers.Decimals))

	for _, data := range rawTransfers.Data {
		bigAmount, success := new(big.Int).SetString(data.Amount, 10)
		if !success {
			return HistoryTransfers{}, errors.New("invalid amount format")
		}

		floatAmount := new(big.Float).SetInt(bigAmount)
		adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
		f64b, _ := adjustedAmount.Float64()

		transfers.Data = append(transfers.Data, HistoryTransfersData{
			DeployHash: data.DeployHash,
			From:       data.From,
			FromPubkey: data.FromPubkey,
			To:         data.To,
			ToPubkey:   data.ToPubkey,
			Amount:     f64b,
			Gas:        data.Gas,
			Height:     data.Height,
		})
	}

	return transfers, nil
}

type NetworkState struct {
	BlockHeight int64 `json:"block_height"`
	CurrentEra  int   `json:"era_id"`
}

func (c *Client) GetState(rpcNode string) (NetworkState, error) {
	req := fmt.Sprintf("%s/state?rpc_node=%s", c.Host, rpcNode)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return NetworkState{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NetworkState{}, err
	}

	var state NetworkState
	err = json.Unmarshal(body, &state)
	if err != nil {
		return NetworkState{}, err
	}

	return state, nil
}

type HistoryDelegateData struct {
	Amount          float64 `json:"amount"`
	ValidatorPubkey string  `json:"validator_pubkey"`
	Era             int     `json:"era"`
	IsFinished      bool    `json:"is_finished"`
	Height          uint64  `json:"height"`
}

type HistoryDelegate struct {
	Data     []HistoryDelegateData `json:"data"`
	Decimals int                   `json:"decimals"`
}

func (c *Client) GetHistoryDelegate(rpcNode string, address string, blockStart int64, blockEnd int64) (HistoryDelegate, error) {
	req := fmt.Sprintf("%s/get_history_delegate?rpc_node=%s&address=%s&block_start=%d&block_end=%d",
		c.Host, rpcNode, address, blockStart, blockEnd)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return HistoryDelegate{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return HistoryDelegate{}, err
	}

	type rawHistoryDelegate struct {
		Data []struct {
			Amount          string `json:"amount"`
			ValidatorPubkey string `json:"validator_pubkey"`
			Era             int    `json:"era"`
			IsFinished      bool   `json:"is_finished"`
			Height          uint64 `json:"height"`
		} `json:"data"`
		Decimals int `json:"decimals"`
	}

	var rawDelegate rawHistoryDelegate
	err = json.Unmarshal(body, &rawDelegate)
	if err != nil {
		return HistoryDelegate{}, err
	}

	var delegate HistoryDelegate
	delegate.Decimals = rawDelegate.Decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(rawDelegate.Decimals))

	for _, data := range rawDelegate.Data {
		bigAmount, success := new(big.Int).SetString(data.Amount, 10)
		if !success {
			return HistoryDelegate{}, errors.New("invalid amount format")
		}

		floatAmount := new(big.Float).SetInt(bigAmount)
		adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
		f64b, _ := adjustedAmount.Float64()

		delegate.Data = append(delegate.Data, HistoryDelegateData{
			Amount:          f64b,
			ValidatorPubkey: data.ValidatorPubkey,
			Era:             data.Era,
			IsFinished:      data.IsFinished,
			Height:          data.Height,
		})
	}

	return delegate, nil
}

type HistoryUndelegateData struct {
	Amount          float64 `json:"amount"`
	ValidatorPubkey string  `json:"validator"`
	Era             int     `json:"era"`
	IsFinished      bool    `json:"is_finished"`
	Height          uint64  `json:"height"`
}

type HistoryUndelegate struct {
	Data     []HistoryUndelegateData `json:"data"`
	Decimals int                     `json:"decimals"`
}

func (c *Client) GetHistoryUndelegate(rpcNode string, address string, blockStart int64, blockEnd int64) (HistoryUndelegate, error) {
	req := fmt.Sprintf("%s/get_history_undelegate?rpc_node=%s&address=%s&block_start=%d&block_end=%d",
		c.Host, rpcNode, address, blockStart, blockEnd)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return HistoryUndelegate{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return HistoryUndelegate{}, err
	}

	type rawHistoryUndelegate struct {
		Data []struct {
			Amount          string `json:"amount"`
			ValidatorPubkey string `json:"validator_pubkey"`
			Era             int    `json:"era"`
			IsFinished      bool   `json:"is_finished"`
			Height          uint64 `json:"height"`
		} `json:"data"`
		Decimals int `json:"decimals"`
	}

	var rawUndelegate rawHistoryUndelegate
	err = json.Unmarshal(body, &rawUndelegate)
	if err != nil {
		return HistoryUndelegate{}, err
	}

	var undelegate HistoryUndelegate
	undelegate.Decimals = rawUndelegate.Decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(rawUndelegate.Decimals))

	for _, data := range rawUndelegate.Data {
		bigAmount, success := new(big.Int).SetString(data.Amount, 10)
		if !success {
			return HistoryUndelegate{}, errors.New("invalid amount format")
		}

		floatAmount := new(big.Float).SetInt(bigAmount)
		adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
		f64b, _ := adjustedAmount.Float64()

		undelegate.Data = append(undelegate.Data, HistoryUndelegateData{
			Amount:          f64b,
			ValidatorPubkey: data.ValidatorPubkey,
			Era:             data.Era,
			IsFinished:      data.IsFinished,
			Height:          data.Height,
		})
	}

	return undelegate, nil
}

type RewardByEra struct {
	Amount    string `json:"amount"`
	Validator string `json:"validator"`
	Era       int    `json:"era"`
	Timestamp string `json:"timestamp"`
}

type RewardsByEra struct {
	Rewards  []RewardByEra `json:"rewards"`
	Decimals int           `json:"decimals"`
}

func (c *Client) GetRewardsByEra(rpcNode string, address string, startEra int, endEra int) (RewardsByEra, error) {
	req := fmt.Sprintf("%s/get_rewards_by_era?rpc_node=%s&address=%s&start_era=%d&end_era=%d",
		c.Host, rpcNode, address, startEra, endEra)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return RewardsByEra{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RewardsByEra{}, err
	}

	var rewards RewardsByEra
	err = json.Unmarshal(body, &rewards)
	if err != nil {
		return RewardsByEra{}, err
	}

	divisor := new(big.Float).SetFloat64(math.Pow10(rewards.Decimals))
	for i, reward := range rewards.Rewards {
		bigAmount, success := new(big.Int).SetString(reward.Amount, 10)
		if !success {
			return RewardsByEra{}, errors.New("invalid amount format")
		}

		floatAmount := new(big.Float).SetInt(bigAmount)
		adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
		f64b, _ := adjustedAmount.Float64()

		rewards.Rewards[i].Amount = fmt.Sprintf("%.6f", f64b)
	}

	return rewards, nil
}

func (c *Client) GetRewardsSummByEra(rpcNode string, address string, startEra int, endEra int) (float64, error) {
	req := fmt.Sprintf("%s/get_rewards_by_era?rpc_node=%s&address=%s&start_era=%d&end_era=%d",
		c.Host, rpcNode, address, startEra, endEra)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}

	var rewards RewardsByEra
	err = json.Unmarshal(body, &rewards)
	if err != nil {
		return 0.0, err
	}
	var summ float64
	summ = 0.0

	divisor := new(big.Float).SetFloat64(math.Pow10(rewards.Decimals))
	for _, reward := range rewards.Rewards {
		bigAmount, success := new(big.Int).SetString(reward.Amount, 10)
		if !success {
			return 0.0, errors.New("invalid amount format")
		}

		floatAmount := new(big.Float).SetInt(bigAmount)
		adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
		f64b, _ := adjustedAmount.Float64()

		summ += f64b
	}

	return summ, nil
}

type EraResponse struct {
	EraID     int    `json:"era_id"`
	Timestamp string `json:"timestamp"`
}

func (c *Client) GetEraByTimestamp(rpcNode string, timestamp time.Time) (EraResponse, error) {
	timestampStr := timestamp.Format(time.RFC3339)
	req := fmt.Sprintf("%s/get_era_by_timestamp?rpc_node=%s&timestamp=%s",
		c.Host, rpcNode, url.QueryEscape(timestampStr))

	resp, err := c.hclient.Get(req)
	if err != nil {
		return EraResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return EraResponse{}, err
	}

	var eraResp EraResponse
	err = json.Unmarshal(body, &eraResp)
	if err != nil {
		return EraResponse{}, err
	}

	return eraResp, nil
}

type APYResponse struct {
	APY float64 `json:"apy"`
}

func (c *Client) CalculateCurrentChainAPY(rpcNode string) (float64, error) {
	req := fmt.Sprintf("%s/calculate_current_chain_apy?rpc_node=%s",
		c.Host, rpcNode)

	resp, err := c.hclient.Get(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var apyResp APYResponse
	err = json.Unmarshal(body, &apyResp)
	if err != nil {
		return 0, err
	}

	return apyResp.APY, nil
}

type Validator struct {
	Address    string  `json:"address"`
	Fee        float32 `json:"fee"`
	Delegators int64   `json:"delegators"`
	Active     bool    `json:"active"`
}

type ValidatorsResponse struct {
	Validators []Validator `json:"validators"`
	Decimals   int         `json:"decimals"`
}

func (c *Client) GetValidators(rpcNode string) (ValidatorsResponse, error) {
	req := fmt.Sprintf("%s/get_validators?rpc_node=%s", c.Host, rpcNode)

	resp, err := c.hclient.Get(req)
	if err != nil {
		return ValidatorsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ValidatorsResponse{}, err
	}

	var validatorsResp ValidatorsResponse
	err = json.Unmarshal(body, &validatorsResp)
	if err != nil {
		return ValidatorsResponse{}, err
	}

	return validatorsResp, nil
}

type Delegate struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	BlockHeight     int    `json:"block_height"`
	IsFinished      bool   `json:"is_finished"`
	ValidatorPubkey string `json:"validator_pubkey"`
}

type Transfer struct {
	Amount     string `json:"amount"`
	DeployHash string `json:"deploy_hash"`
	From       string `json:"from"`
	FromPubkey string `json:"from_pubkey"`
	Gas        string `json:"gas"`
	To         string `json:"to"`
	ToPubkey   string `json:"to_pubkey"`
	Memo       uint64 `json:"memo"`
}

type Undelegate struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Era             int    `json:"era"`
	IsFinished      bool   `json:"is_finished"`
	ValidatorPubkey string `json:"validator_pubkey"`
}

type Reward struct {
	Delagator string `json:"delagator"`
	Validator string `json:"validator"`
	Amount    string `json:"amount"`
}

type BlockEvent struct {
	Delegates   []Delegate   `json:"delegates"`
	Transfers   []Transfer   `json:"transfers"`
	Undelegates []Undelegate `json:"undelegates"`
	Rewards     []Reward     `json:"rewards"`
	EraID       uint32       `json:"EraID"`
	Decimals    int          `json:"decimals"`
	Date        string       `json:"timestamp"`
}

func (c *Client) GetBlockEvents(rpcNode string, height uint64, transfers, delegates, undelegates bool, rewards bool) (*BlockEvent, error) {
	req := fmt.Sprintf("%s/get_block_events?rpc_node=%s&height=%d&transfers=%t&delegates=%t&undelegates=%t&rewards=%t",
		c.Host, rpcNode, height, transfers, delegates, undelegates, rewards)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var blockEvent BlockEvent
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &blockEvent)
	if err != nil {
		return nil, err
	}

	return &blockEvent, nil
}

func (c *Client) GetTimestampByBlock(rpcNode string, height uint64) (string, error) {
	req := fmt.Sprintf("%s/get_timestamp_by_block?rpc_node=%s&block_height=%d", c.Host, rpcNode, height)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type res struct {
		Timestamp string `json:"timestamp"`
	}
	var state res
	err = json.Unmarshal(body, &state)
	if err != nil {
		return "", err
	}

	return state.Timestamp, nil
}

func (c *Client) PutDeploy(rpcNode string, deploy types.Deploy) (string, error) {
	payload, err := json.Marshal(deploy)
	if err != nil {
		return "", err
	}
	req := c.Host + "/put_deploy?rpc_node=" + rpcNode
	request, err := http.NewRequest("POST", req, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	resp, err := c.hclient.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var response string
	//read body to bytes
	data, err := io.ReadAll(resp.Body)
	log.Println("RESPONSER: ", string(data))
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return "", err
	}
	return response, nil
}
