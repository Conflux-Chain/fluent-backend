package service

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type TokenPayConfig struct {
	Recipient           common.Address
	abiEncodedRecipient string

	MinGasFeeRatio       uint64 `default:"120"` // 120% of base fee, e.g. 20 Gdrip * 120% = 24 Gdrip
	minGasFeeRatioBig    *big.Int
	MinGasTipRatio       uint64 `default:"20"` // 20% of base fee, e.g. 20 Gdrip * 20% = 4 Gdrip
	minGasTipRatioBig    *big.Int
	MaxGasCost           uint64 `default:"100000000000000000"` // 0.1 CFX
	maxGasCostBig        *big.Int
	MinSponsorBalance    uint64 `default:"1000000000000000000"` // 1 CFX
	minSponsorBalanceBig *big.Int

	SuggestedGasPriceBumpRatio   uint64 `default:"10"` // 10% price bump for suggested gas price
	SuggestedTokenPriceBumpRatio uint64 `default:"5"`  // 5% price bump for suggested token price

	CheckReceiptInterval        time.Duration `default:"1s"`
	CheckFundingReceiptInterval time.Duration `default:"30ms"`
}

func (config *TokenPayConfig) normalize() {
	// recipient
	var buf [32]byte
	copy(buf[12:], config.Recipient.Bytes())
	config.abiEncodedRecipient = hexutil.Encode(buf[:])

	// numeric configs
	config.minGasFeeRatioBig = new(big.Int).SetUint64(config.MinGasFeeRatio)
	config.minGasTipRatioBig = new(big.Int).SetUint64(config.MinGasTipRatio)
	config.maxGasCostBig = new(big.Int).SetUint64(config.MaxGasCost)
	config.minSponsorBalanceBig = new(big.Int).SetUint64(config.MinSponsorBalance)
}
