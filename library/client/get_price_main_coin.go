package client

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type KucpinApiResponse struct {
	Code string `json:"code"`
	Data struct {
		Price string `json:"price"`
	} `json:"data"`
}

type OKXApiResponse struct {
	Code string `json:"code"`
	Data []struct {
		Last string `json:"last"`
	} `json:"data"`
}

func (c *Client) GetPriceMainCoin() (float64, error) {
	price, err := getKuCoinPrice()
	if err != nil {
		price, err = getOKXPrice()
		if err != nil {
			return 0, err
		}
	}
	return price, nil
}

func getKuCoinPrice() (float64, error) {
	url := "https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=CSPR-USDT"

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var apiResponse KucpinApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(apiResponse.Data.Price, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func getOKXPrice() (float64, error) {
	url := "https://www.okx.com/api/v5/market/ticker?instId=CSPR-USDT"

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var okxApiResponse OKXApiResponse
	err = json.Unmarshal(body, &okxApiResponse)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(okxApiResponse.Data[0].Last, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
