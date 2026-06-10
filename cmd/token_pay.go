package cmd

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Conflux-Chain/fluent-backend/api"
	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/cmd"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/interfaces"
	"github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var (
	testTokenPayArgs struct {
		backend          string // token pay backend URL
		rpc              string // fullnode RPC URL
		faucetPrivateKey string // private key for USDT faucet
	}

	testTokenPayCmd = cobra.Command{
		Use:   "test-token-pay",
		Short: "Test token pay functionality",
		Run:   testTokenPay,
	}
)

func init() {
	testTokenPayCmd.Flags().StringVar(&testTokenPayArgs.backend, "backend", "https://api-testnet.fluentwallet.com/api", "Backend URL for token pay API")
	testTokenPayCmd.Flags().StringVar(&testTokenPayArgs.rpc, "rpc", "https://evmtestnet.confluxrpc.com", "RPC URL for token pay API")
	testTokenPayCmd.Flags().StringVar(&testTokenPayArgs.faucetPrivateKey, "key", "", "Private key for USDT faucet")
	testTokenPayCmd.MarkFlagRequired("key")

	rootCmd.AddCommand(&testTokenPayCmd)
}

func testTokenPay(*cobra.Command, []string) {
	// create web3 client with faucet signer
	opt := web3go.ClientOption{
		SignerManager: signers.MustNewSignerManagerByPrivateKeyStrings([]string{testTokenPayArgs.faucetPrivateKey}),
	}
	client, err := web3go.NewClientWithOption(testTokenPayArgs.rpc, opt)
	cmd.FatalIfErr(err, "Failed to create web3 client")
	defer client.Close()

	// create a test account
	tester, err := NewTokenPayTester(testTokenPayArgs.backend)
	cmd.FatalIfErr(err, "Failed to create token pay tester")
	fmt.Println("Generated test account:", tester.Account())

	// faucet some USDT to test account
	fmt.Println("Begin to claim USDT from faucet ...")
	err = tester.Faucet(client, decimal.NewFromFloat(0.01))
	cmd.FatalIfErr(err, "Failed to faucet tokens")
	fmt.Println("Got USDT from faucet: 0.01 USDT")

	// prepare txs
	price, err := client.Eth.GasPrice()
	cmd.FatalIfErr(err, "Failed to retrieve gas price")

	businessTx, err := tester.PrepareBusinessTx(client, price)
	cmd.FatalIfErr(err, "Failed to prepare business transaction")
	fmt.Println("Prepared business transaction")

	transferTokenTx, err := tester.PrepareTransferTokenTx(client, businessTx)
	cmd.FatalIfErr(err, "Failed to prepare transfer token transaction")
	fmt.Println("Prepared transfer token transaction")

	// simulate txs to make sure they are valid before submission
	err = service.SimulateTx(client, businessTx, tester.Account())
	cmd.FatalIfErr(err, "Failed to simulate business transaction")
	fmt.Println("Succeeded to simulate business transaction")

	err = service.SimulateTx(client, transferTokenTx, tester.Account())
	cmd.FatalIfErr(err, "Failed to simulate transfer token transaction")
	fmt.Println("Succeeded to simulate transfer token transaction")

	// submit txs
	err = tester.Submit(transferTokenTx, businessTx)
	cmd.FatalIfErr(err, "Failed to submit token pay transactions")
	fmt.Println("Submitted token pay transactions to backend and waiting for processing ...")

	// wait for business tx receipt
	err = waitForSuccessfulReceipt(client, businessTx.Hash())
	cmd.FatalIfErr(err, "Business transaction failed")
	fmt.Println("Business transaction executed successfully")
}

type TokenPayTester struct {
	account interfaces.Signer
	backend *TokenPayClient
	config  api.TokenPayConfig
}

func NewTokenPayTester(backend string) (*TokenPayTester, error) {
	api := NewTokenPayClient(backend)

	config, err := api.GetConfig()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve token pay config")
	}

	return &TokenPayTester{
		account: signers.MustNewRandomPrivateKeySigner(),
		backend: api,
		config:  config,
	}, nil
}

func (t *TokenPayTester) Account() common.Address { return t.account.Address() }

func (t *TokenPayTester) Faucet(client *web3go.Client, amount decimal.Decimal) error {
	sm, _ := client.GetSignerManager()
	from := sm.List()[0].Address()

	usdtAddr := common.HexToAddress(t.config.Tokens[0])
	myAddr := t.account.Address()

	caller, signer := client.ToClientForContract()
	erc20, err := contract.NewERC20(usdtAddr, caller)
	if err != nil {
		return errors.WithMessage(err, "Failed to create ERC20 transactor")
	}

	decimals, err := erc20.Decimals(nil)
	if err != nil {
		return errors.WithMessage(err, "Failed to retrieve token decimals")
	}

	amountBig := decimal.New(1, int32(decimals)).Mul(amount).BigInt()

	tx, err := erc20.Transfer(&bind.TransactOpts{Signer: signer, From: from}, myAddr, amountBig)
	if err != nil {
		return errors.WithMessage(err, "Failed to transfer tokens")
	}

	return waitForSuccessfulReceipt(client, tx.Hash())
}

func (t *TokenPayTester) estimateGas(client *web3go.Client, to common.Address, data []byte) (uint64, error) {
	from := t.account.Address()

	req := types.CallRequest{
		From: &from,
		To:   &to,
		Data: data,
	}

	gas, err := client.Eth.EstimateGas(req, nil, nil, nil)
	if err != nil {
		return 0, err
	}

	return gas.Uint64(), nil
}

func (t *TokenPayTester) PrepareBusinessTx(client *web3go.Client, price *big.Int) (*gethTypes.Transaction, error) {
	from := t.account.Address()
	usdt := common.HexToAddress(t.config.Tokens[0])

	// gas fee & tip
	maxFeePerGas := new(big.Int).Mul(price, new(big.Int).SetUint64(t.config.MinGasFeeRatio+t.config.SuggestedGasPriceBumpRatio))
	maxFeePerGas.Div(maxFeePerGas, big.NewInt(100))
	maxTipPerGas := new(big.Int).Mul(price, new(big.Int).SetUint64(t.config.MinGasTipRatio+t.config.SuggestedGasPriceBumpRatio))
	maxTipPerGas.Div(maxTipPerGas, big.NewInt(100))

	// data - approve(spender, amount)
	abi, err := abi.JSON(strings.NewReader(contract.ERC20MetaData.ABI))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse ERC20 ABI")
	}

	spender := signers.MustNewRandomPrivateKeySigner().Address()
	data, err := abi.Pack("approve", spender, big.NewInt(666))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to pack approve data")
	}

	// gas limit
	gasLimit, err := t.estimateGas(client, usdt, data)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to estimate gas limit")
	}

	// nonce
	nonce, err := client.Eth.TransactionCount(from, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve nonce")
	}
	nonceU64 := hexutil.Uint64(nonce.Uint64() + 1) // +1 for pending tx

	// 1559 tx
	txType := uint8(2)

	txArgs := types.TransactionArgs{
		From:                 &from,
		To:                   &usdt,
		Gas:                  (*hexutil.Uint64)(&gasLimit),
		MaxFeePerGas:         (*hexutil.Big)(maxFeePerGas),
		MaxPriorityFeePerGas: (*hexutil.Big)(maxTipPerGas),
		Nonce:                &nonceU64,
		Data:                 (*hexutil.Bytes)(&data),
		TxType:               &txType,
	}

	if err := txArgs.Populate(client.Eth); err != nil {
		return nil, errors.WithMessage(err, "Failed to populate tx args")
	}

	tx, err := txArgs.ToTransaction()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create transaction")
	}

	return t.account.SignTransaction(tx, tx.ChainId())
}

func (t *TokenPayTester) PrepareTransferTokenTx(client *web3go.Client, businessTx *gethTypes.Transaction) (*gethTypes.Transaction, error) {
	from := t.account.Address()
	usdt := common.HexToAddress(t.config.Tokens[0])

	// data - transfer(paymaster, cost)
	abi, err := abi.JSON(strings.NewReader(contract.ERC20MetaData.ABI))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse ERC20 ABI")
	}

	data, err := abi.Pack("transfer", common.HexToAddress(t.config.Recipient), big.NewInt(666))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to pack transfer data")
	}

	// gas limit
	gasLimit, err := t.estimateGas(client, usdt, data)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to estimate gas limit")
	}

	// re-calculate token cost based on gas price & gas
	totalGasLimit := new(big.Int).SetUint64(21000 + gasLimit + businessTx.Gas())
	totalGasCost := new(big.Int).Mul(totalGasLimit, businessTx.GasPrice())
	price, err := t.backend.GetPrice(usdt)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve token price")
	}

	// add buffer to token cost estimation
	price.Mul(price, new(big.Int).SetUint64(100+t.config.SuggestedTokenPriceBumpRatio))
	price.Div(price, big.NewInt(100))

	totalTokenCost := new(big.Int).Mul(totalGasCost, price)
	totalTokenCost.Div(totalTokenCost, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	if data, err = abi.Pack("transfer", common.HexToAddress(t.config.Recipient), totalTokenCost); err != nil {
		return nil, errors.WithMessage(err, "Failed to pack real transfer data")
	}

	// re-estimate gas limit with real token cost, which may be increased a little.
	//
	// otherwise, the transaction simulation may be failed due to insufficient funds.
	if gasLimit, err = t.estimateGas(client, usdt, data); err != nil {
		return nil, errors.WithMessage(err, "Failed to estimate gas limit after token cost calculated")
	}

	// nonce
	nonce := businessTx.Nonce() - 1

	// 1559 tx
	txType := uint8(2)

	txArgs := types.TransactionArgs{
		From:                 &from,
		To:                   &usdt,
		Gas:                  (*hexutil.Uint64)(&gasLimit),
		MaxFeePerGas:         (*hexutil.Big)(businessTx.GasFeeCap()),
		MaxPriorityFeePerGas: (*hexutil.Big)(businessTx.GasTipCap()),
		Nonce:                (*hexutil.Uint64)(&nonce),
		Data:                 (*hexutil.Bytes)(&data),
		ChainID:              (*hexutil.Big)(businessTx.ChainId()),
		TxType:               &txType,
	}

	if err := txArgs.Populate(client.Eth); err != nil {
		return nil, errors.WithMessage(err, "Failed to populate tx args")
	}

	tx, err := txArgs.ToTransaction()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create transaction")
	}

	return t.account.SignTransaction(tx, tx.ChainId())
}

func (t *TokenPayTester) Submit(transferTokenTx, businessTx *gethTypes.Transaction) error {
	rawTransferTokenTx, err := transferTokenTx.MarshalBinary()
	if err != nil {
		return errors.WithMessage(err, "Failed to marshal transfer token transaction")
	}

	rawBusinessTx, err := businessTx.MarshalBinary()
	if err != nil {
		return errors.WithMessage(err, "Failed to marshal business transaction")
	}

	if err = t.backend.Submit(rawTransferTokenTx, rawBusinessTx); err != nil {
		return errors.WithMessage(err, "Failed to submit transactions")
	}

	return nil
}

func waitForSuccessfulReceipt(client *web3go.Client, txHash common.Hash) error {
	for {
		time.Sleep(time.Second)

		receipt, err := client.Eth.TransactionReceipt(txHash)
		if err != nil {
			return errors.WithMessage(err, "Failed to retrieve transaction receipt")
		}

		if receipt == nil {
			continue
		}

		if receipt.Status == nil {
			return errors.New("Transaction receipt has nil status")
		}

		if *receipt.Status == 1 {
			return nil
		}

		if receipt.TxExecErrorMsg == nil {
			return errors.New("Transaction failed without error message")
		}

		return errors.Errorf("Transaction failed with error: %s", *receipt.TxExecErrorMsg)
	}
}
