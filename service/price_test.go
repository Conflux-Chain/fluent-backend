package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"github.com/openweb3/web3go"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

// Initialize through the real constructor using a local ERC20 decimals RPC stub.
func newTestPriceOracle(t *testing.T, timeout time.Duration) *PriceOracle {
	t.Helper()
	rpc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Errorf("decode RPC request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if request.Method != "eth_call" {
			t.Errorf("unexpected RPC method: %s", request.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"result":"0x%064x"}`, request.ID, 6)
	}))
	t.Cleanup(rpc.Close)

	client, err := web3go.NewClient(rpc.URL)
	require.NoError(t, err)
	oracle, err := NewPriceOracle(PriceConfig{
		USDT:           []common.Address{common.HexToAddress("0x1234")},
		RequestTimeout: timeout,
	}, client)
	require.NoError(t, err)
	return oracle
}

func TestGetETHPriceValidationAndConversion(t *testing.T) {
	usdt := common.HexToAddress("0x1234")
	cnh := common.HexToAddress("0x5678")
	for _, tc := range []struct {
		name    string
		token   common.Address
		binance string
		okx     string
		want    string
		wantErr string
	}{
		{"USDT normal", usdt, "0.2512349", "7", "251234", ""},
		{"CNH normal", cnh, "0.25", "7.1234567", "1780864", ""},
		{"USDT zero", usdt, "0", "7", "", "Binance price must be greater than 0"},
		{"USDT negative", usdt, "-0.25", "7", "", "Binance price must be greater than 0"},
		{"USDT malformed", usdt, "invalid", "7", "", "Failed to parse binance price"},
		{"CNH Binance zero", cnh, "0", "7", "", "Binance price must be greater than 0"},
		{"CNH Binance negative", cnh, "-0.25", "7", "", "Binance price must be greater than 0"},
		{"CNH Binance malformed", cnh, "invalid", "7", "", "Failed to parse binance price"},
		{"CNH OKX zero", cnh, "0.25", "0", "", "OKX price must be greater than 0"},
		{"CNH OKX negative", cnh, "0.25", "-7", "", "OKX price must be greater than 0"},
		{"CNH OKX malformed", cnh, "0.25", "invalid", "", "Failed to parse OKX price"},
		{"CNH both negative", cnh, "-0.25", "-7", "", "Binance price must be greater than 0"},
		{"USDT truncates to zero", usdt, "0.0000009", "7", "", "Price in token smallest units must be greater than 0"},
		{"CNH truncates to zero", cnh, "0.0000001", "7", "", "Price in token smallest units must be greater than 0"},
		{"USDT smallest unit", usdt, "0.000001", "7", "1", ""},
		{"CNH smallest unit", cnh, "0.0000002", "5", "1", ""},
		{"CNH truncate after multiplication", cnh, "0.0000002", "7", "1", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/binance":
					fmt.Fprintf(w, `{"symbol":"CFXUSDT","price":%q}`, tc.binance)
				case "/okx":
					fmt.Fprintf(w, `{"code":0,"error_code":"0","data":[{"price":%q}]}`, tc.okx)
				default:
					http.NotFound(w, r)
				}
			}))
			t.Cleanup(server.Close)

			client := resty.New().SetTimeout(time.Second)
			// Route both production URLs to local fixtures without external network access.
			client.OnBeforeRequest(func(_ *resty.Client, request *resty.Request) error {
				switch request.URL {
				case binancePriceUrlCFXUSDT:
					request.URL = server.URL + "/binance"
				case okxPriceUrlUSDTCNY:
					request.URL = server.URL + "/okx"
				default:
					return fmt.Errorf("unexpected price URL: %s", request.URL)
				}
				return nil
			})
			oracle := PriceOracle{
				client:     client,
				usdtTokens: map[common.Address]ERC20TokenStub{usdt: {decimalExp: decimal.New(1, 6)}},
				cnhTokens:  map[common.Address]ERC20TokenStub{cnh: {decimalExp: decimal.New(1, 6)}},
			}

			price, err := oracle.GetETHPrice(tc.token)
			if tc.wantErr != "" {
				require.ErrorContains(t, err, tc.wantErr)
				require.Nil(t, price)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, price.String())
			}
		})
	}
}

func TestPriceOracleRequestTimeoutConfig(t *testing.T) {
	for _, tc := range []struct {
		name    string
		timeout time.Duration
		want    time.Duration
	}{
		{"default", 0, 3 * time.Second},
		{"configured", 250 * time.Millisecond, 250 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			oracle := newTestPriceOracle(t, tc.timeout)
			require.Equal(t, tc.want, oracle.client.GetClient().Timeout)
		})
	}

	t.Run("negative rejected before RPC", func(t *testing.T) {
		oracle, err := NewPriceOracle(PriceConfig{
			USDT:           []common.Address{common.HexToAddress("0x1234")},
			RequestTimeout: -time.Second,
		}, nil)
		require.Nil(t, oracle)
		require.EqualError(t, err, "PriceConfig RequestTimeout must be greater than 0")
	})
}

func TestPriceOracleHTTPTimeout(t *testing.T) {
	for _, source := range []struct {
		name  string
		body  string
		fetch func(*PriceOracle, string) (decimal.Decimal, error)
	}{
		{"Binance", `{"symbol":"CFXUSDT","price":"0.25"}`, (*PriceOracle).getBinancePrice},
		{"OKX", `{"code":0,"error_code":"0","data":[{"price":"0.25"}]}`, (*PriceOracle).getOkxPrice},
	} {
		t.Run(source.name, func(t *testing.T) {
			for _, mode := range []string{"normal", "slow headers", "slow body"} {
				t.Run(mode, func(t *testing.T) {
					oracle := newTestPriceOracle(t, 100*time.Millisecond)
					server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						if mode == "slow body" {
							fmt.Fprint(w, " ")
							w.(http.Flusher).Flush()
						}
						if mode != "normal" {
							select {
							case <-r.Context().Done():
								return
							case <-time.After(2 * time.Second):
								// Bound the test even if the client timeout regresses.
							}
						}
						fmt.Fprint(w, source.body)
					}))
					t.Cleanup(server.Close)

					price, err := source.fetch(oracle, server.URL)
					if mode == "normal" {
						require.NoError(t, err)
						require.True(t, price.Equal(decimal.RequireFromString("0.25")))
					} else {
						require.ErrorContains(t, err, "context deadline exceeded")
						require.True(t, price.IsZero())
					}
				})
			}
		})
	}
}
