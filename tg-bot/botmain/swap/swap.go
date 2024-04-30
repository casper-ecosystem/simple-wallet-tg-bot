package swap

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const SimpleSwapHost = "https://api.simpleswap.io"

type ReqError struct {
	StatusCode int
	Err        error
}

func (r *ReqError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func newSwapReqErr(StatusCode int) error {
	return &ReqError{
		StatusCode: StatusCode,
		Err:        errors.New("Unknown status code != 200: "),
	}
}

type Client struct {
	ApiKey  string
	hclient *http.Client
}

func NewSwapClient(key string) *Client {
	return &Client{
		ApiKey:  key,
		hclient: &http.Client{},
	}
}

func (c *Client) GetCSPRPairs() ([]string, error) {
	req := SimpleSwapHost + "/get_pairs?api_key=" + c.ApiKey + "&fixed=false&symbol=cspr"
	resp, err := c.hclient.Get(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, newSwapReqErr(resp.StatusCode)
	}

	var response []string
	//read body to bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetEstimated(from, to string, amount float64) (float64, error) {
	strAmount := fmt.Sprintf("%f", amount)
	req := SimpleSwapHost + "/get_estimated?api_key=" + c.ApiKey + "&fixed=false&currency_from=" + from + "&currency_to=" + to + "&amount=" + strAmount
	resp, err := c.hclient.Get(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0.0, newSwapReqErr(resp.StatusCode)
	}

	var response string
	//read body to bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return 0.0, err
	}
	if s, err := strconv.ParseFloat(response, 64); err == nil {
		return s, nil
	} else {
		return 0.0, err
	}
}

func (c *Client) CheckExchange(from, to string, amount float64) (bool, error) {
	strAmount := fmt.Sprintf("%f", amount)
	req := SimpleSwapHost + "/check_exchanges?api_key=" + c.ApiKey + "&fixed=false&currency_from=" + from + "&currency_to=" + to + "&amount=" + strAmount
	resp, err := c.hclient.Get(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, newSwapReqErr(resp.StatusCode)
	}

	var response bool
	//read body to bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return false, err
	}
	return response, nil

}

func (c *Client) GetRanges(from, to string) (float64, float64, error) {
	req := SimpleSwapHost + "/get_ranges?api_key=" + c.ApiKey + "&fixed=false&currency_from=" + from + "&currency_to=" + to

	//println req url
	log.Println(req)
	resp, err := c.hclient.Get(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, 0, newSwapReqErr(resp.StatusCode)
	}

	type minMaxRes struct {
		Min string `json:"min"`
		Max string `json:"max"`
	}

	var response minMaxRes
	//read body to bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return 0, 0, err
	}
	var min, max float64
	if response.Min != "" {
		min, err = strconv.ParseFloat(response.Min, 64)
		if err != nil {
			return 0, 0, err
		}
	}
	if response.Max != "" {
		max, err = strconv.ParseFloat(response.Max, 64)
		if err != nil {
			return 0, 0, err
		}

	}
	return min, max, nil
}

type Response struct {
	ID             string    `json:"id"`
	Type           string    `json:"type"`
	Timestamp      time.Time `json:"timestamp"`
	UpdatedAt      time.Time `json:"updated_at"`
	CurrencyFrom   string    `json:"currency_from"`
	CurrencyTo     string    `json:"currency_to"`
	AmountFrom     string    `json:"amount_from"`
	ExpectedAmount string    `json:"expected_amount"`
	AmountTo       string    `json:"amount_to"`
	AddressFrom    string    `json:"address_from"`
	AddressTo      string    `json:"address_to"`

	UserRefundAddress string `json:"user_refund_address"`
	UserRefundExtraID string `json:"user_refund_extra_id"`

	Status      string `json:"status"`
	RedirectURL string `json:"redirect_url"`
}

func (c *Client) MakeExchange(from, to string, addressFrom string, AddressTo string, amount float64, extraId string) (Response, error) {

	type Payload struct {
		Fixed             bool    `json:"fixed"`
		CurrencyFrom      string  `json:"currency_from"`
		CurrencyTo        string  `json:"currency_to"`
		Amount            float64 `json:"amount"`
		AddressTo         string  `json:"address_to"`
		ExtraIDTo         string  `json:"extra_id_to"`
		UserRefundAddress string  `json:"user_refund_address"`
		UserRefundExtraID string  `json:"user_refund_extra_id"`
	}

	data := Payload{
		Fixed:             false,
		CurrencyFrom:      from,
		CurrencyTo:        to,
		Amount:            amount,
		AddressTo:         AddressTo,
		UserRefundAddress: addressFrom,
		ExtraIDTo:         extraId,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return Response{}, nil
	}
	body := bytes.NewReader(payloadBytes)

	fmt.Println(string(payloadBytes))

	req, err := http.NewRequest("POST", "https://api.simpleswap.io/create_exchange?api_key="+c.ApiKey, body)
	if err != nil {
		return Response{}, nil
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		dataResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return Response{}, nil
		}
		log.Println(string(dataResp))
		return Response{}, newSwapReqErr(resp.StatusCode)
	}

	var response Response
	//read body to bytes
	dataResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, nil
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return Response{}, nil
	}
	return response, nil
}
