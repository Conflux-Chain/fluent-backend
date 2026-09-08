package service

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
)

const delegatedCodePrefixHex = "0xef0100" // EIP-7702 standard

var delegatedCodePrefix = hexutil.MustDecode(delegatedCodePrefixHex)

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

	if !bytes.Equal(code[0:3], delegatedCodePrefix) {
		return common.Address{}, fmt.Errorf(
			"Invalid code prefix, expected = %v, got = %v, authority = %v, code = %v",
			delegatedCodePrefixHex, hexutil.Encode(code[0:3]), authority, hexutil.Encode(code),
		)
	}

	return common.BytesToAddress(code[3:]), nil
}

// UnpackArguments unpacks ABI encoded data into the provided target structure (of pointer type).
//
// Generally, the length of the packedData should be 32x bytes, as per the ABI encoding rules.
func UnpackArguments(args abi.Arguments, packedData []byte, v any) error {
	values, err := args.Unpack(packedData)
	if err != nil {
		return errors.WithMessage(err, "Failed to unpack ABI encoded data")
	}

	if err = args.Copy(v, values); err != nil {
		return errors.WithMessage(err, "Failed to copy unpacked values to the target structure")
	}

	return nil
}
