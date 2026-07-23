package api

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	assert.True(t, ok)

	auth := SetCodeAuth{
		ChainId:  1,
		Contract: "0x0000000000000000000000000000000000000002",
		Nonce:    3,
		V:        4,
		R:        "0x0000000000000000000000000000000000000000000000000000000000000005",
		S:        "0x0000000000000000000000000000000000000000000000000000000000000006",
	}
	assert.NoError(t, v.Struct(auth))

	assert.Equal(t, types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(1),
		Address: common.HexToAddress("0x0000000000000000000000000000000000000002"),
		Nonce:   3,
		V:       4,
		R:       *uint256.NewInt(5),
		S:       *uint256.NewInt(6),
	}, auth.ToGeth())
}
