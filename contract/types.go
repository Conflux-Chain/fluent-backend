package contract

import (
	"github.com/Conflux-Chain/go-conflux-util/blockchain/contract/account"
	uniswapv2 "github.com/Conflux-Chain/go-conflux-util/blockchain/contract/defi/uniswap/v2"
	"github.com/Conflux-Chain/go-conflux-util/blockchain/contract/token/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type PackedUserOperation = account.PackedUserOperation

var (
	ERC20ABI           = mustGetABI(erc20.ContractMetaData)
	EntryPointABI      = mustGetABI(account.EntryPointMetaData)
	SmartAccountABI    = mustGetABI(SimpleSmartAccount7702MetaData)
	UniswapV2RouterABI = mustGetABI(uniswapv2.RouterMetaData)
)

// mustGetABI retrieves the ABI from the given metadata.
// This should not panic, because the ABI JSON is auto-generated via abigen tool.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}
