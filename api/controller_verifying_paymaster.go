package api

import (
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type VerifyingPaymasterController struct {
	services service.Services
}

func NewVerifyingPaymasterController(services service.Services) *VerifyingPaymasterController {
	return &VerifyingPaymasterController{services}
}

// Config returns the verifying paymaster configuration exposed to clients.
//
// @ID				aaPaymasterConfig
// @Summary			Get verifying paymaster configuration
// @Description		Returns the configuration of the verifying paymaster.
// @Tags			Paymaster
// @Accept			json
// @Produce			json
// @Success			200	{object}	api.BusinessError{data=VerifyingPaymasterConfig}	"Verifying paymaster configuration"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/paymaster/config	[get]
func (controller *VerifyingPaymasterController) Config(c *gin.Context) (any, error) {
	config := controller.services.Config()
	return NewVerifyingPaymasterConfig(config.VerifyingPaymaster), nil
}

// Stub returns the stub paymasterData of verifying paymaster for gas estimation.
//
// @ID				aaPaymasterStub
// @Summary			Returns the stub paymasterData of verifying paymaster for gas estimation
// @Description		Returns the stub paymasterData of verifying paymaster for gas estimation.
// @Tags			Paymaster
// @Accept			json
// @Produce			json
// @Param			request	body	VerifyingPaymasterStubRequest	true	"Verifying paymaster stub request"
// @Success			200	{object}	api.BusinessError{data=PaymasterAndDataStub}	"Paymaster address and data (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/paymaster/stub	[post]
func (controller *VerifyingPaymasterController) Stub(c *gin.Context) (any, error) {
	var input VerifyingPaymasterStubRequest

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	sender := common.HexToAddress(input.Sender)
	delegation := common.HexToAddress(input.Delegation)

	stub, err := controller.services.VerifyingPaymaster.Stub(sender, delegation)
	if err != nil {
		return nil, err
	}

	return ToPaymasterAndDataStub(stub), nil
}

// Sign validates the given user operation, signs the paymasterData and returns the reassembled paymasterData.
//
// @ID				aaPaymasterSign
// @Summary			Sign paymasterData of given user operation and return reassembled paymasterData
// @Description		Validates the given UserOperation, adds paymaster signature, and returns reassembled paymasterData.
// @Tags			Paymaster
// @Accept			json
// @Produce			json
// @Param			userOp	body	UserOperation	true	"UserOperation for paymaster signing"
// @Success			200	{object}	api.BusinessError{data=string}	"Signed and reassembled paymasterData (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/paymaster/sign	[post]
func (controller *VerifyingPaymasterController) Sign(c *gin.Context) (any, error) {
	var input UserOperation

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	userOp := input.ToPackedUserOperation()

	paymasterAndData, err := controller.services.VerifyingPaymaster.Sign(userOp)
	if err != nil {
		return nil, err
	}

	return ToPaymasterAndDataStub(paymasterAndData).Data, nil
}
