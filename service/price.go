package service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PriceOracle struct{}

// GetETHPrice returns the price of ETH in terms of the given quote token.
func (oracle *PriceOracle) GetETHPrice(quoteToken common.Address) (*big.Int, error) {
	// TODO 1 CFX/USDT
	return big.NewInt(1_000_000_000_000_000_000), nil
}
