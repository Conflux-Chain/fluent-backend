package api

import (
	"math/big"

	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type GasTankController struct {
	services service.Services
}

func NewGasTankController(services service.Services) *GasTankController {
	return &GasTankController{services}
}

// PrepareCredit builds paymasterData for CREDIT mode deposit estimation.
//
// The returned paymasterData should be attached to the UserOperation when wallet clients estimate
// gas in CREDIT mode before submitting a signed operation.
//
// @ID				aaGasTankPrepareCredit
// @Summary			Prepare paymasterData for CREDIT mode deposit estimation
// @Description		Returns paymasterData used to estimate UserOperation gas in CREDIT mode for token deposit flow.
// @Tags			GasTank
// @Accept			json
// @Produce			json
// @Param			request		body	GasTankPrepareCreditRequest	true	"Paymaster data request"
// @Success			200	{object}	api.BusinessError{data=PaymasterAndDataStub}	"Paymaster address and data (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/gastank/stub/credit	[post]
func (controller *GasTankController) PrepareCredit(c *gin.Context) (any, error) {
	var input GasTankPrepareCreditRequest

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	token := common.HexToAddress(input.Token)
	amount, ok := new(big.Int).SetString(input.Amount, 10)
	if !ok || amount.Sign() <= 0 {
		return nil, api.ErrValidationStr("Invalid amount")
	}

	paymasterData, err := controller.services.GasTank.StubCredit(token, amount)
	if err != nil {
		return nil, err
	}

	return ToPaymasterAndDataStub(paymasterData), nil
}

// PrepareRefund builds paymasterData for normal UserOperation gas estimation in REFUND mode.
//
// The returned paymasterData should be attached to the UserOperation when wallet clients estimate
// gas in REFUND mode before signing and sending the operation.
//
// @ID				aaGasTankPrepareRefund
// @Summary			Prepare paymasterData for REFUND mode estimation
// @Description		Returns paymasterData used to estimate gas for a normal UserOperation in REFUND mode.
// @Tags			GasTank
// @Accept			json
// @Produce			json
// @Param			request		body	GasTankPrepareRefundRequest	true	"Paymaster data request"
// @Success			200	{object}	api.BusinessError{data=PaymasterAndDataStub}	"Paymaster address and data (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/gastank/stub/refund	[post]
func (controller *GasTankController) PrepareRefund(c *gin.Context) (any, error) {
	var input GasTankPrepareRefundRequest

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	sender := common.HexToAddress(input.Sender)
	token := common.HexToAddress(input.Token)

	paymasterData, err := controller.services.GasTank.StubRefund(sender, token)
	if err != nil {
		return nil, err
	}

	return ToPaymasterAndDataStub(paymasterData), nil
}

// Sign calculates maxTokenCost for a UserOperation, signs it with paymaster key, and
// returns the reassembled paymasterData.
//
// The wallet should call this endpoint after gas estimation, then place the returned paymasterData
// into UserOperation.paymasterData before final submission.
//
// @ID				aaGasTankSign
// @Summary			Sign paymasterData for UserOperation
// @Description		Calculates maxTokenCost for the given UserOperation, adds paymaster signature, and returns reassembled paymasterData.
// @Description		Encoding format (130 bytes): mode(1) || token(20) || maxTokenCost(32) || validAfter(6) || validUntil(6) || signature(65). maxTokenCost is 0-based, big-endian uint256.
// @Description		CREDIT mode reminder: in input UserOperation.paymasterData, maxTokenCost is the token deposit amount and should be the amount to deposit.
// @Tags			GasTank
// @Accept			json
// @Produce			json
// @Param			userOp	body	UserOperation	true	"UserOperation for maxTokenCost estimation and paymaster signing"
// @Success			200	{object}	api.BusinessError{data=string}	"Signed and reassembled paymasterData (0x-prefixed hex, 130 bytes)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/gastank/sign	[post]
func (controller *GasTankController) Sign(c *gin.Context) (any, error) {
	var input UserOperation

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	userOp := input.ToPackedUserOperation()

	paymasterAndData, err := controller.services.GasTank.Sign(userOp)
	if err != nil {
		return nil, err
	}

	return ToPaymasterAndDataStub(paymasterAndData).Data, nil
}
