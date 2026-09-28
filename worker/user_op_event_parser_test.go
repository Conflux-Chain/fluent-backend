package worker

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/Conflux-Chain/go-conflux-util/blockchain/contract/account"
	utilstore "github.com/Conflux-Chain/go-conflux-util/store"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/sirupsen/logrus"
	logtest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func TestUnpackUserOp(t *testing.T) {
	entryPoint := common.HexToAddress("0x1001")
	paymaster := common.HexToAddress("0x2001")
	sender := common.HexToAddress("0x3001")
	other := common.HexToAddress("0x4001")
	target := contract.PackedUserOperation{
		Sender: sender, Nonce: big.NewInt(7), PreVerificationGas: big.NewInt(100),
		InitCode: common.FromHex("0x7702000000000000000000000000000000000000"),
		CallData: []byte{1}, PaymasterAndData: append(paymaster.Bytes(), make([]byte, 32)...),
	}
	wrongSender := target
	wrongSender.Sender = other
	wrongNonce := target
	wrongNonce.Nonce = big.NewInt(8)
	wrongPaymaster := target
	wrongPaymaster.PaymasterAndData = append(other.Bytes(), make([]byte, 32)...)
	duplicate := target
	duplicate.CallData = []byte{2}

	for _, tc := range []struct {
		name   string
		method string
		to     *common.Address
		ops    []contract.PackedUserOperation
		err    string
		warn   bool
	}{
		{"handleOps", entryPointMethodHandleOps, &entryPoint, []contract.PackedUserOperation{wrongSender, wrongNonce, target}, "", false},
		{"handleAggregatedOps", entryPointMethodHandleAggregatedOps, &entryPoint, []contract.PackedUserOperation{wrongSender, target}, "", false},
		{"first match", entryPointMethodHandleOps, &entryPoint, []contract.PackedUserOperation{target, duplicate}, "", false},
		{"no match", entryPointMethodHandleOps, &entryPoint, []contract.PackedUserOperation{wrongSender, wrongNonce}, "UserOperation not found", false},
		{"paymaster mismatch", entryPointMethodHandleOps, &entryPoint, []contract.PackedUserOperation{wrongPaymaster, target}, "paymaster mismatch", false},
		{"indirect call", entryPointMethodHandleOps, &other, []contract.PackedUserOperation{target}, "", true},
		{"contract creation", entryPointMethodHandleOps, nil, []contract.PackedUserOperation{target}, "", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var input []byte
			var err error
			if tc.method == entryPointMethodHandleAggregatedOps {
				var groups []account.IEntryPointUserOpsPerAggregator
				for _, op := range tc.ops {
					groups = append(groups, account.IEntryPointUserOpsPerAggregator{
						UserOps: []contract.PackedUserOperation{op},
					})
				}
				input, err = contract.EntryPointABI.Pack(tc.method, groups, common.Address{})
			} else {
				input, err = contract.EntryPointABI.Pack(tc.method, tc.ops, common.Address{})
			}
			require.NoError(t, err)
			if tc.warn {
				// Indirect calls must be skipped before interpreting another contract's input.
				input = []byte{1}
			}
			txHash := common.HexToHash("0x1234")
			tx := types.TransactionDetail{
				Hash: txHash, To: tc.to, Input: input,
				Value: big.NewInt(0), R: big.NewInt(0), S: big.NewInt(0), V: big.NewInt(0),
				// An invalid authorization must not prevent parsing finalized events.
				AuthorizationList: []ethtypes.SetCodeAuthorization{{}},
			}
			rpc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var request struct {
					ID     json.RawMessage `json:"id"`
					Method string          `json:"method"`
					Params []common.Hash   `json:"params"`
				}
				if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
					t.Errorf("decode RPC request: %v", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if request.Method != "eth_getTransactionByHash" || len(request.Params) != 1 || request.Params[0] != txHash {
					t.Errorf("unexpected RPC request: %+v", request)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(map[string]any{"jsonrpc": "2.0", "id": request.ID, "result": tx}); err != nil {
					t.Errorf("encode RPC response: %v", err)
				}
			}))
			t.Cleanup(rpc.Close)
			client, err := web3go.NewClient(rpc.URL)
			require.NoError(t, err)
			parser := UserOpEventParser{client: client, entryPointAddr: entryPoint}
			oldHooks := logrus.StandardLogger().ReplaceHooks(make(logrus.LevelHooks))
			t.Cleanup(func() { logrus.StandardLogger().ReplaceHooks(oldHooks) })
			hook := logtest.NewGlobal()
			op, err := parser.unpackUserOp(&account.EntryPointUserOperationEvent{
				Sender: sender, Nonce: big.NewInt(7), Paymaster: paymaster,
				Raw: ethtypes.Log{TxHash: txHash},
			})
			if tc.err != "" {
				require.ErrorContains(t, err, tc.err)
				require.Nil(t, op)
			} else if tc.warn {
				require.NoError(t, err)
				require.Nil(t, op)
			} else {
				require.NoError(t, err)
				require.Equal(t, target.Sender, op.Sender)
				require.Zero(t, target.Nonce.Cmp(op.Nonce))
				require.Equal(t, target.CallData, op.CallData)
			}
			if tc.warn {
				entry := hook.LastEntry()
				require.NotNil(t, entry)
				require.Equal(t, logrus.WarnLevel, entry.Level)
				require.Contains(t, entry.Message, "contract calling EntryPoint")
				require.Equal(t, txHash, entry.Data["txHash"])
				require.Equal(t, entryPoint, entry.Data["entryPoint"])
			} else {
				require.Empty(t, hook.AllEntries())
			}
		})
	}
}

func TestHandleIndexesIndirectCallWithoutRawUserOp(t *testing.T) {
	entryPoint := common.HexToAddress("0x1001")
	paymaster := common.HexToAddress("0x2001")
	sender := common.HexToAddress("0x3001")
	other := common.HexToAddress("0x4001")
	txHash := common.HexToHash("0x1234")
	opHash := common.HexToHash("0x5678")
	paymasterABI, err := contract.VerifyingPaymasterMetaData.GetAbi()
	require.NoError(t, err)
	sponsoredData, err := paymasterABI.Events["Sponsored"].Inputs.NonIndexed().Pack(true, big.NewInt(30), big.NewInt(5))
	require.NoError(t, err)
	sponsoredLog := types.Log{
		Address: paymaster, TxHash: txHash, BlockTimestamp: 1000,
		Topics: []common.Hash{eventHashSponsored, opHash}, Data: sponsoredData,
	}
	eventData, err := contract.EntryPointABI.Events["UserOperationEvent"].Inputs.NonIndexed().Pack(big.NewInt(7), true, big.NewInt(30), big.NewInt(6))
	require.NoError(t, err)
	receipt := types.Receipt{Logs: []*types.Log{{
		Address: entryPoint, TxHash: txHash, Data: eventData,
		Topics: []common.Hash{eventHashUserOperation, opHash, common.BytesToHash(sender.Bytes()), common.BytesToHash(paymaster.Bytes())},
	}}}
	tx := types.TransactionDetail{
		Hash: txHash, To: &other, Input: []byte{1},
		Value: big.NewInt(0), R: big.NewInt(0), S: big.NewInt(0), V: big.NewInt(0),
	}
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
		var result any
		switch request.Method {
		case "eth_getTransactionReceipt":
			result = receipt
		case "eth_getTransactionByHash":
			result = tx
		default:
			t.Errorf("unexpected RPC method: %s", request.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"jsonrpc": "2.0", "id": request.ID, "result": result}); err != nil {
			t.Errorf("encode RPC response: %v", err)
		}
	}))
	t.Cleanup(rpc.Close)
	client, err := web3go.NewClient(rpc.URL)
	require.NoError(t, err)
	paymasterFilterer, err := contract.NewVerifyingPaymasterFilterer(paymaster, nil)
	require.NoError(t, err)
	entryPointFilterer, err := account.NewEntryPointFilterer(entryPoint, nil)
	require.NoError(t, err)
	config := utilstore.NewMemoryConfig()
	db := config.MustOpenOrCreate(store.AllTables...)
	sqlDB, err := db.DB()
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, sqlDB.Close()) })
	scanner := UserOpEventScanner{
		store: store.NewStore(db), config: UserOpEventScanConfig{nextBlock: 1},
		parser: &UserOpEventParser{
			client: client, paymaster: paymaster, entryPointAddr: entryPoint,
			paymasterFilterer: paymasterFilterer, entryPointFilterer: entryPointFilterer,
		},
	}
	require.NoError(t, scanner.handle([]types.Log{sponsoredLog}, 10))
	var records []store.UserOp
	require.NoError(t, db.Find(&records).Error)
	require.Len(t, records, 1)
	record := records[0]
	require.Equal(t, opHash.Hex(), record.Hash)
	require.Equal(t, sender.Hex(), record.Sender)
	require.Equal(t, "0x7", record.Nonce)
	require.True(t, record.Success)
	require.Equal(t, "30", record.ActualGasCost.String())
	require.Equal(t, uint64(6), record.ActualGasUsed)
	require.Equal(t, int64(1000), record.BlockTime.Unix())
	require.Empty(t, record.RawUserOp)
	count, err := scanner.store.UserOp.GetCountByBlockTimestamp(sender, time.Unix(999, 0))
	require.NoError(t, err)
	require.Equal(t, int64(1), count)
	value, found, err := scanner.store.Config.Get(configNameEventScanNextBlock)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, "10", value)
	require.Equal(t, uint64(10), scanner.config.nextBlock)
}
