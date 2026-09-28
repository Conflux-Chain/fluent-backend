package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
