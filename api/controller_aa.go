package api

import (
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/holiman/uint256"
	"github.com/openweb3/web3go/types/enums"
)

type AccountAbstractController struct {
	services service.Services
}

func NewAccountAbstractController(services service.Services) *AccountAbstractController {
	return &AccountAbstractController{services}
}

// SendAuth sends a transaction to set code for the given EIP-7702 auth message.
//
// @ID				aaAuthSend
// @Summary			Sends EIP-7702 transaction
// @Description		Sends a transaction with the given EIP-7702 auth message.
// @Tags			AccountAbstract
// @Accept			json
// @Produce			json
// @Param			auth	body	SetCodeAuth	true	"Standard EIP-7702 auth message"
// @Success			200	{object}	api.BusinessError{data=string}	"Transaction hash"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/auth	[post]
func (controller *AccountAbstractController) SendAuth(c *gin.Context) (any, error) {
	var input SetCodeAuth

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	r, err := uint256.FromHex(input.R)
	if err != nil {
		return nil, api.ErrValidationStrf("Invalid signature R: %v", err.Error())
	}

	s, err := uint256.FromHex(input.S)
	if err != nil {
		return nil, api.ErrValidationStrf("Invalid signature S: %v", err.Error())
	}

	return controller.services.AccountAbstract.SendSetCodeTransaction(gethTypes.SetCodeAuthorization{
		ChainID: *uint256.NewInt(input.ChainId),
		Address: common.HexToAddress(input.Contract),
		Nonce:   input.Nonce,
		V:       input.V,
		R:       *r,
		S:       *s,
	})
}

// GetAuthStatus returns the EIP-7702 auth result of given transaction hash.
//
// @ID				aaAuthStatus
// @Summary			Query EIP-7702 auth result
// @Description		Query the EIP-7702 auth result of given transaction hash. Note, if transaction not found (code 1001), user could send a transaction again.
// @Tags			AccountAbstract
// @Accept			json
// @Produce			json
// @Param			txHash	path	string	true	"EIP-7702 transaction hash"
// @Success			200	{object}	api.BusinessError{data=SetCodeResult}	"Auth result"
// @Failure			600	{object}	api.BusinessError{data=string}			"Internal server error"
// @Router			/aa/auth/{txHash}	[get]
func (controller *AccountAbstractController) GetAuthStatus(c *gin.Context) (any, error) {
	txHash := c.Param("txHash")
	if len(txHash) == 0 {
		return nil, api.ErrValidationStr("Tx hash not specified")
	}

	if len(txHash) != 66 {
		return nil, api.ErrValidationStr("Invalid tx hash")
	}

	result, ok, err := controller.services.AccountAbstract.GetSetCodeResult(common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}

	if !ok {
		return SetCodeResult{}, nil
	}

	if result.Result == enums.SetAuthOutcomeSuccess {
		return SetCodeResult{
			Executed: true,
			Success:  true,
		}, nil
	}

	return SetCodeResult{
		Executed: true,
		Error:    string(result.Result),
	}, nil
}
