package service

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/pkg/errors"
)

var (
	ErrRPCError = api.NewBusinessError(101, "RPC error, please try again later")

	// account abstract
	ErrAccountAbstractTxNotFound = api.NewBusinessError(1001, "Set code transaction not found")
)

func NewRPCError(err error, message string, args ...any) *api.BusinessError {
	if len(args) == 0 {
		err = errors.WithMessage(err, message)
	} else {
		err = errors.WithMessagef(err, message, args...)
	}

	return ErrRPCError.WithData(err.Error())
}
