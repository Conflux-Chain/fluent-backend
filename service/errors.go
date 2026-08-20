package service

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/pkg/errors"
)

var (
	ErrRPCError = api.NewBusinessError(101, "RPC error, please try again later or check the detailed error message for more information")

	// account abstract
	ErrAccountAbstractTxNotFound = api.NewBusinessError(1001, "Set code transaction not found")

	// Gas tank
	ErrGasTankInsufficientBalance = api.NewBusinessError(2001, "Insufficient token balance in gas tank")

	// token pay
	ErrTokenPayPriceTooLow               = api.NewBusinessError(3001, "Gas price too low")
	ErrTokenPayGasCostTooHigh            = api.NewBusinessError(3002, "Gas cost too high")
	ErrTokenPayTransferTokenAmountTooLow = api.NewBusinessError(3003, "Transferred token amount too low")
	ErrTokenPaySimulateTxFailed          = api.NewBusinessError(3004, "Failed to simulate transaction")
	ErrTokenPaySponsorBalanceNotEnough   = api.NewBusinessError(3005, "Sponsor balance is insufficient, please contact the administrator")

	// verifying paymaster
	ErrVerifyingPaymasterMaxGasCostExceeded     = api.NewBusinessError(4001, "Max gas cost exceeded")
	ErrVerifyingPaymasterInvalidSmartAccount    = api.NewBusinessError(4002, "Smart account is not in whitelist")
	ErrVerifyingPaymasterContractNotWhitelisted = api.NewBusinessError(4003, "Contract is not in whitelist")
	ErrVerifyingPaymasterPaused                 = api.NewBusinessError(4004, "Verifying paymaster contract is paused")
)

func NewRPCError(err error, message string, args ...any) *api.BusinessError {
	if len(args) == 0 {
		err = errors.WithMessage(err, message)
	} else {
		err = errors.WithMessagef(err, message, args...)
	}

	return ErrRPCError.WithData(err.Error())
}
