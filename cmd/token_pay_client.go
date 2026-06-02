package cmd

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Conflux-Chain/fluent-backend/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type TokenPayClient struct {
	inner *resty.Client
}

func NewTokenPayClient(baseURL string) *TokenPayClient {
	return &TokenPayClient{
		inner: resty.New().SetBaseURL(baseURL).SetTimeout(5 * time.Second),
	}
}

func (client *TokenPayClient) GetConfig() (api.TokenPayConfig, error) {
	return doHttpRequest[api.TokenPayConfig](resty.MethodGet, "/tokenpay/config", client.inner.R())
}

func (client *TokenPayClient) GetPrice(token common.Address) (*big.Int, error) {
	price, err := doHttpRequest[string](resty.MethodGet, "/tokenpay/price", client.inner.R().SetQueryParam("token", token.Hex()))
	if err != nil {
		return nil, err
	}

	priceInt, ok := new(big.Int).SetString(price, 10)
	if !ok {
		return nil, fmt.Errorf("invalid price format: %s", price)
	}

	return priceInt, nil
}

func (client *TokenPayClient) Submit(rawTransferTokenTx, rawBusinessTx []byte) error {
	_, err := doHttpRequest[any](resty.MethodPost, "/tokenpay/submit",
		client.inner.R().
			SetHeader("Content-Type", "application/json").
			SetBody(api.TokenPayRequest{
				RawTransferTokenTx: hexutil.Encode(rawTransferTokenTx),
				RawBusinessTx:      hexutil.Encode(rawBusinessTx),
			}),
	)

	return err
}

func doHttpRequest[T any](method, url string, request *resty.Request) (data T, err error) {
	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    T      `json:"data"`
	}

	response, err := request.SetResult(&result).Execute(method, url)
	if err != nil {
		return data, errors.WithMessage(err, "Failed to do HTTP request")
	}

	if response.IsError() {
		return data, fmt.Errorf("HTTP request failed, status = %v", response.Status())
	}

	if result.Code != 0 {
		return data, fmt.Errorf("API error, code = %v, message = %v, data = %v", result.Code, result.Message, result.Data)
	}

	return result.Data, nil
}
