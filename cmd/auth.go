package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/Conflux-Chain/go-conflux-util/cmd"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/signers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	authCmdArgs struct {
		URL        string
		PrivateKey string
		Contract   string
	}

	authCmd = cobra.Command{
		Use:   "auth",
		Short: "Generate EIP-7702 auth message",
		Run:   generateAuth,
	}
)

func init() {
	authCmd.Flags().StringVar(&authCmdArgs.URL, "url", "https://evmtestnet.confluxrpc.com", "Fullnode RPC URL")
	authCmd.Flags().StringVar(&authCmdArgs.PrivateKey, "key", "0x1111111111222222222233333333334444444444555555555566666666667777", "Private key to sign auth message")
	authCmd.Flags().StringVar(&authCmdArgs.Contract, "contract", "", "Delegated contract address, empty indicates reset code")

	rootCmd.AddCommand(&authCmd)
}

func generateAuth(*cobra.Command, []string) {
	// signer
	sm := signers.MustNewSignerManagerByPrivateKeyStrings([]string{authCmdArgs.PrivateKey})
	signer := sm.List()[0]
	fmt.Println("Signer:", signer.Address())

	// rpc client
	client, err := web3go.NewClient(authCmdArgs.URL)
	cmd.FatalIfErr(err, "Failed to create RPC client")
	defer client.Close()

	// chain id
	chainId, err := client.Eth.ChainId()
	cmd.FatalIfErr(err, "Failed to retrieve chain ID")

	var chainIdU64 uint64
	if chainId == nil {
		logrus.Warn("Chain ID unavailable")
	} else {
		chainIdU64 = *chainId
	}
	fmt.Println("Chain ID:", chainIdU64)

	// nonce
	nonce, err := client.Eth.TransactionCount(signer.Address(), nil)
	cmd.FatalIfErr(err, "Failed to retrieve signer nonce")
	fmt.Println("Signer nonce:", nonce)

	// construct auth message and sign
	auth := types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(chainIdU64),
		Address: common.HexToAddress(authCmdArgs.Contract),
		Nonce:   nonce.Uint64(),
	}

	signedAuth, err := signer.SignSetCodeAuthorization(auth)
	cmd.FatalIfErr(err, "Failed to sign auth message")

	content, err := json.MarshalIndent(signedAuth, "", "    ")
	cmd.FatalIfErr(err, "Failed to marshal signed auth message")
	fmt.Println(string(content))
}
