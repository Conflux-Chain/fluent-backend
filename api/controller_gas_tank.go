package api

import (
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

// Stub builds paymasterData for UserOperation gas estimation.
//
// @ID				aaGasTankStub
// @Summary			Prepare paymasterData for gas estimation
// @Description		Returns paymasterData used for user operation gas estimation.
// @Tags			GasTank
// @Accept			json
// @Produce			json
// @Param			request		body	GasTankStubRequest	true	"Paymaster data stub request"
// @Success			200	{object}	api.BusinessError{data=PaymasterAndDataStub}	"Paymaster address and data (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/gastank/stub	[post]
func (controller *GasTankController) Stub(c *gin.Context) (any, error) {
	var input GasTankStubRequest

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	sender := common.HexToAddress(input.Sender)
	token := common.HexToAddress(input.Token)

	paymasterData, err := controller.services.GasTank.Stub(sender, token)
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
// @Description		Encoding format (129 bytes): token(20) || maxTokenCost(32) || validAfter(6) || validUntil(6) || signature(65). maxTokenCost is hex encoded.
// @Tags			GasTank
// @Accept			json
// @Produce			json
// @Param			userOp	body	UserOperation	true	"UserOperation for maxTokenCost estimation and paymaster signing"
// @Success			200	{object}	api.BusinessError{data=string}	"Signed and reassembled paymasterData"
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
