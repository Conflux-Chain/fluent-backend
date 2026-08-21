package api

import (
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
)

type VerifyingPaymasterController struct {
	services service.Services
}

func NewVerifyingPaymasterController(services service.Services) *VerifyingPaymasterController {
	return &VerifyingPaymasterController{services}
}

// Stub returns the stub paymasterData of verifying paymaster for gas estimation.
//
// @ID				aaPaymasterStub
// @Summary			Returns the stub paymasterData of verifying paymaster for gas estimation
// @Description		Returns the stub paymasterData of verifying paymaster for gas estimation.
// @Tags			Paymaster
// @Accept			json
// @Produce			json
// @Success			200	{object}	api.BusinessError{data=PaymasterAndDataStub}	"Paymaster address and data (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/paymaster/stub	[get]
func (controller *VerifyingPaymasterController) Stub(c *gin.Context) (any, error) {
	address, data := controller.services.VerifyingPaymaster.Stub()

	return PaymasterAndDataStub{
		Address: address.Hex(),
		Data:    hexutil.Encode(data),
	}, nil
}

// Sign validates the given user operation, signs the paymasterData and returns the reassembled paymasterData.
//
// @ID				aaPaymasterSign
// @Summary			Sign paymasterData of given user operation and return reassembled paymasterData
// @Description		Validates the given UserOperation, adds paymaster signature, and returns reassembled paymasterData.
// @Description		Encoding format (77 bytes): validAfter(6) || validUntil(6) || signature(65).
// @Description		Note: the on-chain paymaster contract will verify the delegated contract address, so users may be punished
// @Description 	if sending another inconsistent EIP-7702 auth message to the bundler.
// @Tags			Paymaster
// @Accept			json
// @Produce			json
// @Param			userOp	body	UserOperationWithAuth	true	"UserOperation for paymaster signing"
// @Success			200	{object}	api.BusinessError{data=string}	"Signed and reassembled paymasterData (0x-prefixed hex, 129 bytes)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/paymaster/sign	[post]
func (controller *VerifyingPaymasterController) Sign(c *gin.Context) (any, error) {
	var input UserOperationWithAuth

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	userOp := input.ToPackedUserOperation()
	delegatedContract := common.HexToAddress(input.DelegatedContract)

	paymasterData, err := controller.services.VerifyingPaymaster.Sign(userOp, delegatedContract)
	if err != nil {
		return nil, err
	}

	return hexutil.Encode(paymasterData), nil
}
