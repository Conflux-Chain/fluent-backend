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

func (controller *VerifyingPaymasterController) Stub(c *gin.Context) (any, error) {
	stub := controller.services.VerifyingPaymaster.Stub()

	return hexutil.Encode(stub), nil
}

func (controller *VerifyingPaymasterController) Sign(c *gin.Context) (any, error) {
	var input UserOperationWithAuth

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	userOp := input.ToPackedUserOperation()

	delegatedContract := common.HexToAddress(input.DelegatedContract)

	paymasterData, err := controller.services.VerifyingPaymaster.Sign(userOp, delegatedContract)
	if err != nil {
		return nil, api.ErrInternal(err)
	}

	return hexutil.Encode(paymasterData), nil
}
