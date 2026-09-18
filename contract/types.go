package contract

import (
	"strings"

	"github.com/Conflux-Chain/go-conflux-util/blockchain/contract/account"
	"github.com/Conflux-Chain/go-conflux-util/blockchain/contract/token/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type PackedUserOperation = account.PackedUserOperation

var (
	ERC20ABI        = mustInitABI(erc20.ContractMetaData.ABI)
	EntryPointABI   = mustInitABI(account.EntryPointMetaData.ABI)
	SmartAccountABI = mustInitABI(SimpleSmartAccount7702MetaData.ABI)
)

// mustInitABI initializes an ABI from its JSON representation.
// It panics if the ABI cannot be parsed.
//
// It should not panic, because the ABI JSON is auto-generated via abigen tool.
func mustInitABI(ABI string) abi.ABI {
	result, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		panic(err)
	}

	return result
}
