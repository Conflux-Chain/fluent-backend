package service

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
)

const delegatedCodePrefix = "0xef0100" // EIP-7702 standard

// GetDelegatedContract reads the on-chain code of authority and extracts the 20-byte contract
// address from the EIP-7702 delegation designator (prefix 0xef0100 + address).
// Returns empty address when the authority has no code (not yet delegated).
func GetDelegatedContract(client *web3go.Client, authority common.Address) (common.Address, error) {
	code, err := client.Eth.CodeAt(authority, nil)
	if err != nil {
		return common.Address{}, NewRPCError(err, "Failed to retrieve authority code")
	}

	codeLen := len(code)
	if codeLen == 0 {
		return common.Address{}, nil
	}

	if codeLen != 23 {
		return common.Address{}, fmt.Errorf(
			"Invalid code length, expected = 23, got = %v, authority = %v, code = %v",
			codeLen, authority, hexutil.Encode(code),
		)
	}

	if prefix := hexutil.Encode(code[0:3]); prefix != delegatedCodePrefix {
		return common.Address{}, fmt.Errorf(
			"Invalid code prefix, expected = %v, got = %v, authority = %v, code = %v",
			delegatedCodePrefix, prefix, authority, hexutil.Encode(code),
		)
	}

	return common.BytesToAddress(code[3:]), nil
}
